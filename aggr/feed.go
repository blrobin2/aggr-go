package aggr

import (
	"fmt"

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

	fmt.Println(feed)
	return []Album{}
}
