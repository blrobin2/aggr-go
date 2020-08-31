package aggr

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/blrobin2/aggr-go/album"
	"github.com/blrobin2/aggr-go/feed"
)

func getAlbums(lowestScore int, currentMonth time.Month) []album.Album {
	a := feed.Pitchfork()
	album.DefaultSort(a)
	return album.DefaultFilter(lowestScore, currentMonth, a)
}

func main() {
	currentYear, currentMonth, _ := time.Now().Date()
	albums := getAlbums(80, currentMonth)
	albumsJSON, err := json.Marshal(albums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(albumsJSON)
}
