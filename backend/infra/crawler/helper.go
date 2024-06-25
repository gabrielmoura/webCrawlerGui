package crawler

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/db"
	"WebCrawlerGui/backend/infra/log"
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
func isDenyPostfix(url string, denySuffixes []string) bool {
	for _, denySuffix := range denySuffixes {
		if strings.HasSuffix(strings.ToLower(url), denySuffix) {
			return true
		}
	}
	return false
}

// isAllowedSchema checks if the link has an acceptable schema
func isAllowedSchema(link string, acceptableSchema []string) bool {
	nLink, err := url.Parse(link)
	if err != nil {
		log.Logger.Debug("error parsing link in checking schema", zap.String("Link", link), zap.Error(err))
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
func isAllowedMIME(contentType string, allowedMIMEs []string) bool {
	for _, allowedMIME := range allowedMIMEs {
		if strings.Contains(contentType, allowedMIME) {
			return true
		}
	}
	return false
}

// checkTLD checks if the link has an acceptable TLD
func checkTLD(link string) bool {
	if len(config.Conf.General.Tlds) > 0 {
		linkUrl, err := url.Parse(link)
		if err != nil {
			return false
		}
		for _, tld := range config.Conf.General.Tlds {
			if strings.HasSuffix(linkUrl.Hostname(), tld) {
				return true
			}
		}
		return false
	}
	return true
}

func handleAddToQueue(links []string, depth int) {
	for _, link := range links {
		if checkTLD(link) && isAllowedSchema(link, config.AcceptableSchema) {
			err := db.DB.AddToQueue(link, depth)
			if err != nil {
				log.Logger.Error("error adding link to queue", zap.String("Link", link), zap.Error(err))
				return
			}
		}

	}
}

// isLocalLink verifica se o link é local,
// caso definido para ignorar não adiciona a fila
func isLocalLink(link *url.URL) bool {
	if !config.Conf.General.IgnoreLocal {
		return false
	}
	return link.Host == "localhost" || link.Host == "127.0.0.1"
}

func prepareLink(link string) (*url.URL, error) {
	linkUrl, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if isLocalLink(linkUrl) {
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

	if isDenyPostfix(linkUrl.Path, config.DenySuffixes) {
		return nil, ErrDenySuffix
	}

	return linkUrl, nil
}
func prepareParentLink(parentLink, link string) (*url.URL, error) {

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
	log.Logger.Debug("New URL", zap.String("URL", nURL.String()))

	return nURL, nil
}
func isStatusErr(status int, url *url.URL) bool {
	if isSuccess(status) && isI2P(url) && url.Query().Has("i2paddresshelper") {
		handleListHelperI2P(url.String())
		return true
	}

	switch status {
	case http.StatusConflict:
		if isI2P(url) {
			handleListHelperI2P(url.String())
		}
		return true
	case http.StatusTooManyRequests:
		// TODO: Adicionar URL na lista de espera, pois o proxy pode estar bloqueando
		return true
	default:
		return !isSuccess(status)
	}
}

func isSuccess(status int) bool {
	return status >= 200 && status < 300
}

func isI2P(url *url.URL) bool {
	return strings.HasSuffix(url.Host, ".i2p")
}

func handleListHelperI2P(url string) {
	// TODO: Lógica para adicionar a lista de helpers encontrados
	log.Logger.Info("Helper I2P encontrado", zap.String("URL", url))
}
