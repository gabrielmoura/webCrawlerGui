package crawler

import (
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
	"go.uber.org/zap"
	"sync"
)

var visitedMutex sync.Mutex

//var pagesMutex sync.Mutex

// SetVisited adds a URL to the cache to mark it as visited.
func SetVisited(url string) {
	visitedMutex.Lock()
	err := db.DB.SetVisited(url)
	if err != nil {
		log.Logger.Error("error setting visited", zap.Error(err))
	}
	visitedMutex.Unlock()
}

// GetVisited retrieves a URL from the cache to check if it has been visited.
func GetVisited(url string) bool {
	visitedMutex.Lock()
	defer visitedMutex.Unlock()
	return db.DB.IsVisited(url)
}

// SetPage adds a page to the database.
func SetPage(page *data.Page) {
	err := db.DB.WritePage(page)
	if err != nil {
		return
	}
}

// GetPage retrieves a page from the database.
func GetPage(url string) *data.Page {
	p, err := db.DB.ReadPage(url)
	if err != nil {
		return nil
	}
	return p
}
