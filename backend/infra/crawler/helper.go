package crawler

import (
	"WebCrawlerGui/backend/config"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"
)

var (
	invalidSchemaErr = errors.New("invalid schema")
	ErrLocalLink     = errors.New("local link")
	ErrEmptyPath     = errors.New("empty path")
	ErrDenySuffix    = errors.New("deny suffix")
)

// isDenyPostfix checks if the link has a deny postfix
func (c CrawlerService) isDenyPostfix(url string, denySuffixes []string) bool {
	for _, denySuffix := range denySuffixes {
		if strings.HasSuffix(strings.ToLower(url), denySuffix) {
			return true
		}
	}
	return false
}

// isAllowedSchema checks if the link has an acceptable schema
func (c CrawlerService) isAllowedSchema(link string, acceptableSchema []string) bool {
	nLink, err := url.Parse(link)
	if err != nil {
		c.logger.Debug("error parsing link in checking schema", zap.String("Link", link), zap.Error(err))
		return false
	}
	for _, schema := range acceptableSchema {
		if nLink.Scheme == schema {
			return true
		}
	}
	return false
}

// isAllowedMIME checks if the link has an acceptable MIME type
func (c CrawlerService) isAllowedMIME(contentType string, allowedMIMEs []string) bool {
	for _, allowedMIME := range allowedMIMEs {
		if strings.Contains(contentType, allowedMIME) {
			return true
		}
	}
	return false
}

// checkTLD checks if the link has an acceptable TLD
func (c CrawlerService) checkTLD(link string) bool {
	if len(config.Conf.Filter.Tlds) > 0 {
		linkUrl, err := url.Parse(link)
		if err != nil {
			return false
		}
		for _, tld := range config.Conf.Filter.Tlds {
			if strings.HasSuffix(linkUrl.Hostname(), tld) {
				return true
			}
		}
		return false
	}
	return true
}

func (c CrawlerService) handleAddToQueue(links []string, depth int) {
	for _, link := range links {
		if c.checkTLD(link) && c.isAllowedSchema(link, config.AcceptableSchema) {
			err := c.db.AddToQueue(link, depth)
			if err != nil {
				c.logger.Error("error adding link to queue", zap.String("Link", link), zap.Error(err))
				return
			}
		}

	}
}

// isLocalLink verifica se o link é local,
// caso definido para ignorar não adiciona a fila
func (c CrawlerService) isLocalLink(link *url.URL) bool {
	if !config.Conf.Filter.IgnoreLocal {
		return false
	}
	return link.Host == "localhost" || link.Host == "127.0.0.1"
}

func (c CrawlerService) prepareLink(link string) (*url.URL, error) {
	linkUrl, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if c.isLocalLink(linkUrl) {
		return nil, ErrLocalLink
	}

	if linkUrl.Scheme == "" {
		return nil, invalidSchemaErr
	}
	q, _ := url.ParseQuery(linkUrl.RawQuery)
	q.Del("utm_source")
	q.Del("utm_medium")
	q.Del("utm_campaign")
	q.Del("utm_term")
	q.Del("utm_content")
	q.Del("#")
	linkUrl.RawQuery = q.Encode()

	if c.isDenyPostfix(linkUrl.Path, config.DenySuffixes) {
		return nil, ErrDenySuffix
	}

	return linkUrl, nil
}
func (c CrawlerService) prepareParentLink(parentLink, link string) (*url.URL, error) {

	// Remove o primeiro caractere se for uma barra ou ponto
	if strings.HasPrefix(link, "/") || strings.HasPrefix(link, ".") {
		link = link[1:]
	}

	nURL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if nURL.Path == "" {
		return nil, ErrEmptyPath
	}

	pURL, err := url.Parse(parentLink)
	if err != nil {
		return nil, err
	}

	nURL.Host = pURL.Host
	nURL.Scheme = pURL.Scheme
	c.logger.Debug("New URL", zap.String("URL", nURL.String()))

	return nURL, nil
}
func (c CrawlerService) isStatusErr(status int, url *url.URL) bool {
	if status == http.StatusOK {
		return false
	}

	i2p := strings.HasSuffix(url.Host, ".i2p")
	switch status {
	case http.StatusConflict:
		if i2p {
			handleListHelperI2P(url.String())
		}
		return true
	default:
		return status < 200 || status >= 300
	}
}

func handleListHelperI2P(url string) {
	// TODO: Lógica para adicionar a lista de helpers encontrados
}
