package crawler

import (
	"WebCrawlerGui/backend/infra/data"
	"sync"
)

var visitedMutex sync.Mutex
var pagesMutex sync.Mutex

// SetVisited adds a URL to the cache to mark it as visited.
func (c CrawlerService) SetVisited(url string) {
	visitedMutex.Lock()
	c.db.SetVisited(url)
	visitedMutex.Unlock()
}

// GetVisited retrieves a URL from the cache to check if it has been visited.
func (c CrawlerService) GetVisited(url string) bool {
	visitedMutex.Lock()
	defer visitedMutex.Unlock()
	return c.db.IsVisited(url)
}

// SetPage adds a page to the database.
func (c CrawlerService) SetPage(url string, page *data.Page) {
	//pagesMutex.Lock()
	err := c.db.WritePage(page)
	if err != nil {
		return
	}
	//pages[url] = page
	//pagesMutex.Unlock()
}

// GetPage retrieves a page from the database.
func (c CrawlerService) GetPage(url string) *data.Page {
	//pagesMutex.Lock()
	//defer pagesMutex.Unlock()
	p, err := c.db.ReadPage(url)
	if err != nil {
		return nil
	}
	return p
}
