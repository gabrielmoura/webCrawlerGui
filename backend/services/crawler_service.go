package services

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/helper"
	"WebCrawlerGui/backend/infra/crawler"
	"WebCrawlerGui/backend/infra/data"
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

func Crawling(appName string) *CrawlerService {
	db.InitDB(appName)
	return &CrawlerService{}
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
	err = db.DB.DeleteQueuePrefix(fmt.Sprintf("%s://%s", link.Scheme, link.Host))
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
	if err != nil {
		log.Logger.Error("Error getting Hosts.txt", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting Hosts.txt",
		}
	}
	defer resp.Body.Close()

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
	queue, err := db.DB.ReadQueuePaginated(pag.Limit, pag.Offset)
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
	queue, err := db.DB.GetAllQueue()
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
	err := db.DB.DeleteQueue(url)
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

// GetTreePages retorna uma árvore de páginas agrupadas por host
func (c CrawlerService) GetTreePages(pageN, size int) types.JSResp {
	pages, _ := db.DB.GetPaginatedPage(pageN, size)
	pageMap := make(map[string]*data.Page)
	hostMap := make(map[string]*types.TreeNode)
	visited := make(map[string]bool)

	// Construir um mapa para acesso rápido às páginas por URL e agrupar por host
	for i := range pages {
		page := &pages[i]
		pageMap[page.Url] = page

		u, err := url.Parse(page.Url)
		if err != nil {
			continue // Ignorar URLs inválidas
		}

		host := helper.NormalizeURL(u.Scheme + "://" + u.Host)

		if _, exists := hostMap[host]; !exists {
			// Incluir título e descrição do nó pai
			hostMap[host] = &types.TreeNode{
				Title:       page.Title,
				Description: page.Description,
				URL:         host,
				Children:    []types.TreeNode{},
			}
		}

		// Evitar adicionar a URL do nó pai aos filhos
		if helper.NormalizeURL(page.Url) != host {
			hostMap[host].Children = append(hostMap[host].Children, helper.PageToTreeNode(page, pageMap, host, visited))
		}
	}

	var treeNodes []types.TreeNode
	for _, node := range hostMap {
		treeNodes = append(treeNodes, *node)
	}

	return types.JSResp{
		Success: true,
		Data:    treeNodes,
	}
}

func (c CrawlerService) GetAllFailed() types.JSResp {
	failed, err := db.DB.GetAllFailed()
	if err != nil {
		log.Logger.Error("Error getting failed", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting failed",
		}
	}
	if len(failed) == 0 {
		return types.JSResp{
			Success: true,
			Msg:     "There is no data to display",
		}
	}
	return types.JSResp{
		Success: true,
		Data:    failed,
	}
}
func (c CrawlerService) DeleteFailed(url string) types.JSResp {
	err := db.DB.DeleteFailed(url)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error deleting failed",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Failed deleted",
	}
}
func (c CrawlerService) DeleteAllFailed(prefix string) types.JSResp {
	err := db.DB.DeleteFailedPrefix(prefix)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error deleting failed",
		}
	}
	log.Logger.Info("Failed deleted", zap.String("prefix", prefix))
	return types.JSResp{
		Success: true,
		Msg:     "Failed deleted",
	}
}

func (c CrawlerService) GetStatistics() types.JSResp {
	stats, err := db.DB.GetStatistics()
	if err != nil {
		log.Logger.Error("Error getting statistics", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error getting statistics",
		}
	}
	return types.JSResp{
		Success: true,
		Data:    stats,
	}
}
