package crawler

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
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
	Ctx context.Context
}

var Cs CrawlerService

func Mount(
	Ctx context.Context) {

	Cs = CrawlerService{
		Ctx: Ctx,
	}
}

// processPage processa uma página, extrai links e dados
func processPage(pageUrl string, depth int) {

	log.Logger.Debug(fmt.Sprintf("Looping queue, depth: %d", depth))
	if depth > config.Conf.General.MaxDepth {
		log.Logger.Info(fmt.Sprintf("Reached max depth of %d, %d", config.Conf.General.MaxDepth, depth))
		return
	}
	// Só processa uma página se ela ainda não foi visitada

	if GetVisited(pageUrl) {
		return
	}

	log.Logger.Info(fmt.Sprintf("Visiting %s", pageUrl))
	plainText, htmlDoc, err := visitLink(pageUrl)
	if err != nil {
		if errors.Is(err, mimeNotAllow) {
			//log.Logger.Info(fmt.Sprintf("MIME not allowed: %s", pageUrl))
			return
		}
		log.Logger.Debug(fmt.Sprintf("Error checking link: %s", err))
		return
	}

	links, err := extractLinks(pageUrl, htmlDoc)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Error extracting links: %s", err))
		return
	}

	dataPage, err := extractData(htmlDoc)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Error extracting data: %s", err))
		return
	}
	words, _ := countWordsInText(plainText)

	dataPage.Words = words
	dataPage.Url = pageUrl
	dataPage.Links = links
	dataPage.Timestamp = time.Now()
	dataPage.Visited = true

	SetPage(pageUrl, dataPage)

	SetVisited(pageUrl)

	handleAddToQueue(links, depth+1)
}
func HandleQueue(initialURL string) {
	// Só processa a fila se ela não estiver vazia
	log.Logger.Info("Handling queue")

	ok, _, err := db.DB.GetFromQueue()

	if err != nil { // Check if queue is empty
		log.Logger.Info("Queue is empty", zap.Error(err))
	}
	if ok == "" {
		db.DB.AddToQueue(initialURL, 0)
	}
	LoopQueue()
}

func LoopQueue() {
	for {
		if config.Conf.General.EnableProcessing {
			links, err := db.DB.GetFromQueueV2(config.Conf.General.MaxConcurrency) // Get a batch of links
			if err != nil {
				log.Logger.Error("Error getting from queue", zap.Error(err))
			}
			if len(links) == 0 {
				log.Logger.Info("Queue is empty", zap.Any("Links", links))
				break
			}

			for _, link := range links {
				if link.Url == "" {
					continue
				}
				Wg.Add(1) // Para cada link, incrementa o WaitGroup

				go func(link data.QueueType) {
					defer Wg.Done()
					processPage(link.Url, link.Depth)
				}(link)
			}
			Wg.Wait()
		}
		time.Sleep(1 * time.Second)
		log.Logger.Debug("Looping queue")

	}
}

func visitLink(pageUrl string) ([]byte, *html.Node, error) {
	resp, err := httpRequest(pageUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching URL %s: %w", pageUrl, err)
	}
	defer resp.Body.Close()

	if isStatusErr(resp.StatusCode, resp.Request.URL) {
		// TODO: Implementar lógica para por em outra fila e ternar novamente
		log.Logger.Info("Status Error", zap.String("URL", pageUrl), zap.String("Status", resp.Status))
		return nil, nil, ErrUnexpectedStatus
	}

	// Streamlined MIME type check and early return
	if !isAllowedMIME(resp.Header.Get("Content-Type"), config.AcceptableMimeTypes) {
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
