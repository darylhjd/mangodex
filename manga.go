package mangodex

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	MangaListPath         = "manga"
	CreateMangaPath       = MangaListPath
	ViewMangaPath         = "manga/%s"
	UpdateMangaPath       = ViewMangaPath
	DeleteMangaPath       = ViewMangaPath
	AddMangaInListPath    = "manga/%s/list/%s"
	RemoveMangaInListPath = AddMangaInListPath
	UnfollowMangaPath     = "manga/%s/follow"
	FollowMangaPath       = UnfollowMangaPath
	MangaFeedPath         = "manga/%s/feed"
	MangaReadMarkersPath  = "manga/%s/read"
)

type MangaList struct {
	Results []MangaResponse `json:"results"`
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Total   int             `json:"total"`
}

type MangaResponse struct {
	Result        string       `json:"result"`
	Data          Manga        `json:"data"`
	Relationships Relationship `json:"relationships"`
}

func (mr *MangaResponse) GetResult() string {
	return mr.Result
}

type Manga struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes MangaAttributes `json:"attributes"`
}

type Relationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type MangaAttributes struct {
	Title                  LocalisedString    `json:"title"`
	AltTitles              []*LocalisedString `json:"altTitles"`
	Description            LocalisedString    `json:"description"`
	IsLocked               bool               `json:"isLocked"`
	Links                  []*string          `json:"links"`
	OriginalLanguage       string             `json:"originalLanguage"`
	LastVolume             *string            `json:"lastVolume"`
	LastChapter            *string            `json:"lastChapter"`
	PublicationDemographic *string            `json:"publicationDemographic"`
	Status                 *string            `json:"status"`
	Year                   int                `json:"year"`
	ContentRating          *string            `json:"contentRating"`
	Tags                   []*LocalisedString `json:"tags"`
	Version                int                `json:"version"`
	CreatedAt              string             `json:"createdAt"`
	UpdatedAt              string             `json:"updatedAt"`
}

type LocalisedString struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type ReadMarkersResponse struct {
	Result string   `json:"result"`
	Data   []string `json:"data"`
}

func (rmr *ReadMarkersResponse) GetResult() string {
	return rmr.Result
}

// MangaList : Get a list of manga based on query parameters.
// https://api.mangadex.org/docs.html#operation/get-search-manga
func (dc *DexClient) MangaList(params url.Values) (*MangaList, error) {
	return dc.MangaListContext(context.Background(), params)
}

// MangaListContext : MangaList with custom context.
func (dc *DexClient) MangaListContext(ctx context.Context, params url.Values) (*MangaList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = MangaListPath

	// Set query parameters
	u.RawQuery = params.Encode()

	var l MangaList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}

// CreateManga : Create a new manga.
// https://api.mangadex.org/docs.html#operation/post-manga
func (dc *DexClient) CreateManga(newManga io.Reader) (*MangaResponse, error) {
	return dc.CreateMangaContext(context.Background(), newManga)
}

// CreateMangaContext : CreateManga with custom context.
func (dc *DexClient) CreateMangaContext(ctx context.Context, newManga io.Reader) (*MangaResponse, error) {
	var mr MangaResponse
	err := dc.responseOp(ctx, http.MethodPost, CreateMangaPath, newManga, &mr)
	return &mr, err
}

// ViewManga : View a manga by ID.
// https://api.mangadex.org/docs.html#operation/get-manga-id
func (dc *DexClient) ViewManga(id string) (*MangaResponse, error) {
	return dc.ViewMangaContext(context.Background(), id)
}

// ViewMangaContext : ViewManga with custom context.
func (dc *DexClient) ViewMangaContext(ctx context.Context, id string) (*MangaResponse, error) {
	var mr MangaResponse
	err := dc.responseOp(ctx, http.MethodGet, fmt.Sprintf(ViewMangaPath, id), nil, &mr)
	return &mr, err
}

// UpdateManga : Update a Manga.
// https://api.mangadex.org/docs.html#operation/put-manga-id
func (dc *DexClient) UpdateManga(id string, upManga io.Reader) (*MangaResponse, error) {
	return dc.UpdateMangaContext(context.Background(), id, upManga)
}

// UpdateMangaContext : UpdateManga with custom context.
func (dc *DexClient) UpdateMangaContext(ctx context.Context, id string, upManga io.Reader) (*MangaResponse, error) {
	var mr MangaResponse
	err := dc.responseOp(ctx, http.MethodPut, fmt.Sprintf(UpdateMangaPath, id), upManga, &mr)
	return &mr, err
}

