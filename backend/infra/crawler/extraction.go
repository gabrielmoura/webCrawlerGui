package crawler

import (
	"WebCrawlerGui/backend/config"
	"WebCrawlerGui/backend/infra/data"
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var InvalidMeta = errors.New("invalid meta")

// countWordsInText Extrai e conta a frequência de palavras do conteúdo HTML, ignorando palavras irrelevantes comuns.
func (c CrawlerService) countWordsInText(data []byte) (map[string]int, error) {
	c.logger.Debug("Word Count")
	// Etapa 1: Ignorar determinadas tags HTML
	htmlRegex := regexp.MustCompile("(?s)<(script|style|noscript|link|meta)[^>]*?>.*?</(script|style|noscript|link|meta)>")
	parcialPlainText := htmlRegex.ReplaceAll(data, []byte(""))

	// Etapa 2: remover tags HTML
	tagsRegex := regexp.MustCompile("<([^>]*)>")
	plainText := tagsRegex.ReplaceAll(parcialPlainText, []byte(""))

	// Etapa 3: Normalizar texto
	normalizedText := bytes.ToLower(plainText)

	// Etapa 4: Remova caracteres especiais e divida em palavras
	wordRegex := regexp.MustCompile("[^\\pL\\pN\\pZ'-]+")
	noSpecialCh := wordRegex.ReplaceAll(normalizedText, []byte(" "))
	words := bytes.Split(noSpecialCh, []byte(" "))

	// Etapa 5: Conte a frequência das palavras (ignorando palavras comuns)
	wordCounts := make(map[string]int)
	for _, wordBytes := range words {
		word := string(bytes.TrimSpace(wordBytes))

		// Pule palavras curtas e palavras de parada comuns
		if len(word) < 2 || c.containsMap(config.CommonStopWords, word) {
			continue
		}
		wordCounts[word]++
		c.logger.Debug("Word: ", zap.Int(word, wordCounts[word]))
	}

	return wordCounts, nil
}

// ContainsMap Verifica se uma palavra está em uma lista de stop words comuns,
func (c CrawlerService) containsMap(wordMap map[string][]string, item string) bool {
	for key, slice := range wordMap {
		// Ignora a primeira string do mapa (chave vazia ou primeira chave lexicograficamente)
		if key == "" {
			continue
		}

		for _, a := range slice {
			if a == item {
				return true
			}
		}
	}
	return false
}

func (c CrawlerService) extractData(n *html.Node) (*data.Page, error) {
	var dataPage data.Page

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				c.extractTitle(n, &dataPage)
			case "meta":
				c.extractMeta(n, &dataPage)
			case "script":
				c.extractJSONLD(n, &dataPage)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(n)
	return &dataPage, nil
}

func (c CrawlerService) extractDescription(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "name" && a.Val == "description" {
			for _, a := range n.Attr {
				if a.Key == "content" {
					return a.Val
				}
			}
		}
	}
	return ""
}

func (c CrawlerService) extractOG(n *html.Node) (map[string]string, error) {
	ogData := make(map[string]string)
	if n.Data != "meta" {
		return nil, InvalidMeta
	}
	for _, a := range n.Attr {
		if a.Key == "property" && len(a.Val) > 3 && a.Val[:3] == "og:" {
			for _, b := range n.Attr {
				if b.Key == "content" {
					ogData[a.Val] = b.Val
				}
			}
		}
	}

	return ogData, nil
}

func (c CrawlerService) extractKeywords(n *html.Node) ([]string, error) {
	if n.Data != "meta" {
		return nil, InvalidMeta
	}
	for _, a := range n.Attr {
		if a.Key == "name" && a.Val == "keywords" {
			for _, a := range n.Attr {
				if a.Key == "content" {
					return strings.Split(a.Val, ","), nil
				}
			}
		}
	}
	return nil, nil
}

func (c CrawlerService) extractManifest(n *html.Node) string {
	if n.Data == "link" {
		var isManifest bool
		for _, a := range n.Attr {
			if a.Key == "rel" && a.Val == "manifest" {
				isManifest = true
			}
			if a.Key == "href" && isManifest {
				return a.Val
			}
		}
	}
	return ""
}

func (c CrawlerService) extractMeta(n *html.Node, dataPage *data.Page) {
	if n.Data == "meta" {
		if dataPage.Meta == nil {
			dataPage.Meta = metaNull
		}
		description := c.extractDescription(n)
		if description != "" {
			dataPage.Description = description
		}

		ogData, err := c.extractOG(n)
		if err == nil {
			for k, v := range ogData {
				dataPage.Meta.OG[k] = v
			}
		}

		keywords, err := c.extractKeywords(n)
		if keywords != nil && err == nil {
			dataPage.Meta.Keywords = keywords
		}
	}

	manifest := c.extractManifest(n)
	if manifest != "" {
		dataPage.Meta.Manifest = manifest
	}

	if dataPage.Meta.Manifest == "" && len(dataPage.Meta.OG) == 0 && len(dataPage.Meta.Keywords) == 0 {
		dataPage.Meta = nil
	}
}

func (c CrawlerService) extractTitle(n *html.Node, dataPage *data.Page) {
	if n.Data == "title" && n.FirstChild != nil {
		dataPage.Title = n.FirstChild.Data
	}
}

// extractLinks Extrai links de um documento HTML.
func (c CrawlerService) extractLinks(parentLink string, n *html.Node) ([]string, error) {
	var links []string

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urlE, err := c.prepareLink(a.Val)
					if err != nil {
						if errors.Is(invalidSchemaErr, err) {
							preparedLink, err := c.prepareParentLink(parentLink, a.Val)
							if err != nil {
								continue
							}
							urlE = preparedLink
						}
						c.logger.Debug(fmt.Sprintf("Error preparing link: %s", err))
						continue
					}
					links = append(links, urlE.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(n)
	return links, nil
}
func (c CrawlerService) extractJSONLD(n *html.Node, dataPage *data.Page) {
	for _, a := range n.Attr {
		if a.Key == "type" && a.Val == "application/ld+json" {
			if dataPage.Meta == nil {
				dataPage.Meta = metaNull
			}
			if n.FirstChild != nil {
				content := n.FirstChild.Data
				dataPage.Meta.Ld = content
			}
		}
	}
}

var metaNull = &data.MetaData{
	OG:       make(map[string]string),
	Keywords: []string{},
	Manifest: "",
	Ld:       "",
}
