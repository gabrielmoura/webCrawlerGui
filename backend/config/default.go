package config

import "WebCrawlerGui/backend/config/stopwords"

var QueueDataPrefix = "queueIndex"
var VisitedIndexName = "visitedIndex"
var PageDataIndexName = "pageDataIndex"
var PageDataPrefix = "pageData"

var PrefixFailedData = "failed"

// AcceptableMimeTypes Mimes aceitos, checagem quando visitado
var AcceptableMimeTypes = []string{
	"text/html",
	"text/plain",
	"text/xml",
	"application/xml",
	"application/xhtml+xml",
	"application/rss+xml",
	"application/atom+xml",
	"application/rdf+xml",
	"application/json",
	"application/ld+json",
	"application/vnd.geo+json",
	"application/xml-dtd",
	"application/rss+xml",
	"application/atom+xml",
	"application/rdf+xml",
	"application/json",
	"application/ld+json",
	"application/vnd.geo+json",
}

// AcceptableSchema Schemas permitidos
var AcceptableSchema = []string{
	"http",
	"https",
	"",
}

// DenySuffixes Impede urls com estes sufixos de serem visitadas.
var DenySuffixes = []string{
	".css",
	".js",
	".png",
	".jpg",
	".jpeg",
	".gif",
	".svg",
	".ico",
	".mp4",
	".mp3",
	".avi",
	".flv",
	".mpeg",
	".webp",
	".webm",
	".woff",
	".woff2",
	".ttf",
	".eot",
	".otf",
	".pdf",
	".zip",
	".tar",
	".gz",
	".bz2",
	".xz",
	".7z",
	".rar",
	".apk",
	".exe",
	".dmg",
	".img",
	".pdf",
}

// CommonStopWords Palavras de parada comuns (personalize conforme necessário)
var CommonStopWords = map[string][]string{
	"en":    stopwords.En,
	"pt":    stopwords.Pt,
	"ru":    stopwords.Ru,
	"es":    stopwords.Es,
	"hindi": stopwords.Hin,
	"ch":    stopwords.Ch,
}
