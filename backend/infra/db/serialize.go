package db

import (
	"WebCrawlerGui/backend/infra/data"
	"WebCrawlerGui/backend/infra/pb"

	"google.golang.org/protobuf/proto"
)

// pageMarshal converts a data.Page instance into a protobuf-encoded byte slice.
func pageMarshal(page *data.Page) ([]byte, error) {
	pbPage := &pb.Page{
		Url:         page.Url,
		Title:       page.Title,
		Description: page.Description,
		Words:       page.Words,
		Links:       page.Links,
		Meta: &pb.PageMeta{
			Keywords: page.Meta.Keywords,
			Manifest: page.Meta.Manifest,
			Ld:       page.Meta.Ld,
			OG:       page.Meta.OG,
		},
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
	page.Meta = &data.MetaData{
		Keywords: pbPage.Meta.Keywords,
		Manifest: pbPage.Meta.Manifest,
		Ld:       pbPage.Meta.Ld,
		OG:       pbPage.Meta.OG,
	}
	return nil
}
