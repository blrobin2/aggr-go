package main

import (
	"fmt"
	"time"

	"github.com/blrobin2/aggr-go/aggr"
)

func getAlbums() []aggr.Album {
	a := []aggr.Album{}
	done := make(chan []aggr.Album)

	for _, feed := range aggr.Feeds() {
		go func(res func() []aggr.Album) {
			done <- res()
		}(feed)
	}

	for range aggr.Feeds() {
		res := <-done
		fmt.Println(a)
		a = append(a, res...)
	}

	return a
}

func organizeAlbums(lowestScore float32, currentMonth time.Month, a []aggr.Album) []aggr.Album {
	a = aggr.UniqueAlbums(a)
	aggr.DefaultSort(a)
	return aggr.DefaultFilter(lowestScore, currentMonth, a)
}

func main() {
	_, currentMonth, _ := time.Now().Date()
	albums := organizeAlbums(float32(7.5), currentMonth, getAlbums())

	for _, album := range albums {
		fmt.Printf("%s by %s was released on %d-%02d-%02d and has a score of %.1f\n",
			album.Title, album.Artist, album.Date.Year(), album.Date.Month(), album.Date.Day(), album.Score,
		)
	}
}
