package services

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/crawler"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/types"
	"context"
	"go.uber.org/zap"
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

func Crawling(appName string) *CrawlerService {
	db.InitDB(appName)
	return &CrawlerService{}
}
