package main

import (
	"fmt"
	"time"

	"github.com/blrobin2/aggr-go/aggr"
)

func getAlbums(lowestScore float32, currentMonth time.Month) []aggr.Album {
	a := aggr.PitchforkFeed()
	aggr.DefaultSort(a)
	return aggr.DefaultFilter(lowestScore, currentMonth, a)
}

func main() {
	_, currentMonth, _ := time.Now().Date()
	albums := getAlbums(float32(8.0), currentMonth)
	fmt.Println(albums)
	// albumsJSON, err := json.Marshal(albums)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(albumsJSON)
}
