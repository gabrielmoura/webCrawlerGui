package crawler

import (
	"WebCrawlerGui/backend/config"
	"net/http"
	"net/url"
	"time"
)

func proxyClient() *http.Client {
	urlProxy, _ := url.Parse(config.Conf.General.ProxyURL)
	transport := &http.Transport{
		Proxy: http.ProxyURL(urlProxy),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // Definir um timeout de 5 segundos
	}

	return client
}
func httpClient() *http.Client {
	if config.Conf.General.ProxyEnabled {
		return proxyClient()
	} else {
		return &http.Client{
			Timeout: 5 * time.Second, // Definir um timeout de 5 segundos
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
			},
		}
	}
}

func httpRequest(pageUrl string) (resp *http.Response, err error) {
	client := httpClient()
	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Set("User-Agent", config.Conf.General.UserAgent)
	resp, err = client.Do(req)
	return
}
