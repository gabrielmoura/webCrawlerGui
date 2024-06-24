package helper

import (
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/types"
	"net/url"
)

// NormalizeURL Função para normalizar URLs removendo a barra final
func NormalizeURL(urlStr string) string {
	if urlStr[len(urlStr)-1] == '/' {
		return urlStr[:len(urlStr)-1]
	}
	return urlStr
}

// PageToTreeNode Função recursiva para converter Page para TreeNode utilizando um mapa para acesso rápido e agrupando por host
func PageToTreeNode(page *data.Page, pageMap map[string]*data.Page, host string, visited map[string]bool) types.TreeNode {
	if visited[page.Url] {
		return types.TreeNode{}
	}

	visited[page.Url] = true

	treeNode := types.TreeNode{
		Title:       page.Title,
		Description: page.Description,
		URL:         page.Url,
	}

	for _, link := range page.Links {
		normalizedLink := NormalizeURL(link)
		u, err := url.Parse(normalizedLink)
		if err != nil {
			continue // Ignorar URLs inválidas
		}

		if NormalizeURL(u.Scheme+"://"+u.Host) == host {
			if childPage, found := pageMap[link]; found {
				if normalizedLink != host { // Evitar adicionar a URL do nó pai aos filhos
					childTreeNode := PageToTreeNode(childPage, pageMap, host, visited)

					if childTreeNode.Title != "" {
						treeNode.Children = append(treeNode.Children, childTreeNode)
					}
				}
			}
		}
	}

	return treeNode
}
