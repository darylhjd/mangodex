package mangodex

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	GetMangaCoverListPath = "cover"
)

// CoverService : Provides Cover services provided by the API.
type CoverService service

// CoverArtList : A response for getting a list of CoverArts.
type CoverArtList struct {
	Result   string  `json:"result"`
	Response string  `json:"response"`
	Data     []Cover `json:"data"`
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	Total    int     `json:"total"`
}

func (cal *CoverArtList) GetResult() string {
	return cal.Result
}

type Cover struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes CoverAttributes `json:"attributes"`
}

type CoverAttributes struct {
	Volume      *string `json:"volume"`
	FileName    string  `json:"fileName"`
	Description *string `json:"description"`
	Version     int     `json:"version"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// GetMangaCoverList : Get manga cover by ID.
// https://api.mangadex.org/docs.html#operation/get-cover
func (s *CoverService) GetMangaCoverList(ids []string, isMangaCover bool) (*CoverArtList, error) {
	return s.GetMangaCoverContextList(context.Background(), ids, isMangaCover)
}

// GetMangaCoverContextList : GetMangaCover with custom context.
func (s *CoverService) GetMangaCoverContextList(ctx context.Context, ids []string, isMangaCover bool) (*CoverArtList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = GetMangaCoverListPath

	// Set request body.
	req := map[string][]string{}
	if isMangaCover {
		req["manga"] = ids
	} else {
		req["ids"] = ids
	}
	rBytes, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	var r CoverArtList
	err = s.client.RequestAndDecode(ctx, http.MethodGet, u.String(), bytes.NewBuffer(rBytes), &r)
	return &r, err
}
