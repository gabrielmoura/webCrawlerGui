package db

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/helper"
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/log"
	"errors"
	"fmt"
	"net/url"
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
	DB         *Database
)

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
	//opts.ValueLogFileSize = 100 << 20 // 100 MB

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
		d.OptimizeCacheNow()
	}
}
func (d Database) OptimizeCacheNow() {
	log.Logger.Info("Optimizing cache")
	blockWrite.Lock()
	err := d.db.RunValueLogGC(0.9)
	if err != nil && !errors.Is(badger.ErrNoRewrite, err) {
		log.Logger.Info("error optimizing cache", zap.Error(err))
	}
	blockWrite.Unlock()
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
	return err == nil
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

	pageKey := []byte(fmt.Sprintf("%s:%s", config.PageDataPrefix, page.Url))
	indexKey := []byte(config.PageDataIndexName)

	return d.db.Update(func(txn *badger.Txn) error {
		// Marshal the page data
		pageBytes, err := pageMarshal(page)
		if err != nil {
			return fmt.Errorf("error marshaling page: %w", err)
		}

		// Set the page data in the database
		if err := txn.Set(pageKey, pageBytes); err != nil {
			return fmt.Errorf("error setting page data: %w", err)
		}

		// Get the index from the database
		item, err := txn.Get(indexKey)
		var pageIndex data.PageIndex

		if err != nil {
			if errors.Is(badger.ErrKeyNotFound, err) {
				// If the index does not exist, initialize it
				pageIndex = data.PageIndex{}
			} else {
				return fmt.Errorf("error getting index: %w", err)
			}
		} else {
			if err := item.Value(func(val []byte) error {
				return indexPageUnmarshal(val, &pageIndex)
			}); err != nil {
				return fmt.Errorf("error unmarshaling index: %w", err)
			}
		}

		// Update the index with the new page URL
		pageIndex.Keys = append(pageIndex.Keys, page.Url)

		// Marshal the updated index
		indexBytes, err := indexPageMarshal(&pageIndex)
		if err != nil {
			return fmt.Errorf("error marshaling index: %w", err)
		}

		// Set the updated index in the database
		if err := txn.Set(indexKey, indexBytes); err != nil {
			return fmt.Errorf("error setting index data: %w", err)
		}

		return nil
	})
}

// ReadPage recupera uma página do banco de dados por URL
func (d Database) ReadPage(url string) (*data.Page, error) {
	var page data.Page
	Key := []byte(fmt.Sprintf("%s:%s", config.PageDataPrefix, url))

	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(Key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return pageUnmarshal(val, &page)
		})
	})
	if errors.Is(badger.ErrKeyNotFound, err) {
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
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataPrefix))

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
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataPrefix))

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

func (d Database) GetAllPage() ([]data.Page, error) {
	var pages []data.Page
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataPrefix))

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
func (d Database) GetPaginatedPage(pageNumber, pageSize int) ([]data.Page, error) {
	var pages []data.Page
	keyIndex := []byte(config.PageDataIndexName)

	err := d.db.View(func(txn *badger.Txn) error {
		index, err := txn.Get(keyIndex)
		if err != nil {
			return err
		}
		var pageIndex data.PageIndex
		if err := index.Value(func(val []byte) error {
			return indexPageUnmarshal(val, &pageIndex)
		}); err != nil {
			return err
		}

		start := pageNumber * pageSize
		end := start + pageSize
		if end > len(pageIndex.Keys) {
			end = len(pageIndex.Keys)
		}
		for _, key := range pageIndex.Keys[start:end] {
			item, err := txn.Get([]byte(fmt.Sprintf("%s:%s", config.PageDataPrefix, key)))
			if err != nil {
				return err
			}
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
	prefixKey := []byte(fmt.Sprintf("%s:", config.PageDataPrefix))

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

func (d Database) ImportData(pages []data.Page) error {
	// TODO: checar se o dado existe no banco de dados antes de importar, um WriteCheckPage
	var errs []error
	for _, page := range pages {
		err := d.WritePage(&page)
		if err != nil {
			errs = append(errs, err)
		}
		err = d.SetVisited(page.Url)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		log.Logger.Error("error writing page", zap.Errors("errors", errs))
		return fmt.Errorf("error writing page: %v", errs)
	}
	return nil
}

// GetStatistics recupera estatísticas do banco de dados
func (d Database) GetStatistics() (data.Statistic, error) {
	var statistic data.Statistic

	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(config.PageDataIndexName))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			var pageIndex data.PageIndex
			if err := indexPageUnmarshal(val, &pageIndex); err != nil {
				return err
			}

			statistic.TotalPages = len(pageIndex.Keys)

			// Group and count pages by host
			statistic.TotalPagesPerHost = make(map[string]int)
			for _, link := range pageIndex.Keys {
				u, err := url.Parse(link)
				if err != nil {
					continue
				}

				host := helper.NormalizeURL(u.Scheme + "://" + u.Host)
				statistic.TotalPagesPerHost[host] = statistic.TotalPagesPerHost[host] + 1
			}

			return nil
		})
	})

	return statistic, err
}