// DeleteManga : Delete a Manga through ID.
// https://api.mangadex.org/docs.html#operation/delete-manga-id
func (dc *DexClient) DeleteManga(id string) error {
	return dc.DeleteMangaContext(context.Background(), id)
}

// DeleteMangaContext : DeleteManga with custom context.
func (dc *DexClient) DeleteMangaContext(ctx context.Context, id string) error {
	return dc.responseOp(ctx, http.MethodDelete, fmt.Sprintf(DeleteMangaPath, id), nil, nil)
}

// AddMangaInList : Add a Manga to a custom list.
// https://api.mangadex.org/docs.html#operation/post-manga-id-list-listId
func (dc *DexClient) AddMangaInList(mangaId, listId string) error {
	return dc.AddMangaInListContext(context.Background(), mangaId, listId)
}

// AddMangaInListContext : AddMangaInList with custom context.
func (dc *DexClient) AddMangaInListContext(ctx context.Context, mangaId, listId string) error {
	return dc.responseOp(ctx, http.MethodPost, fmt.Sprintf(AddMangaInListPath, mangaId, listId), nil, nil)
}

// RemoveMangaInList : Remove a Manga from a custom list.
// https://api.mangadex.org/docs.html#operation/delete-manga-id-list-listId
func (dc *DexClient) RemoveMangaInList(mangaId, listId string) error {
	return dc.RemoveMangaInListContext(context.Background(), mangaId, listId)
}

// RemoveMangaInListContext : RemoveMangaInList with custom context.
func (dc *DexClient) RemoveMangaInListContext(ctx context.Context, mangaId, listId string) error {
	return dc.responseOp(ctx, http.MethodDelete, fmt.Sprintf(RemoveMangaInListPath, mangaId, listId), nil, nil)
}

// UnfollowManga : Unfollow a Manga by ID.
// https://api.mangadex.org/docs.html#operation/delete-manga-id-follow
func (dc *DexClient) UnfollowManga(id string) error {
	return dc.UnfollowMangaContext(context.Background(), id)
}

// UnfollowMangaContext : UnfollowManga with custom context.
func (dc *DexClient) UnfollowMangaContext(ctx context.Context, id string) error {
	return dc.responseOp(ctx, http.MethodDelete, fmt.Sprintf(UnfollowMangaPath, id), nil, nil)
}

// FollowManga : Follow a Manga by ID.
// https://api.mangadex.org/docs.html#operation/post-manga-id-follow
func (dc *DexClient) FollowManga(id string) error {
	return dc.FollowMangaContext(context.Background(), id)
}

// FollowMangaContext : FollowManga with custom context.
func (dc *DexClient) FollowMangaContext(ctx context.Context, id string) error {
	return dc.responseOp(ctx, http.MethodPost, fmt.Sprintf(FollowMangaPath, id), nil, nil)
}

// MangaFeed : Get Manga feed by ID.
// https://api.mangadex.org/docs.html#operation/get-manga-id-feed
func (dc *DexClient) MangaFeed(id string, params url.Values) (*ChapterList, error) {
	return dc.MangaFeedContext(context.Background(), id, params)
}

// MangaFeedContext : MangaFeed with custom context.
func (dc *DexClient) MangaFeedContext(ctx context.Context, id string, params url.Values) (*ChapterList, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = fmt.Sprintf(MangaFeedPath, id)

	// Set request parameters
	u.RawQuery = params.Encode()

	var l ChapterList
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &l)
	return &l, err
}

// MangaReadMarkers : Get list of Chapter IDs that are marked as read for a specified manga ID.
// https://api.mangadex.org/docs.html#operation/get-manga-chapter-readmarkers
func (dc *DexClient) MangaReadMarkers(id string) (*ReadMarkersResponse, error) {
	return dc.MangaReadMarkersContext(context.Background(), id)
}

// MangaReadMarkersContext : MangaReadMarkers with custom context.
func (dc *DexClient) MangaReadMarkersContext(ctx context.Context, id string) (*ReadMarkersResponse, error) {
	var rmr ReadMarkersResponse
	err := dc.responseOp(ctx, http.MethodGet, fmt.Sprintf(MangaReadMarkersPath, id), nil, &rmr)
	return &rmr, err
}
