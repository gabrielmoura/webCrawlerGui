package services

import (
	"context"

	"WebCrawlerGui/backend/infra/crawler"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/storage"
	"WebCrawlerGui/backend/types"
	"go.uber.org/zap"
)

type CrawlingService struct {
	logger *zap.Logger
	conf   types.PreferencesGeneral
	db     *db.Database
}

func (s CrawlingService) Start(ctx context.Context, conf *PreferencesService) {
	s.logger.Info("Crawling service started")

	cfg := conf.pref
	s.conf = cfg.GetPreferences().General

	service := crawler.InitCrawlerService(ctx, s.logger, s.conf, s.db)

	service.LoopQueue()
}
func (s CrawlingService) AddToQueue(url string) types.JSResp {
	//crawler.AddToQueue(url)
	err := s.db.AddToQueue(url, 0)
	if err != nil {
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
func (s CrawlingService) GetAllQueue() types.JSResp {
	queue, err := s.db.Read()
	if err != nil {
		s.logger.Error("Error getting queue", zap.Error(err))
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
func (s CrawlingService) DeleteQueue(url string) types.JSResp {
	err := s.db.Delete(url)
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
func (s CrawlingService) Search(args string) types.JSResp {
	// queue, err := s.db.Search(args)
	// if err != nil {
	// 	s.logger.Error("Error searching queue", zap.Error(err))
	// 	return types.JSResp{
	// 		Success: false,
	// 		Msg:     "Error searching queue",
	// 	}
	// }
	// if len(queue) == 0 {
	// 	return types.JSResp{
	// 		Success: true,
	// 		Msg:     "There is no data to display",
	// 		Data:    []string{},
	// 	}
	// }
	return types.JSResp{
		Success: true,
		Msg:     "Not implemented",
	}
}
func Crawling(appName string) *CrawlingService {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	dbPath := storage.NewCacheStore(appName, "")

	database := db.InitDB(dbPath.ConfPath, logger)
	return &CrawlingService{
		logger: logger,
		db:     database,
	}
}
