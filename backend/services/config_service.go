package services

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/consts"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/types"
	"context"
	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"
	"go.uber.org/zap"
)

type ConfigService struct {
	cref *config.Config
}

func InitConfigService(appName string, appVersion string) *ConfigService {
	return &ConfigService{cref: config.InitConfig(appName, appVersion)}
}
func (c *ConfigService) SaveWindowPosition(x, y int) {
	if x > 0 || y > 0 {
		c.cref.Behavior.WindowPosX, c.cref.Behavior.WindowPosY = x, y
		err := c.cref.SaveFileConfig()
		if err != nil {
			log.Logger.Error("Error saving window position", zap.Error(err))
		}
	}
}
func (c *ConfigService) GetWindowPosition(ctx context.Context) (x, y int) {
	x, y = c.cref.Behavior.WindowPosX, c.cref.Behavior.WindowPosY
	width, height := c.cref.Behavior.WindowWidth, c.cref.Behavior.WindowHeight
	var screenWidth, screenHeight int

	if screens, err := runtime2.ScreenGetAll(ctx); err == nil {
		for _, screen := range screens {
			if screen.IsCurrent {
				screenWidth, screenHeight = screen.Size.Width, screen.Size.Height
				break
			}
		}
	}
	if screenWidth <= 0 || screenHeight <= 0 {
		screenWidth, screenHeight = consts.DEFAULT_WINDOW_WIDTH, consts.DEFAULT_WINDOW_HEIGHT
	}
	if x <= 0 || x+width > screenWidth || y <= 0 || y+height > screenHeight {
		// out of screen, reset to center
		x, y = (screenWidth-width)/2, (screenHeight-height)/2
	}
	return
}

func (c *ConfigService) GetWindowSize() (width, height int, maximised bool) {
	width, height, maximised = c.cref.Behavior.WindowWidth, c.cref.Behavior.WindowHeight, c.cref.Behavior.WindowMaximised
	if width <= 0 {
		width = consts.DEFAULT_WINDOW_WIDTH
	}
	if height <= 0 {
		height = consts.DEFAULT_WINDOW_HEIGHT
	}
	return
}
func (c *ConfigService) SaveWindowSize(width, height int, maximised bool) {
	c.cref.Behavior.WindowWidth, c.cref.Behavior.WindowHeight, c.cref.Behavior.WindowMaximised = width, height, maximised
	err := c.cref.SaveFileConfig()
	if err != nil {
		log.Logger.Error("Error saving window size", zap.Error(err))
	}
}
func (c *ConfigService) GetPreferences() (resp types.JSResp) {
	err := c.cref.LoadConfig()
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     err.Error(),
		}
	}
	resp.Success = true
	resp.Data = types.Preferences{
		Behavior: c.cref.Behavior,
		General:  c.cref.General,
	}
	return
}
func (c *ConfigService) SetPreferences(pf types.Preferences) (resp types.JSResp) {
	c.cref.Behavior = pf.Behavior
	c.cref.General = pf.General
	err := c.cref.SaveFileConfig()
	if err != nil {
		resp.Success = false
		resp.Msg = "Error saving preferences"
		log.Logger.Error("Error saving preferences", zap.Error(err))
		return
	}
	resp.Success = true
	resp.Msg = "Preferences saved"
	return
}
