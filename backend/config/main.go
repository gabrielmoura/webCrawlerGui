package config

import (
	"WebCrawlerGui/backend/consts"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/types"
	"github.com/vrischmann/userdir"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"sync"
)

var Conf *Config

type Config struct {
	sync.RWMutex `yaml:"-" json:"-"`
	AppName      string                    `json:"app_name" yaml:"app_name"`
	AppVersion   string                    `json:"app_version" yaml:"app_version"`
	Behavior     types.PreferencesBehavior `json:"behavior" yaml:"behavior"`
	General      types.PreferencesGeneral  `json:"general" yaml:"general"`
}

func InitConfig(appName string, appVersion string) *Config {
	cs := &Config{
		AppName:    appName,
		AppVersion: appVersion,
	}
	err := cs.LoadConfig()
	if err != nil {
		log.Logger.Error("Error loading config", zap.Error(err))
		panic(err)
	}
	Conf = cs
	return cs
}

func (c *Config) getPath(fileName string) string {
	return path.Join(userdir.GetConfigHome(), c.AppName, fileName)
}

func ensureDirExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) SaveFileConfig() error {
	c.Lock()
	defer c.Unlock()

	cData, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	dir := path.Dir(c.getPath(""))
	if err := ensureDirExists(dir); err != nil {
		return err
	}

	if err := os.WriteFile(c.getPath("config.yml"), cData, 0666); err != nil { // Alterei para 0666 para ser mais seguro
		return err
	}

	return nil
}

func (c *Config) DefaultConfig() {
	c.Behavior = types.PreferencesBehavior{
		WindowMaximised: true,
		AsideWidth:      consts.DEFAULT_ASIDE_WIDTH,
		WindowWidth:     consts.DEFAULT_WINDOW_WIDTH,
		WindowHeight:    consts.DEFAULT_WINDOW_HEIGHT,
	}
	c.General = types.PreferencesGeneral{
		AppName:           c.AppName,
		Theme:             "auto",
		Language:          "auto",
		FontSize:          consts.DEFAULT_FONT_SIZE,
		ScanSize:          consts.DEFAULT_SCAN_SIZE,
		CheckUpdate:       true,
		TimeFormat:        "02-Jan-2006",
		TimeZone:          "America/Sao_Paulo",
		MaxConcurrency:    10,
		MaxDepth:          2,
		EnableProcessing:  false,
		UserAgent:         "Go-http-client/1.1",
		ProxyEnabled:      false,
		ProxyURL:          "",
		Tlds:              []string{},
		IgnoreLocal:       false,
		EnableImportHosts: false,
	}
}

func (c *Config) LoadConfig() error {
	//c.RLock()
	//defer c.RUnlock()
	if _, err := os.Stat(c.getPath("config.yml")); os.IsNotExist(err) {
		err := ensureDirExists(c.getPath(""))
		if err != nil {
			log.Logger.Info("creating config directory", zap.Error(err))
		}
		c.DefaultConfig()
		return c.SaveFileConfig()
	}

	cData, err := os.ReadFile(c.getPath("config.yml"))
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(cData, c); err != nil {
		return err
	}

	return nil
}
