package mangodex

import (
	"net/url"
	"testing"
)

func TestCover(t *testing.T) {
	c := NewDexClient()
	list, err := c.Manga.GetMangaList(url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list.Data[0].ID)

	resp, err := c.Cover.GetMangaCoverList([]string{list.Data[0].ID}, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
