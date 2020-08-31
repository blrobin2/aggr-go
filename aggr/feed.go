package aggr

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

// Feed is a function with side-effects that returns a slice of albums
type Feed func() []Album

// Feeds are the feeds we want to consume for Albums
func Feeds() []Feed {
	return []Feed{StereogumFeed, PitchforkFeed}
}

// StereogumFeed parses the Heavy Rotation RSS feed provided by stereogum
func StereogumFeed() []Album {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.stereogum.com/heavy-rotation/feed/")
	if err != nil {
		fmt.Println(err)
		return []Album{}
	}

	albums := []Album{}
	albumChan := make(chan Album)
	for _, item := range feed.Items {
		go func(album *gofeed.Item) {
			artistTitle := strings.Split(album.Title, " – ")

			albumChan <- Album{
				Title:  artistTitle[1],
				Artist: artistTitle[0],
				Date:   *album.PublishedParsed,
				Score:  8.0,
			}
		}(item)
	}

	for range feed.Items {
		album := <-albumChan
		albums = append(albums, album)
	}

	return albums
}

// PitchforkFeed parses the RSS feed provided by pitchfork.com
func PitchforkFeed() []Album {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://pitchfork.com/rss/reviews/albums/")
	if err != nil {
		fmt.Println(err)
		return []Album{}
	}

	albums := []Album{}
	albumChan := make(chan Album)
	for _, item := range feed.Items {
		go func(album *gofeed.Item) {
			artistTitle := strings.Split(album.Title, ": ")

			albumChan <- Album{
				Title:  artistTitle[1],
				Artist: artistTitle[0],
				Date:   *album.PublishedParsed,
				Score:  getPitchforkScore(album.Link),
			}
		}(item)
	}

	for range feed.Items {
		album := <-albumChan
		albums = append(albums, album)
	}

	return albums
}

func getPitchforkScore(link string) float32 {
	score := float32(0.0)
	response, err := http.Get(link)
	if err != nil {
		return score
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return score
	}

	document.Find(".score").Each(func(i int, s *goquery.Selection) {
		scoreString := s.Text()
		score64, err := strconv.ParseFloat(scoreString, 32)
		if err != nil {
			panic(err)
		}

		score = float32(score64)
	})

	return score
}
