package main

import (
	"WebCrawlerGui/backend/infra/log"
	"context"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"go.uber.org/zap"

	"embed"
	"runtime"

	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"

	"WebCrawlerGui/backend/consts"
	"WebCrawlerGui/backend/services"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

var version = "0.0.2"
var AppName = "WebCrawlerGui"

func main() {
	err := log.InitLogger()
	if err != nil {
		panic(err)
	}

	// Create an instance of the app structure
	prefSvc := services.InitConfigService(AppName, version)
	srv := services.InitSystem()
	crw := services.Crawling(AppName)

	// menu
	appMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.AppMenu())
		appMenu.Append(menu.EditMenu())
		appMenu.Append(menu.WindowMenu())
	}
	windowWidth, windowHeight, maximised := prefSvc.GetWindowSize()
	windowStartState := options.Normal
	if maximised {
		windowStartState = options.Maximised
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:            AppName,
		Width:            windowWidth,
		Height:           windowHeight,
		MinWidth:         consts.MIN_WINDOW_WIDTH,
		MinHeight:        consts.MIN_WINDOW_HEIGHT,
		WindowStartState: windowStartState,
		DisableResize:    false,
		//Frameless:                runtime.GOOS != "darwin",
		Menu:                     appMenu,
		EnableDefaultContextMenu: true,
		StartHidden:              true,
		LogLevel:                 logger.DEBUG,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "e3984e08-28dc-4e3d-b70a-45e961589cdc",
			OnSecondInstanceLaunch: srv.OnSecondInstanceLaunch,
		},
		OnStartup: func(ctx context.Context) {
			srv.Start(ctx, version, AppName, icon)
			crw.Handle(ctx)
		},
		OnDomReady: func(ctx context.Context) {
			x, y := prefSvc.GetWindowPosition(ctx)
			runtime2.WindowSetPosition(ctx, x, y)
			runtime2.WindowShow(ctx)
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			x, y := runtime2.WindowGetPosition(ctx)
			prefSvc.SaveWindowPosition(x, y)
			crw.HandleClose()
			return false
		},
		Bind: []interface{}{
			srv,
			crw,
			prefSvc,
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   AppName + " " + version,
				Message: "A modern lightweight cross-platform Redis desktop client.\n\nCopyright © 2024",
				Icon:    icon,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: true,
		},
		Linux: &linux.Options{
			ProgramName:         AppName,
			Icon:                icon,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		log.Logger.Panic("Error running Wails", zap.Error(err))
	}
}
