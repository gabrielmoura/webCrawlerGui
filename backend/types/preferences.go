package types

import "WebCrawlerGui/backend/consts"

type Preferences struct {
	Behavior PreferencesBehavior `json:"behavior" yaml:"behavior"`
	General  PreferencesGeneral  `json:"general" yaml:"general"`
}

func NewPreferences() Preferences {
	return Preferences{
		Behavior: PreferencesBehavior{
			AsideWidth:   consts.DEFAULT_ASIDE_WIDTH,
			WindowWidth:  consts.DEFAULT_WINDOW_WIDTH,
			WindowHeight: consts.DEFAULT_WINDOW_HEIGHT,
		},
		General: PreferencesGeneral{
			AppName:  "WebCrawler",
			Theme:    "auto",
			Language: "auto",
			FontSize: consts.DEFAULT_FONT_SIZE,
			ScanSize: consts.DEFAULT_SCAN_SIZE,
			//KeyIconStyle: 0,
			CheckUpdate: true,
			//AllowTrack:   true,

			TimeFormat:       "02-Jan-2006",
			TimeZone:         "America/Sao_Paulo",
			MaxConcurrency:   10,
			MaxDepth:         2,
			EnableProcessing: false,

			UserAgent:    "Go-http-client/1.1",
			ProxyEnabled: false,
			ProxyURL:     "",
			Tlds:         []string{},
			IgnoreLocal:  false,
		},
	}
}

type PreferencesBehavior struct {
	Welcomed        bool `json:"welcomed" yaml:"welcomed"`
	AsideWidth      int  `json:"asideWidth" yaml:"aside_width"`
	WindowWidth     int  `json:"windowWidth" yaml:"window_width"`
	WindowHeight    int  `json:"windowHeight" yaml:"window_height"`
	WindowMaximised bool `json:"windowMaximised" yaml:"window_maximised"`
	WindowPosX      int  `json:"windowPosX" yaml:"window_pos_x"`
	WindowPosY      int  `json:"windowPosY" yaml:"window_pos_y"`
	DarkMode        bool `json:"darkMode" yaml:"dark_mode"`
}

type PreferencesGeneral struct {
	Theme       string   `json:"theme" yaml:"theme"`
	Language    string   `json:"language" yaml:"language"`
	Font        string   `json:"font" yaml:"font,omitempty"`
	FontFamily  []string `json:"fontFamily" yaml:"font_family,omitempty"`
	FontSize    int      `json:"fontSize" yaml:"font_size"`
	CheckUpdate bool     `json:"checkUpdate" yaml:"check_update"`
	ScanSize    int      `json:"scanSize" yaml:"scan_size"`

	MaxConcurrency int    `mapstructure:"MAX_CONCURRENCY" json:"maxConcurrency" yaml:"maxConcurrency"`
	MaxDepth       int    `mapstructure:"MAX_DEPTH" json:"maxDepth" yaml:"maxDepth"`
	AppName        string `mapstructure:"APP_NAME" json:"appName" json:"appName"`
	TimeFormat     string `mapstructure:"TIME_FORMAT" json:"timeFormat" yaml:"timeFormat"`
	TimeZone       string `mapstructure:"TIME_ZONE" json:"timeZone" yaml:"timeZone"`

	Tlds        []string `mapstructure:"TLDS" json:"tlds" yaml:"tlds"`
	IgnoreLocal bool     `mapstructure:"IGNORE_LOCAL" json:"ignoreLocal" yaml:"ignoreLocal"`

	ProxyEnabled bool   `mapstructure:"PROXY_ENABLED" json:"proxyEnabled" yaml:"proxyEnabled"`
	ProxyURL     string `mapstructure:"PROXY_URL" json:"proxyURL" yaml:"proxyURL"`

	UserAgent string `mapstructure:"USER_AGENT" json:"userAgent" yaml:"user_agent"`

	EnableProcessing bool `mapstructure:"ENABLE_PROCESSING" json:"enableProcessing" yaml:"enableProcessing"`
}
