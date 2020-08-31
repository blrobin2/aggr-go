package aggr

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

// Pitchfork parses the RSS feed provided by pitchfork.com
func Pitchfork() []album.Album {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://pitchfork.com/rss/reviews/albums/")
	if err != nil {
		fmt.Println(err)
		return []album.Album{}
	}

	fmt.Println(feed)
	return []album.Album{}
}
