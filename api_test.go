package mangodex

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"
)

var client = NewDexClient()

func TestLogin(t *testing.T) {
	err := client.Auth.Login(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		t.Error("Login failed.")
	}
	fmt.Printf("%v\n", client)
}

func TestGetLoggedUser(t *testing.T) {
	user, err := client.User.GetLoggedUser()
	if err != nil {
		t.Error("Getting user failed.")
	}
	t.Log(user)
}

func TestGetMangaList(t *testing.T) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(100))
	params.Set("offset", strconv.Itoa(0))
	// Include Author relationship
	params.Set("includes[]", AuthorRel)
	// If it is a search, then we add the search term.
	_, err := client.Manga.GetMangaList(params)
	if err != nil {
		t.Errorf("Getting manga failed: %s\n", err.Error())
	}
}
