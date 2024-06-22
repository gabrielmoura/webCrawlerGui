package services

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/crawler"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/types"
	"bufio"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

type CrawlerService struct {
	Ctx context.Context
}

func (c CrawlerService) Handle(ctx context.Context) {
	log.Logger.Info("Crawling service started")
	c.Ctx = ctx
	crawler.Mount(c.Ctx)
	go crawler.LoopQueue()
}
func (c CrawlerService) HandleClose() {
	log.Logger.Info("Crawling service stopped")
	err := db.DB.CloseDB()
	if err != nil {
		log.Logger.Error("Error closing database", zap.Error(err))
		return
	}
}

func (c CrawlerService) Start() types.JSResp {
	log.Logger.Info("Crawling service started")
	config.Conf.General.EnableProcessing = true
	err := config.Conf.SaveFileConfig()
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error saving configuration",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Crawling service started",
	}
}
func (c CrawlerService) Stop() types.JSResp {
	log.Logger.Info("Crawling service stopped")
	config.Conf.General.EnableProcessing = false
	err := config.Conf.SaveFileConfig()
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error saving configuration",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Crawling service stopped",
	}
}
func (c CrawlerService) AddToQueue(url string) types.JSResp {
	err := db.DB.AddToQueue(url, 0)
	if err != nil {
		log.Logger.Error("Error adding URL to queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error adding URL to queue",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "URL added to queue",
	}
}

// RemoveFromQueueByHost remove URLs da fila por host
func (c CrawlerService) RemoveFromQueueByHost(host string) types.JSResp {
	link, err := url.Parse(host)
	if err != nil {
		log.Logger.Error("Error removing URL from queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error removing URL from queue",
		}
	}
	err = db.DB.DeletePrefix(fmt.Sprintf("%s://%s", link.Scheme, link.Host))
	if err != nil {
		log.Logger.Error("Error removing URL from queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error removing URL from queue",
		}
	}

	return types.JSResp{
		Success: true,
		Msg:     "URL removed from queue",
	}
}

// AddHotsTxt adiciona um arquivo Hosts.txt à fila, útil para adicionar uma lista de URLs
func (c CrawlerService) AddHotsTxt(url string) types.JSResp {
	if !config.Conf.General.ProxyEnabled {
		return types.JSResp{
			Success: false,
			Msg:     "Proxy is not enabled",
		}
	}
	resp, err := crawler.HttpRequest(url)
	defer resp.Body.Close()
	if err != nil {
		log.Logger.Error("Error getting Hosts.txt", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting Hosts.txt",
		}
	}
	if resp.StatusCode != http.StatusOK {
		log.Logger.Error("Error getting Hosts.txt", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting Hosts.txt",
		}
	}
	scanner := bufio.NewScanner(resp.Body)

	var errorsToQueue []error
	for scanner.Scan() {
		err = db.DB.AddToQueue(scanner.Text(), 0)
		if err != nil {
			errorsToQueue = append(errorsToQueue, err)
		}
	}

	if len(errorsToQueue) > 0 {
		log.Logger.Error("Error adding Hosts.txt to queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error adding Hosts.txt to queue",
		}
	}

	if err := scanner.Err(); err != nil {
		log.Logger.Error("Error reading Hosts.txt", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error reading Hosts.txt",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Hosts.txt added to queue",
	}
}

func (c CrawlerService) GetPaginatedQueue(pag types.Paginated) types.JSResp {
	queue, err := db.DB.ReadPaginated(pag.Limit, pag.Offset)
	if err != nil {
		log.Logger.Error("Error getting queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting queue",
		}
	}
	if len(queue) == 0 {
		return types.JSResp{
			Success: true,
			Msg:     "There is no data to display",
			Data:    []string{},
		}
	}
	return types.JSResp{
		Success: true,
		Data:    queue,
	}
}

func (c CrawlerService) GetAllQueue() types.JSResp {
	queue, err := db.DB.Read()
	if err != nil {
		log.Logger.Error("Error getting queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting queue",
		}
	}
	if len(queue) == 0 {
		return types.JSResp{
			Success: true,
			Msg:     "There is no data to display",
			Data:    []string{},
		}
	}
	return types.JSResp{
		Success: true,
		Data:    queue,
	}
}
func (c CrawlerService) DeleteQueue(url string) types.JSResp {
	err := db.DB.Delete(url)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error deleting queue",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Queue deleted",
	}
}

// Search pesquisa páginas por título, descrição ou conteúdo
func (c CrawlerService) Search(args string) types.JSResp {
	queue, err := db.DB.Search(args)
	if err != nil {
		log.Logger.Error("Error searching queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error searching queue",
		}
	}
	if len(queue) == 0 {
		return types.JSResp{
			Success: true,
			Msg:     "There is no data to display",
		}
	}
	return types.JSResp{
		Success: true,
		Data:    queue,
	}
}
func (c CrawlerService) SearchWords(args []string) types.JSResp {
	queue, err := db.DB.SearchV2(args)
	if err != nil {
		log.Logger.Error("Error searching queue", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error searching queue",
		}
	}
	if len(queue) == 0 {
		return types.JSResp{
			Success: true,
			Msg:     "There is no data to display",
		}
	}
	return types.JSResp{
		Success: true,
		Data:    queue,
	}
}

func Crawling(appName string) *CrawlerService {
	db.InitDB(appName)
	return &CrawlerService{}
}
