package data

import (
	"time"
)

type Page struct {
	Url         string           `json:"url" bson:"url" db:"url"`
	Links       []string         `json:"links" bson:"links" db:"links"`
	Title       string           `json:"title" bson:"title" db:"title"`
	Description string           `json:"description" bson:"description" db:"description"`
	Meta        *MetaData        `json:"meta,omitempty" bson:"meta" db:"meta"`
	Visited     bool             `json:"visited" bson:"visited" db:"visited"`
	Timestamp   time.Time        `json:"timestamp" bson:"timestamp" db:"timestamp"`
	Words       map[string]int32 `json:"words" bson:"words" db:"words"`
}
type MetaData struct {
	OG       map[string]string `json:"og,omitempty" bson:"og"`
	Keywords []string          `json:"keywords,omitempty" bson:"keywords"`
	Manifest string            `json:"manifest,omitempty" bson:"manifest"`
	Ld       string            `json:"ld,omitempty" bson:"ld"`
}
type PageVisited struct {
	Url     string `json:"url" bson:"url"`
	Visited bool   `json:"visited" bson:"visited"`
}

type PageSearch struct {
	Url   string `json:"url" bson:"url"`
	Title string `json:"title" bson:"title"`
}
type PageSearchWithFrequency struct {
	Url       string `json:"url" bson:"url"`
	Title     string `json:"title" bson:"title"`
	Frequency int32  `json:"frequency" bson:"frequency"`
}

type QueueType struct {
	Url   string `json:"url"`
	Depth int    `json:"depth"`
}

// PageIndex é a estrutura que contém as chaves para as páginas
type PageIndex struct {
	Keys []string `json:"url,omitempty"` // Lista ordenada de chaves para as páginas
}

type FailedType struct {
	Url    string `json:"url"`
	Reason string `json:"reason,omitempty"`
}

type Statistic struct {
	TotalPages        int            `json:"total_pages,omitempty"`
	TotalPagesPerHost map[string]int `json:"total_pages_per_host,omitempty"`
}
