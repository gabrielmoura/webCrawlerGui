package types

type Preferences struct {
	Behavior PreferencesBehavior `json:"behavior" yaml:"behavior"`
	General  PreferencesGeneral  `json:"general" yaml:"general"`
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
