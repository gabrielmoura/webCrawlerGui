package types

import "WebCrawlerGui/backend/consts"

type Preferences struct {
	Behavior PreferencesBehavior  `json:"behavior" yaml:"behavior"`
	General  PreferencesGeneral   `json:"general" yaml:"general"`
	Editor   PreferencesEditor    `json:"editor" yaml:"editor"`
	Cli      PreferencesCli       `json:"cli" yaml:"cli"`
	Decoder  []PreferencesDecoder `json:"decoder" yaml:"decoder,omitempty"`
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

			TimeFormat:     "02-Jan-2006",
			TimeZone:       "America/Sao_Paulo",
			MaxConcurrency: 10,
			MaxDepth:       2,

			UserAgent: "Go-http-client/1.1",
			Cache: &CacheConfig{
				DBDir: "/tmp/WebCrawler",
				Mode:  "mem",
			},
			Proxy: &Proxy{
				Enabled:  false,
				ProxyURL: "",
			},
			Filter: &Filter{
				Tlds:        []string{},
				IgnoreLocal: false,
			},
		},
		Editor: PreferencesEditor{
			FontSize:    consts.DEFAULT_FONT_SIZE,
			ShowLineNum: true,
			ShowFolding: true,
			DropText:    true,
			Links:       true,
		},
		Cli: PreferencesCli{
			FontSize:    consts.DEFAULT_FONT_SIZE,
			CursorStyle: "block",
		},
		Decoder: []PreferencesDecoder{},
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

//type PreferencesGeneral struct {
//	Theme           string   `json:"theme" yaml:"theme"`
//	Language        string   `json:"language" yaml:"language"`
//	Font            string   `json:"font" yaml:"font,omitempty"`
//	FontFamily      []string `json:"fontFamily" yaml:"font_family,omitempty"`
//	FontSize        int      `json:"fontSize" yaml:"font_size"`
//	ScanSize        int      `json:"scanSize" yaml:"scan_size"`
//	KeyIconStyle    int      `json:"keyIconStyle" yaml:"key_icon_style"`
//	UseSysProxy     bool     `json:"useSysProxy" yaml:"use_sys_proxy,omitempty"`
//	UseSysProxyHttp bool     `json:"useSysProxyHttp" yaml:"use_sys_proxy_http,omitempty"`
//	CheckUpdate     bool     `json:"checkUpdate" yaml:"check_update"`
//	SkipVersion     string   `json:"skipVersion" yaml:"skip_version,omitempty"`
//	AllowTrack      bool     `json:"allowTrack" yaml:"allow_track"`
//}

type PreferencesEditor struct {
	Font        string   `json:"font" yaml:"font,omitempty"`
	FontFamily  []string `json:"fontFamily" yaml:"font_family,omitempty"`
	FontSize    int      `json:"fontSize" yaml:"font_size"`
	ShowLineNum bool     `json:"showLineNum" yaml:"show_line_num"`
	ShowFolding bool     `json:"showFolding" yaml:"show_folding"`
	DropText    bool     `json:"dropText" yaml:"drop_text"`
	Links       bool     `json:"links" yaml:"links"`
}

type PreferencesCli struct {
	FontFamily  []string `json:"fontFamily" yaml:"font_family,omitempty"`
	FontSize    int      `json:"fontSize" yaml:"font_size"`
	CursorStyle string   `json:"cursorStyle" yaml:"cursor_style,omitempty"`
}

type PreferencesDecoder struct {
	Name       string   `json:"name" yaml:"name"`
	Enable     bool     `json:"enable" yaml:"enable"`
	Auto       bool     `json:"auto" yaml:"auto"`
	DecodePath string   `json:"decodePath" yaml:"decode_path"`
	DecodeArgs []string `json:"decodeArgs" yaml:"decode_args,omitempty"`
	EncodePath string   `json:"encodePath" yaml:"encode_path"`
	EncodeArgs []string `json:"encodeArgs" yaml:"encode_args,omitempty"`
}

type PreferencesGeneral struct {
	Theme       string   `json:"theme" yaml:"theme"`
	Language    string   `json:"language" yaml:"language"`
	Font        string   `json:"font" yaml:"font,omitempty"`
	FontFamily  []string `json:"fontFamily" yaml:"font_family,omitempty"`
	FontSize    int      `json:"fontSize" yaml:"font_size"`
	CheckUpdate bool     `json:"checkUpdate" yaml:"check_update"`
	ScanSize    int      `json:"scanSize" yaml:"scan_size"`

	MaxConcurrency int          `mapstructure:"MAX_CONCURRENCY" json:"maxConcurrency" yaml:"maxConcurrency"`
	MaxDepth       int          `mapstructure:"MAX_DEPTH" json:"maxDepth" yaml:"maxDepth"`
	AppName        string       `mapstructure:"APP_NAME" json:"appName" json:"appName"`
	TimeFormat     string       `mapstructure:"TIME_FORMAT" json:"timeFormat" yaml:"timeFormat"`
	TimeZone       string       `mapstructure:"TIME_ZONE" json:"timeZone" yaml:"timeZone"`
	Cache          *CacheConfig `mapstructure:"CACHE" json:"cache" yaml:"cache"`
	Proxy          *Proxy       `mapstructure:"PROXY" json:"proxy" yaml:"proxy"`
	Filter         *Filter      `mapstructure:"FILTER" json:"filter" yaml:"filter"`
	UserAgent      string       `mapstructure:"USER_AGENT" json:"userAgent" yaml:"user_agent"`
}
type CacheConfig struct {
	DBDir string `mapstructure:"DB_DIR" json:"DBDir" yaml:"DBDir"`
	Mode  string `mapstructure:"MODE" json:"Mode" yaml:"Mode"` // "mem" or "disc'
}
type Proxy struct {
	Enabled  bool   `mapstructure:"ENABLED" json:"enabled" yaml:"enabled"`
	ProxyURL string `mapstructure:"PROXY_URL" json:"proxyURL" yaml:"proxyURL"`
}
type Filter struct {
	Tlds        []string `mapstructure:"TLDS" json:"tlds" yaml:"tlds"`
	IgnoreLocal bool     `mapstructure:"IGNORE_LOCAL" json:"ignoreLocal" yaml:"ignoreLocal"`
}
