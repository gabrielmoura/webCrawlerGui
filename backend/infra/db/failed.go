package db

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"fmt"
	"github.com/dgraph-io/badger/v4"
)

// AddToFailed adds a URL to the failed list.
func (d Database) AddToFailed(url string, reason string) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.PrefixFailedData, url))

	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, []byte(reason))
	})
}

// GetFailed retrieves a failed URL.
func (d Database) GetFailed(url string) error {
	key := []byte(fmt.Sprintf("%s:%s", config.PrefixFailedData, url))

	return d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return nil
		})
	})
}

// GetAllFailed retrieves all failed URLs.
func (d Database) GetAllFailed() ([]data.FailedType, error) {
	var urls []data.FailedType

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(config.PrefixFailedData)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			url := string(item.KeyCopy(nil))
			reason := ""
			err := item.Value(func(val []byte) error {
				reason = string(val)
				return nil
			})
			if err != nil {
				return err
			}
			urls = append(urls, data.FailedType{Url: url[len(config.PrefixFailedData)+1:], Reason: reason})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error reading from failed: %v", err)
	}

	return urls, nil
}

// DeleteFailed removes a URL from the failed list.
func (d Database) DeleteFailed(url string) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.PrefixFailedData, url))
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

// DeleteFailedPrefix removes all URLs with a given prefix from the failed list.
func (d Database) DeleteFailedPrefix(prefix string) error {
	blockWrite.RLock()
	defer blockWrite.RUnlock()
	key := []byte(fmt.Sprintf("%s:%s", config.PrefixFailedData, prefix))
	return d.db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(key); it.ValidForPrefix(key); it.Next() {
			item := it.Item()
			err := txn.Delete(item.Key())
			if err != nil {
				return err
			}
		}
		return nil
	})
}
