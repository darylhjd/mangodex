package mangodex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	GetUserPath                         = "user/%s"
	GetLoggedUserPath                   = "user/me"
	GetUserFollowedMangaChapterFeedPath = "user/follows/manga/feed"
	GetUserFollowedMangaListPath        = "user/follows/manga"
	GetUserMangaReadingStatusPath       = "manga/status"
	GetUserFollowedScanGroupListPath    = "user/follows/group"
)

type UserList struct {
	Results []UserResponse `json:"results"`
	Limit   int            `json:"limit"`
	Offset  int            `json:"offset"`
	Total   int            `json:"total"`
}

type UserResponse struct {
	Result        string         `json:"result"`
	Data          User           `json:"data"`
	Relationships []Relationship `json:"relationships"`
}

func (r *UserResponse) GetResult() string {
	return r.Result
}

type User struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes UserAttributes `json:"attributes"`
}

type UserAttributes struct {
	Username string `json:"username"`
	Version  int    `json:"version"`
}

// GetUser : Return a UserResponse.
func (dc *DexClient) GetUser(id string) (*UserResponse, error) {
	return dc.GetUserContext(context.Background(), id)
}

// GetUserContext : GetUser with custom context.
func (dc *DexClient) GetUserContext(ctx context.Context, id string) (*UserResponse, error) {
	var r UserResponse
	err := dc.responseOp(ctx, http.MethodGet, fmt.Sprintf(GetUserPath, id), nil, &r)
	return &r, err
}

// GetLoggedUser : Return logged UserResponse.
func (dc *DexClient) GetLoggedUser() (*UserResponse, error) {
	return dc.GetLoggedUserContext(context.Background())
}

// GetLoggedUserContext : GetLoggedUser with custom context.
func (dc *DexClient) GetLoggedUserContext(ctx context.Context) (*UserResponse, error) {
	var r UserResponse
	err := dc.responseOp(ctx, http.MethodGet, GetLoggedUserPath, nil, &r)
	return &r, err
}

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

// GetUserFollowedMangaList : Return list of followed Manga.
// https://api.mangadex.org/docs.html#operation/get-user-follows-manga
func (dc *DexClient) GetUserFollowedMangaList(limit, offset int) (*MangaList, error) {
	return dc.GetUserFollowedMangaListContext(context.Background(), limit, offset)
}

// GetUserFollowedMangaListContext : GetUserFollowedMangaListPath with custom context.
func (dc *DexClient) GetUserFollowedMangaListContext(ctx context.Context, limit, offset int) (*MangaList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = GetUserFollowedMangaListPath

	// Set required query parameters
	q := u.Query()
	q.Add("limit", strconv.Itoa(limit))
	q.Add("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()

	var l MangaList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}

// GetUserMangaReadingStatus : Get reading status for all manga for logged user.
// https://api.mangadex.org/docs.html#operation/get-manga-status
func (dc *DexClient) GetUserMangaReadingStatus() (*MangaReadingStatusResponse, error) {
	return dc.GetUserMangaReadingStatusContext(context.Background())
}

// GetUserMangaReadingStatusContext : GetUserMangaReadingStatus with custom context.
func (dc *DexClient) GetUserMangaReadingStatusContext(ctx context.Context) (*MangaReadingStatusResponse, error) {
	var r MangaReadingStatusResponse
	err := dc.responseOp(ctx, http.MethodGet, GetUserMangaReadingStatusPath, nil, &r)
	return &r, err
}

// GetUserFollowedScanGroupList : Return list of followed ScanGroup.
// https://api.mangadex.org/docs.html#operation/get-user-follows-group
func (dc *DexClient) GetUserFollowedScanGroupList(limit, offset int) (*ScanGroupList, error) {
	return dc.GetUserFollowedScanGroupListContext(context.Background(), limit, offset)
}

// GetUserFollowedScanGroupListContext : GetUserFollowedScanGroupList with custom context.
func (dc *DexClient) GetUserFollowedScanGroupListContext(ctx context.Context, limit, offset int) (*ScanGroupList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = GetUserFollowedScanGroupListPath

	// Set query parameters
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()

	var l ScanGroupList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}
