package db

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"strconv"
)

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
			err := item.Value(func(val []byte) error {
				depth, _ = strconv.Atoi(string(val))
				return nil
			})
			if err != nil {
				return err
			}
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
func (d Database) ReadPaginated(limit, offset int) ([]data.QueueType, error) {
	var urls []data.QueueType
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()
		prefix := []byte(config.QueueName)
		count := 0
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			if count >= offset && count < offset+limit {
				item := it.Item()
				url := string(item.KeyCopy(nil))
				depth := 0
				item.Value(func(val []byte) error {
					depth, _ = strconv.Atoi(string(val))
					return nil
				})
				urls = append(urls, data.QueueType{Url: url[len(config.QueueName)+1:], Depth: depth})
			}
			count++
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error reading from queue: %v", err)
	}
	return urls, nil
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
