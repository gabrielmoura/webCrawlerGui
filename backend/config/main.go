package config

import (
	"strings"
)

func splitComma(txt string) []string {
	if txt == "" {
		return []string{}
	}
	return strings.Split(txt, ",")
}

type Config struct {
	MaxConcurrency int          `mapstructure:"MAX_CONCURRENCY"`
	MaxDepth       int          `mapstructure:"MAX_DEPTH"`
	PostgresURI    string       `mapstructure:"POSTGRES_URI"`
	AppName        string       `mapstructure:"APP_NAME"`
	TimeFormat     string       `mapstructure:"TIME_FORMAT"`
	TimeZone       string       `mapstructure:"TIME_ZONE"`
	InicialURL     string       `mapstructure:"URL"`
	Cache          *CacheConfig `mapstructure:"CACHE"`
	Proxy          *Proxy       `mapstructure:"PROXY"`
	Filter         *Filter      `mapstructure:"FILTER"`
	UserAgent      string       `mapstructure:"USER_AGENT"`
}
type CacheConfig struct {
	DBDir string `mapstructure:"DB_DIR"`
	Mode  string `mapstructure:"MODE"` // "mem" or "disc'
}
type Proxy struct {
	Enabled  bool   `mapstructure:"ENABLED"`
	ProxyURL string `mapstructure:"PROXY_URL"`
}
type Filter struct {
	Tlds        []string `mapstructure:"TLDS"`
	IgnoreLocal bool     `mapstructure:"IGNORE_LOCAL"`
}

var Conf *Config

func loadDefault() error {
	cfg := &Config{

		AppName:        "WebCrawler",
		TimeFormat:     "02-Jan-2006",
		TimeZone:       "America/Sao_Paulo",
		MaxConcurrency: 10,
		MaxDepth:       2,
		PostgresURI:    "postgres://postgres:Strong@P4ssword@localhost/crawler",
		UserAgent:      "Go-http-client/1.1",
		Cache: &CacheConfig{
			DBDir: "/tmp/WebCrawler",
			Mode:  "mem",
		},
		Proxy: &Proxy{
			Enabled:  false,
			ProxyURL: "",
		},
		Filter: &Filter{
			Tlds:        splitComma(""),
			IgnoreLocal: false,
		},
	}
	// Atualiza a variável global Conf
	Conf = cfg
	return nil
}
