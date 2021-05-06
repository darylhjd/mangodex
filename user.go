package mangodex

import (
	"context"
	"net/http"
	"net/url"
)

const (
	GetUserFollowedMangaFeed = "user/follows/manga/feed"
)

// GetUserFollowedMangaFeed : Return Chapter feed.
// https://api.mangadex.org/docs.html#operation/get-user-follows-manga-feed
func (dc *DexClient) GetUserFollowedMangaFeed(params url.Values) (*ChapterList, error) {
	return dc.GetUserFollowedMangaFeedContext(context.Background(), params)
}

// GetUserFollowedMangaFeedContext : GetUserFollowedMangaFeed with custom context.
func (dc *DexClient) GetUserFollowedMangaFeedContext(ctx context.Context, params url.Values) (*ChapterList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = GetUserFollowedMangaFeed

	// Set query parameters
	u.RawQuery = params.Encode()

	var l ChapterList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}
