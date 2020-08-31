package aggr

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
)

// PitchforkFeed parses the RSS feed provided by pitchfork.com
func PitchforkFeed() []Album {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://pitchfork.com/rss/reviews/albums/")
	if err != nil {
		fmt.Println(err)
		return []Album{}
	}

	albums := []Album{}

	for _, album := range feed.Items {
		artistTitle := strings.Split(album.Title, ": ")
		albums = append(albums, Album{
			Title:  artistTitle[1],
			Artist: artistTitle[0],
			Date:   *album.PublishedParsed,
			Score:  80,
		})
	}

	return albums
}
