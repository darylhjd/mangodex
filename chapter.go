package mangodex

import (
	"context"
	"fmt"
	"net/http"
)

const (
	DeleteChapterPath = "chapter/%s"
)

type ChapterList struct {
	Results []ChapterResponse `json:"results"`
	Limit   int               `json:"limit"`
	Offset  int               `json:"offset"`
	Total   int               `json:"total"`
}

type ChapterResponse struct {
	Result        string         `json:"result"`
	Data          Chapter        `json:"data"`
	Relationships []Relationship `json:"relationships"`
}

type Chapter struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes ChapterAttributes `json:"attributes"`
}

type ChapterAttributes struct {
	Title              string   `json:"title"`
	Volume             int      `json:"volume"`
	Chapter            string   `json:"chapter"`
	TranslatedLanguage string   `json:"translatedLanguage"`
	Hash               string   `json:"hash"`
	Data               []string `json:"data"`
	DataSaver          []string `json:"dataSaver"`
	Uploader           string   `json:"uploader"`
	Version            int      `json:"version"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedAt          string   `json:"updatedAt"`
	PublishAt          string   `json:"publishAt"`
}

// DeleteChapter : Remove a chapter by ID.
// https://api.mangadex.org/docs.html#operation/delete-chapter-id
func (dc *DexClient) DeleteChapter(id string) error {
	return dc.DeleteChapterContext(context.Background(), id)
}

// DeleteChapterContext : DeleteChapter with custom context.
func (dc *DexClient) DeleteChapterContext(ctx context.Context, id string) error {
	return dc.responseOp(ctx, http.MethodDelete, fmt.Sprintf(DeleteChapterPath, id), nil, nil)
}
