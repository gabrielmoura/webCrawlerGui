package db

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/log"
	"errors"
	"fmt"
	"path"
	"sort"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/vrischmann/userdir"
	"go.uber.org/zap"
)

var (
	// blockWrite é mutex para controle de otimização dos logs
	blockWrite sync.RWMutex
)

var DB *Database

type Database struct {
	db *badger.DB
}

func getDbPath(appName string) string {
	dbPath := path.Join(userdir.GetDataHome(), appName)
	log.Logger.Info("BD", zap.String("path", dbPath))
	return dbPath
}

// InitDB inicializa a sessão do banco de dados
func InitDB(appName string) *Database {
	opts := badger.DefaultOptions(getDbPath(appName))
	//opts.Logger = nil // Desativa logs de debug
	opts.CompactL0OnClose = true
	opts.NumCompactors = 2
	opts.ValueLogFileSize = 100 << 20 // 100 MB

	db, err := badger.Open(opts)
	if err != nil {
		log.Logger.Fatal("error initializing database", zap.Error(err))
	}

	cdb := &Database{
		db: db,
	}
	//go cdb.OptimizeCache(5)
	DB = cdb
	return cdb
}

// CloseDB fecha a sessão do banco de dados
func (d Database) CloseDB() error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	return d.db.Close()
}
func (d Database) SyncCache() error {
	time.Sleep(1 * time.Second)
	log.Logger.Info("Syncing cache")

	visited, err := d.AllVisited()
	if err != nil {
		return fmt.Errorf("error getting all visited: %v", err)
	}
	if len(visited) == 0 {
		return nil
	}

	for _, link := range visited {
		err := d.SetVisited(link)
		if err != nil {
			return fmt.Errorf("error setting visitedIndex: %v", err)
		}
	}
	return nil
}

// OptimizeCache Otimiza Banco a cada n minutos
func (d Database) OptimizeCache(nMinute int) {
	for {
		time.Sleep(time.Duration(nMinute) * time.Minute)
		log.Logger.Info("Optimizing cache")
		blockWrite.Lock()
		err := d.db.RunValueLogGC(0.9)
		if err != nil && !errors.Is(badger.ErrNoRewrite, err) {
			log.Logger.Info("error optimizing cache", zap.Error(err))
		}
		blockWrite.Unlock()
	}
}
func (d Database) IsVisited(url string) bool {
	key := []byte(fmt.Sprintf("%s:%s", config.VisitedIndexName, url))
	err := d.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
func (d Database) SetVisited(url string) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.VisitedIndexName, url))
	err := d.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, []byte{})
		if err != nil {
			return err
		}
		return nil // return nil to commit the transaction
	})
	if err != nil {
		return fmt.Errorf("error setting visited: %v", err)
	}
	return nil
}

func (d Database) AddToQueue(url string, depth int) error {
	err := d.Enqueue(url, depth)
	if err != nil {
		return fmt.Errorf("error adding to queue: %v", err)
	}
	return nil
}
func (d Database) GetFromQueue() (string, int, error) {
	url, depth, err := d.Dequeue()
	if err != nil {
		return "", 0, fmt.Errorf("error getting from queue: %v", err)
	}
	return url, depth, nil
}
func (d Database) GetFromQueueV2(getNumber int) ([]data.QueueType, error) {
	var urls []data.QueueType
	for i := 0; i < getNumber; i++ {
		url, depth, err := d.Dequeue()
		if err != nil {
			return nil, fmt.Errorf("error getting from queue: %v", err)
		}
		if url != "" {
			urls = append(urls, data.QueueType{Url: url, Depth: depth})
		}
	}
	return urls, nil
}

// WritePage insere uma nova página no banco de dados
func (d Database) WritePage(page *data.Page) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	Key := []byte(fmt.Sprintf("%s:%s", config.PageDataIndexName, page.Url))

	return d.db.Update(func(txn *badger.Txn) error {
		bytes, err := pageMarshal(page)
		if err != nil {
			return err
		}
		return txn.Set(Key, bytes)
	})
}

// ReadPage recupera uma página do banco de dados por URL
func (d Database) ReadPage(url string) (*data.Page, error) {
	var page data.Page
	Key := []byte(fmt.Sprintf("%s:%s", config.PageDataIndexName, url))

	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(Key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return pageUnmarshal(val, &page)
		})
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	return &page, err
}

// AllVisited recupera todos os URLs visitados
func (d Database) AllVisited() ([]string, error) {
	var urls []string
	prefixKey := []byte(fmt.Sprintf("%s:", config.VisitedIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}

			if page.Visited {
				urls = append(urls, page.Url)
			}
		}
		return nil
	})

	return urls, err
}

// SearchByTitleOrDescription pesquisa páginas por título ou descrição
func (d Database) SearchByTitleOrDescription(searchTerm string) ([]data.PageSearch, error) {
	var pages []data.PageSearch
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}

			if containsIgnoreCase(page.Title, searchTerm) || containsIgnoreCase(page.Description, searchTerm) {
				pages = append(pages, data.PageSearch{
					Url:   page.Url,
					Title: page.Title,
				})
			}
		}
		return nil
	})

	return pages, err
}

// SearchByContent pesquisa páginas por conteúdo e ordena por frequência
func (d Database) SearchByContent(searchTerm string) ([]data.PageSearchWithFrequency, error) {
	var pages []data.PageSearchWithFrequency
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}

			if freq, ok := page.Words[searchTerm]; ok {
				pages = append(pages, data.PageSearchWithFrequency{
					Url:       page.Url,
					Title:     page.Title,
					Frequency: freq,
				})
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Sort by frequency in descending order
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Frequency > pages[j].Frequency
	})

	return pages, nil
}

// Search queries pages by title, description, or content.
func (d Database) Search(searchTerm string) ([]data.Page, error) {
	var pages []data.Page
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}
			if containsIgnoreCase(page.Title, searchTerm) ||
				containsIgnoreCase(page.Description, searchTerm) ||
				containsWord(page.Words, searchTerm) {
				pages = append(pages, page)
			}
		}
		return nil
	})

	return pages, err
}

func (d Database) SearchWords(searchTerms []string) ([]data.Page, error) {
	var pages []data.Page
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}

			for _, search := range searchTerms {
				if containsWord(page.Words, search) {
					pages = append(pages, page)
				}
			}
		}
		return nil
	})

	return pages, err
}

func (d Database) GetAllPage() ([]data.Page, error) {
	var pages []data.Page
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}
			pages = append(pages, page)
		}
		return nil
	})

	return pages, err
}

func (d Database) SearchV2(searchTerms []string) ([]data.Page, error) {
	var pages []data.Page
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataIndexName))

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixKey
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixKey); it.ValidForPrefix(prefixKey); it.Next() {
			item := it.Item()
			var page data.Page

			if err := item.Value(func(val []byte) error {
				return pageUnmarshal(val, &page)
			}); err != nil {
				return err
			}

			if containsIgnoreCaseSlice(page.Title, searchTerms) || containsIgnoreCaseSlice(page.Description, searchTerms) {
				pages = append(pages, page)
			}
			for _, search := range searchTerms {
				if containsWord(page.Words, search) {
					pages = append(pages, page)
				}
			}

		}
		return nil
	})

	return pages, err
}
