package crawler

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/types"
	"bytes"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/html"
	"io"
	"sync"
	"time"
)

var Wg sync.WaitGroup
var (
	mimeNotAllow        = errors.New("mime: not allowed")
	ErrUnexpectedStatus = errors.New("unexpected status")
)

type CrawlerService struct {
	logger  *zap.Logger
	conf    types.PreferencesGeneral
	ctx     context.Context
	db      *db.Database
	enabled bool
}

func InitCrawlerService(ctx context.Context, logger *zap.Logger,
	conf types.PreferencesGeneral, db *db.Database) *CrawlerService {

	return &CrawlerService{
		logger: logger,
		conf:   conf,
		ctx:    ctx,
		db:     db,
	}
}

// processPage processa uma página, extrai links e dados
func (c CrawlerService) processPage(pageUrl string, depth int) {

	c.logger.Debug(fmt.Sprintf("Looping queue, depth: %d", depth))
	if depth > c.conf.MaxDepth {
		c.logger.Info(fmt.Sprintf("Reached max depth of %d, %d", c.conf.MaxDepth, depth))
		return
	}
	// Só processa uma página se ela ainda não foi visitada

	if c.GetVisited(pageUrl) {
		return
	}

	c.logger.Info(fmt.Sprintf("Visiting %s", pageUrl))
	plainText, htmlDoc, err := c.visitLink(pageUrl)
	if err != nil {
		if errors.Is(err, mimeNotAllow) {
			//log.Logger.Info(fmt.Sprintf("MIME not allowed: %s", pageUrl))
			return
		}
		c.logger.Debug(fmt.Sprintf("Error checking link: %s", err))
		return
	}

	links, err := c.extractLinks(pageUrl, htmlDoc)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error extracting links: %s", err))
		return
	}

	dataPage, err := c.extractData(htmlDoc)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error extracting data: %s", err))
		return
	}
	words, _ := c.countWordsInText(plainText)

	dataPage.Words = words
	dataPage.Url = pageUrl
	dataPage.Links = links
	dataPage.Timestamp = time.Now()
	dataPage.Visited = true

	c.SetPage(pageUrl, dataPage)

	c.SetVisited(pageUrl)

	c.handleAddToQueue(links, depth+1)
}

//func HandleQueue(initialURL string) {
//	// Só processa a fila se ela não estiver vazia
//	log.Logger.Info("Handling queue")
//	ok, _, err := cache.GetFromQueue()
//
//	if err != nil { // Check if queue is empty
//		log.Logger.Info("Queue is empty", zap.Error(err))
//	}
//	if ok == "" {
//		cache.AddToQueue(initialURL, 0)
//	}
//	LoopQueue()
//}

func (c CrawlerService) LoopQueue() {

	for {
		if c.enabled {
			if c.ctx.Done() != nil {
				break
			}
			links, notLinks := c.db.GetFromQueueV2(c.conf.MaxConcurrency) // Get a batch of links
			if len(links) == 0 || notLinks != nil {
				c.logger.Info("Queue is empty", zap.Error(notLinks))
				break
			}

			for _, link := range links {
				if link.Url == "" {
					continue
				}
				Wg.Add(1) // Para cada link, incrementa o WaitGroup

				go func(link data.QueueType) {
					defer Wg.Done()
					c.processPage(link.Url, link.Depth)
				}(link)
			}
			Wg.Wait()
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CrawlerService) visitLink(pageUrl string) ([]byte, *html.Node, error) {
	resp, err := c.httpRequest(pageUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching URL %s: %w", pageUrl, err)
	}
	defer resp.Body.Close()

	if c.isStatusErr(resp.StatusCode, resp.Request.URL) {
		// TODO: Implementar lógica para por em outra fila e ternar novamente
		c.logger.Info("Status Error", zap.String("URL", pageUrl), zap.String("Status", resp.Status))
		return nil, nil, ErrUnexpectedStatus
	}

	// Streamlined MIME type check and early return
	if !c.isAllowedMIME(resp.Header.Get("Content-Type"), config.AcceptableMimeTypes) {
		return nil, nil, mimeNotAllow
	}

	// Efficiently read the response body into a buffer
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse HTML from the buffered content
	htmlDoc, err := html.Parse(bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	return bodyBytes, htmlDoc, nil
}
