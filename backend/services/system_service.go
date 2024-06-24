package services

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/consts"
	data2 "WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/types"
	"context"
	"encoding/json"
	"fmt"
	"fyne.io/systray"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.uber.org/zap"
	"net/http"
	"os"
	runtime2 "runtime"
	"strings"
	"sync"
	"time"
)

type SystemService struct {
	ctx        context.Context
	appVersion string
	appName    string
}
type latestRelease struct {
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

var System *SystemService
var onceSystem sync.Once

func InitSystem() *SystemService {
	if System == nil {
		onceSystem.Do(func() {
			System = &SystemService{
				appVersion: "0.0.0",
			}
			go System.loopWindowEvent()
		})
	}
	return System
}
func onReadySystray(ctx context.Context, appIcon []byte) func() {
	return func() {
		systray.SetTemplateIcon(appIcon, appIcon)
		systray.SetTitle("Web Crawler")
		systray.SetTooltip("Web Crawler")
		mStart := systray.AddMenuItem("Iniciar", "Iniciar a aplicação")
		mPause := systray.AddMenuItem("Pausar", "Pausar a aplicação")
		systray.AddSeparator()
		mRestore := systray.AddMenuItem("Restaurar", "Restaurar configurações padrão")
		mQuit := systray.AddMenuItem("Fechar", "Fechar a aplicação")
		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					runtime.Quit(ctx)
					systray.Quit()
				case <-mPause.ClickedCh:
					runtime.EventsEmit(ctx, "pause")
				case <-mRestore.ClickedCh:
					runtime.Show(ctx)
				case <-mStart.ClickedCh:
					runtime.EventsEmit(ctx, "start")
				}
			}
		}()
	}
}
func onExitSystray() {
	fmt.Println("Exiting...")
}

func (s *SystemService) Start(ctx context.Context, version string, appName string, appIcon []byte) {
	s.ctx = ctx
	s.appVersion = version
	s.appName = appName

	// maximize the window if screen size is lower than the minimum window size
	if screen, err := runtime.ScreenGetAll(ctx); err == nil && len(screen) > 0 {
		for _, sc := range screen {
			if sc.IsCurrent {
				if sc.Size.Width < consts.MIN_WINDOW_WIDTH || sc.Size.Height < consts.MIN_WINDOW_HEIGHT {
					runtime.WindowMaximise(ctx)
					break
				}
			}
		}
	}
	runtime.EventsOn(s.ctx, "start", func(optionalData ...interface{}) {
		config.Conf.General.EnableProcessing = true
		err := config.Conf.SaveFileConfig()
		if err != nil {
			log.Logger.Error("Error saving config", zap.Error(err))
		}
	})
	runtime.EventsOn(s.ctx, "pause", func(optionalData ...interface{}) {
		config.Conf.General.EnableProcessing = false
		err := config.Conf.SaveFileConfig()
		if err != nil {
			log.Logger.Error("Error saving config", zap.Error(err))
		}
	})

	go systray.Run(onReadySystray(ctx, appIcon), onExitSystray)
}

