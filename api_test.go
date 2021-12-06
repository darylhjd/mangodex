package mangodex

import (
	"net/url"
	"strconv"
	"testing"
)

var client = NewDexClient()

func TestLogin(t *testing.T) {
	err := client.Auth.Login("-", "-")
	if err != nil {
		t.Error("Login failed.")
	}
}

func TestGetMangaList(t *testing.T) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(100))
	params.Set("offset", strconv.Itoa(0))
	// Include Author relationship
	params.Set("includes[]", AuthorRel)
	// If it is a search, then we add the search term.
	list, err := client.Manga.GetMangaList(params)
	if err != nil {
		t.Errorf("Getting manga failed: %s\n", err.Error())
	}
	t.Log(list)
}
