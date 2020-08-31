package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/blrobin2/aggr-go/aggr"
)

func getAlbums(lowestScore int, currentMonth time.Month) []aggr.Album {
	a := aggr.PitchforkFeed()
	aggr.DefaultSort(a)
	return aggr.DefaultFilter(lowestScore, currentMonth, a)
}

func main() {
	_, currentMonth, _ := time.Now().Date()
	albums := getAlbums(80, currentMonth)
	albumsJSON, err := json.Marshal(albums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(albumsJSON)
}