func (s *SystemService) Info() (resp types.JSResp) {
	resp.Success = true
	resp.Data = struct {
		OS      string `json:"os"`
		Arch    string `json:"arch"`
		Version string `json:"version"`
	}{
		OS:      runtime2.GOOS,
		Arch:    runtime2.GOARCH,
		Version: s.appVersion,
	}
	return
}
func (s *SystemService) loopWindowEvent() {
	var fullscreen, maximised, minimised, normal bool
	var width, height int
	var dirty bool
	for {
		time.Sleep(300 * time.Millisecond)
		if s.ctx == nil {
			continue
		}

		dirty = false
		if f := runtime.WindowIsFullscreen(s.ctx); f != fullscreen {
			// full-screen switched
			fullscreen = f
			dirty = true
		}

		if w, h := runtime.WindowGetSize(s.ctx); w != width || h != height {
			// window size changed
			width, height = w, h
			dirty = true
		}

		if m := runtime.WindowIsMaximised(s.ctx); m != maximised {
			maximised = m
			dirty = true
		}

		if m := runtime.WindowIsMinimised(s.ctx); m != minimised {
			minimised = m
			dirty = true
		}

		if n := runtime.WindowIsNormal(s.ctx); n != normal {
			normal = n
			dirty = true
		}

		if dirty {
			runtime.EventsEmit(s.ctx, "window_changed", map[string]any{
				"fullscreen": fullscreen,
				"width":      width,
				"height":     height,
				"maximised":  maximised,
				"minimised":  minimised,
				"normal":     normal,
			})

			if !fullscreen && !minimised {
				// save window size and position
				b := config.Conf.Behavior
				b.WindowWidth, b.WindowHeight, b.WindowMaximised = width, height, maximised
				err := config.Conf.SaveFileConfig()
				if err != nil {
					log.Logger.Error("Error saving window size", zap.Error(err))
				}
			}
		}
	}
}

func (s *SystemService) OnSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	secondInstanceArgs := secondInstanceData.Args

	println("user opened second instance", strings.Join(secondInstanceData.Args, ","))
	println("user opened second from", secondInstanceData.WorkingDirectory)
	runtime.WindowUnminimise(s.ctx)
	runtime.Show(s.ctx)
	go runtime.EventsEmit(s.ctx, "launchArgs", secondInstanceArgs)
}

func (s *SystemService) CheckForUpdate() (resp types.JSResp) {
	// request latest version
	res, err := http.Get(fmt.Sprintf("https://api.github.com/repos/gabrielmoura/%s/releases/latest", s.appName))
	if err != nil || res.StatusCode != http.StatusOK {
		resp.Msg = "network error"
		return
	}

	var respObj latestRelease
	err = json.NewDecoder(res.Body).Decode(&respObj)
	if err != nil {
		resp.Msg = "invalid content"
		return
	}

	// compare with current version
	resp.Success = true
	resp.Data = map[string]any{
		"version":  s.appVersion,
		"latest":   respObj.TagName,
		"page_url": respObj.HtmlUrl,
	}
	return
}

func (s *SystemService) ExportData() types.JSResp {
	saveDialogOptions := runtime.SaveDialogOptions{
		Title:           "Export data",
		DefaultFilename: "data.json",
		Filters:         []runtime.FileFilter{{DisplayName: "JSON Files", Pattern: "*.json"}},
	}
	filename, err := runtime.SaveFileDialog(s.ctx, saveDialogOptions)
	if err != nil {
		log.Logger.Error("Error saving file", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error saving file",
		}
	}

	data, err := db.DB.GetAllPage()
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error getting data",
		}
	}
	if len(data) == 0 {
		return types.JSResp{
			Success: false,
			Msg:     "No data to export",
		}
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error exporting data",
		}
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error exporting data",
		}
	}
	return types.JSResp{
		Success: true,
		Msg:     "Data exported",
	}
}
func (s *SystemService) ImportData() types.JSResp {
	openDialogOptions := runtime.OpenDialogOptions{
		Title: "Import data",
		Filters: []runtime.FileFilter{
			{DisplayName: "JSON Files", Pattern: "*.json"},
		},
	}
	filename, err := runtime.OpenFileDialog(s.ctx, openDialogOptions)
	if err != nil {
		log.Logger.Error("Error opening file", zap.Error(err))
		return types.JSResp{
			Success: false,
			Msg:     "Error opening file",
		}
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error reading file",
		}
	}
	var pages []data2.Page
	err = json.Unmarshal(data, &pages)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error parsing data",
		}
	}

	err = db.DB.ImportData(pages)
	if err != nil {
		return types.JSResp{
			Success: false,
			Msg:     "Error importing data",
		}
	}
	db.DB.OptimizeCacheNow()
	return types.JSResp{
		Success: true,
		Msg:     "Data imported",
	}
}
