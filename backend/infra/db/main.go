package db

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// blockWrite é mutex para controle de otimização dos logs
	blockWrite sync.RWMutex
)

type Database struct {
	db     *badger.DB
	logger *zap.Logger
}

// InitDB inicializa a sessão do banco de dados
func InitDB(path string, logger *zap.Logger) *Database {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil // Desativa logs de debug
	opts.CompactL0OnClose = true
	opts.NumCompactors = 2
	opts.ValueLogFileSize = 100 << 20 // 100 MB

	db, err := badger.Open(opts)
	if err != nil {
		logger.Fatal("error initializing database", zap.Error(err))
	}
	return &Database{
		db:     db,
		logger: logger,
	}
}

// CloseDB fecha a sessão do banco de dados
func (d Database) CloseDB() error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	return d.db.Close()
}
func (d Database) SyncCache() error {
	time.Sleep(1 * time.Second)
	d.logger.Info("Syncing cache")

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
func (d Database) OptimizeCache() {
	for {
		time.Sleep(2 * time.Minute)
		d.logger.Info("Optimizing cache")
		blockWrite.Lock()
		err := d.db.RunValueLogGC(0.9)
		if err != nil && !errors.Is(badger.ErrNoRewrite, err) {
			d.logger.Info("error optimizing cache", zap.Error(err))
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
	return d.db.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(page)
		if err != nil {
			return err
		}
		return txn.Set([]byte(page.Url), data)
	})
}

// ReadPage recupera uma página do banco de dados por URL
func (d Database) ReadPage(url string) (*data.Page, error) {
	var page data.Page
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(url))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &page)
		})
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	return &page, err
}

// IsVisited verifica se uma URL foi visitada
//func (d Database) IsVisited(url string) bool {
//	var page data.Page
//	err := d.db.View(func(txn *badger.Txn) error {
//		item, err := txn.Get([]byte(url))
//		if err != nil {
//			return err
//		}
//		return item.Value(func(val []byte) error {
//			return json.Unmarshal(val, &page)
//		})
//	})
//	return err == nil && page.Visited
//}

// AllVisited recupera todos os URLs visitados
func (d Database) AllVisited() ([]string, error) {
	var urls []string
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var page data.Page
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &page)
			})
			if err != nil {
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
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var page data.Page
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &page)
			})
			if err != nil {
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

// containsIgnoreCase verifica se um texto contém outro ignorando maiúsculas e minúsculas
func containsIgnoreCase(text, substr string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(substr))
}

// SearchByContent pesquisa páginas por conteúdo e ordena por frequência
func (d Database) SearchByContent(searchTerm string) ([]data.PageSearchWithFrequency, error) {
	var pages []data.PageSearchWithFrequency
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var page data.Page
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &page)
			})
			if err != nil {
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

	// Ordenar por frequência
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Frequency > pages[j].Frequency
	})

	return pages, nil
}

// Search pesquisa páginas por título, descrição ou conteúdo
func (d Database) Search(searchTerm string) ([]data.PageSearch, error) {
	var pages []data.PageSearch
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var page data.Page
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &page)
			})
			if err != nil {
				return err
			}
			if containsIgnoreCase(page.Title, searchTerm) ||
				containsIgnoreCase(page.Description, searchTerm) ||
				func() bool {
					_, ok := page.Words[searchTerm]
					return ok
				}() {
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

// FILAS

type Queue interface {
	Enqueue(url string, depth int) error
	Dequeue() (string, int, error)

	Read() ([]data.QueueType, error)
	Delete(url string) error
}

// Enqueue adds a URL to the queue.
func (d Database) Enqueue(url string, depth int) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.QueueName, url))
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, []byte(strconv.Itoa(depth)))
	})
}

// Dequeue retrieves and removes a URL from the queue.
func (d Database) Dequeue() (string, int, error) {
	blockWrite.RLock()
	defer blockWrite.RUnlock()

	var url string
	var depth int

	err := d.db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false // We don't need the values
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(config.QueueName)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			url = string(item.KeyCopy(nil)) // Copy the key to avoid issues
			item.Value(func(val []byte) error {
				depth, _ = strconv.Atoi(string(val))
				return nil
			})
			return txn.Delete(item.Key()) // Remove from queue after retrieval
		}
		return nil // No items in queue
	})

	if err != nil {
		return "", 0, fmt.Errorf("error dequeuing from queue: %v", err)
	}

	if url != "" {
		url = url[len(config.QueueName)+1:]
	}

	return url, depth, nil
}

// Read retrieves all URLs from the queue.
func (d Database) Read() ([]data.QueueType, error) {
	var urls []data.QueueType

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(config.QueueName)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			url := string(item.KeyCopy(nil))
			depth := 0
			item.Value(func(val []byte) error {
				depth, _ = strconv.Atoi(string(val))
				return nil
			})
			urls = append(urls, data.QueueType{Url: url[len(config.QueueName)+1:], Depth: depth})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error reading from queue: %v", err)
	}

	return urls, nil
}

// Delete removes a URL from the queue.
func (d Database) Delete(url string) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.QueueName, url))
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}
