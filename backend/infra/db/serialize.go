package db

import (
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/log"
	"WebCrawlerGui/backend/infra/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

// pageMarshal converts a data.Page instance into a protobuf-encoded byte slice.
func pageMarshal(page *data.Page) ([]byte, error) {
	pbPage := &pb.Page{
		Url:         page.Url,
		Title:       page.Title,
		Description: page.Description,
		Words:       page.Words,
		Links:       page.Links,
		Timestamp:   page.Timestamp.Format(time.RFC3339),
	}
	if page.Meta != nil {
		pbPage.Meta = &pb.PageMeta{
			Keywords: page.Meta.Keywords,
			Manifest: page.Meta.Manifest,
			Ld:       page.Meta.Ld,
			OG:       page.Meta.OG,
		}
	}

	return proto.Marshal(pbPage)
}

// pageUnmarshal converts a protobuf-encoded byte slice into a data.Page instance.
func pageUnmarshal(bytes []byte, page *data.Page) error {
	var pbPage pb.Page
	if err := proto.Unmarshal(bytes, &pbPage); err != nil {
		return err
	}
	page.Url = pbPage.Url
	page.Title = pbPage.Title
	page.Description = pbPage.Description
	page.Words = pbPage.Words
	page.Links = pbPage.Links
	timeStamp, err := time.Parse(time.RFC3339, pbPage.Timestamp)
	if err != nil {
		log.Logger.Error("Error parsing timestamp", zap.Error(err))
	}
	page.Timestamp = timeStamp

	if pbPage.Meta != nil {
		page.Meta = &data.MetaData{
			Keywords: pbPage.Meta.Keywords,
			Manifest: pbPage.Meta.Manifest,
			Ld:       pbPage.Meta.Ld,
			OG:       pbPage.Meta.OG,
		}
	}
	return nil
}

func indexPageMarshal(index *data.PageIndex) ([]byte, error) {
	pbPageIndex := &pb.PageIndex{
		Keys: index.Keys,
	}
	//return json.Marshal(index)
	return proto.Marshal(pbPageIndex)
}
func indexPageUnmarshal(bytes []byte, index *data.PageIndex) error {
	//return json.Unmarshal(bytes, index)
	var pbPageIndex pb.PageIndex
	if err := proto.Unmarshal(bytes, &pbPageIndex); err != nil {
		return err
	}
	index.Keys = pbPageIndex.Keys
	return nil
}
