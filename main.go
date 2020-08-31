package main

import (
	"fmt"
	"time"

	"github.com/blrobin2/aggr-go/aggr"
)

func getAlbums() []aggr.Album {
	albums := []aggr.Album{}
	albumChan := make(chan []aggr.Album)

	for _, feed := range aggr.Feeds() {
		go func(feedFunc func() []aggr.Album) {
			albumChan <- feedFunc()
		}(feed)
	}

	for range aggr.Feeds() {
		response := <-albumChan
		albums = append(albums, response...)
	}

	return albums
}

func organizeAlbums(lowestScore float32, currentMonth time.Month, albums []aggr.Album) []aggr.Album {
	albums = aggr.UniqueAlbums(albums)
	albums = aggr.DefaultFilter(lowestScore, currentMonth, albums)
	aggr.DefaultSort(albums)
	return albums
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
