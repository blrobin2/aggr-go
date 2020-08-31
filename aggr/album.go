package aggr

import (
	"sort"
	"strings"
	"time"
)

// Album represents an album that has been rated by some source
type Album struct {
	Artist string
	Title  string
	Date   time.Time
	Score  int
}

type lessFunc func(a1, a2 *Album) bool

// MultiSorter is a means of sorting by multiple fields
type MultiSorter struct {
	albums []Album
	less   []lessFunc
}

// Len returns the length of the albums stored in the sorter
func (ms *MultiSorter) Len() int {
	return len(ms.albums)
}

// OrderedBy returns a Sorted that sorts using the less functions, in order
func OrderedBy(less ...lessFunc) *MultiSorter {
	return &MultiSorter{
		less: less,
	}
}

// Swap swaps the place of two values, part of sort.Interface
func (ms *MultiSorter) Swap(i, j int) {
	ms.albums[i], ms.albums[j] = ms.albums[j], ms.albums[i]
}

func (ms *MultiSorter) Less(i, j int) bool {
	p, q := &ms.albums[i], &ms.albums[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q
			return true
		case less(q, p):
			// p > q
			return false
		}
		// p ==q; next comparator
	}

	// All comparisons are equal, so just return the last
	// comparison result
	return ms.less[k](p, q)
}

// Sort sorts the Album slice according to the less functions passed to OrderedBy
func (ms *MultiSorter) Sort(albums []Album) {
	ms.albums = albums
	sort.Sort(ms)
}

func byArtist(a1, a2 *Album) bool {
	return strings.ToLower(a1.Artist) < strings.ToLower(a2.Artist)
}

func byTitle(a1, a2 *Album) bool {
	return strings.ToLower(a1.Title) < strings.ToLower(a2.Title)
}

// DefaultSort is the default means of sorting a slice of Albums
func DefaultSort(albums []Album) {
	OrderedBy(byArtist, byTitle).Sort(albums)
}

// Filter removes albums from the slice based on condition function
func Filter(albums []Album, cond func(Album) bool) []Album {
	result := []Album{}
	for i := range albums {
		if cond(albums[i]) {
			result = append(result, albums[i])
		}
	}
	return result
}

// DefaultFilter is the default means of removing unwanted albums
func DefaultFilter(lowestScore int, currentMonth time.Month, albums []Album) []Album {
	return Filter(albums, func(album Album) bool {
		return scoreIsHighEnough(lowestScore, album) && cameOutThisMonth(currentMonth, album)
	})
}

func scoreIsHighEnough(lowestScore int, album Album) bool {
	return album.Score >= lowestScore
}

func cameOutThisMonth(currentMonth time.Month, album Album) bool {
	return album.Date.Month() == currentMonth
}
