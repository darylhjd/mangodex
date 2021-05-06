package mangodex

import "os"

var (
	testClient = NewDexClient()
	user, pwd  = os.Getenv("USER"), os.Getenv("PASSWORD")
)
