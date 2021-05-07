package mangodex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	GetUserFollowedMangaChapterFeedPath = "user/follows/manga/feed"
	CustomListMangaChapterFeedPath      = "list/%s/feed"
)

// GetUserFollowedMangaChapterFeed : Return Chapter feed.
// https://api.mangadex.org/docs.html#operation/get-user-follows-manga-feed
func (dc *DexClient) GetUserFollowedMangaChapterFeed(params url.Values) (*ChapterList, error) {
	return dc.GetUserFollowedMangaChapterFeedContext(context.Background(), params)
}

// GetUserFollowedMangaChapterFeedContext : GetUserFollowedMangaChapterFeedPath with custom context.
func (dc *DexClient) GetUserFollowedMangaChapterFeedContext(ctx context.Context, params url.Values) (*ChapterList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = GetUserFollowedMangaChapterFeedPath

	// Set query parameters
	u.RawQuery = params.Encode()

	var l ChapterList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}

// CustomListMangaChapterFeed : Return Chapter feed from custom manga list.
// https://api.mangadex.org/docs.html#operation/get-list-id-feed
func (dc *DexClient) CustomListMangaChapterFeed(id string, params url.Values) (*ChapterList, error) {
	return dc.CustomListMangaChapterFeedContext(context.Background(), id, params)
}

// CustomListMangaChapterFeedContext : CustomListMangaFeed with custom context.
func (dc *DexClient) CustomListMangaChapterFeedContext(ctx context.Context, id string, params url.Values) (*ChapterList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = fmt.Sprintf(CustomListMangaChapterFeedPath, id)

	// Set query parameters
	u.RawQuery = params.Encode()

	var l ChapterList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}
