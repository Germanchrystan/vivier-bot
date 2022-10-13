package composer

import (
	"math/rand"
	"time"

	"github.com/wesovilabs/koazee"
)

type Coordinate [2]int

type Chord struct {
	Coors []Coordinate
	Notes []Note
}

func newChord(coors []Coordinate, xNote, yNote Note) Chord {
	var notes []Note
	for _, coor := range coors {
		freq := xNote.Freq*float32(coor[0]) + yNote.Freq*float32(coor[1])
		note := FindNearestNote(freq, NotesCatalog)
		notes = append(notes, note)
	}
	return Chord{
		Coors: coors,
		Notes: notes,
	}
}

// Create a slice of Vectors for chords. Using koazee lib to use "contains" method
func NewChordVectors(num int) []Coordinate {
	var coors []Coordinate
	stream := koazee.StreamOf(coors)
	for len(coors) != num {
		max := getRandomIntBetweenRanges(num, 2)
		coor := NewCoord(max)
		contains, _ := stream.Contains(coor)
		if !contains {
			coors = append(coors, coor)
			stream = koazee.StreamOf(coors)
		}
	}
	return coors
}

// Generates a random coord, using a num argument as the sum of both numbers
func NewCoord(num int) Coordinate {
	partition := getRandomIntBetweenRanges(num, 1)
	return Coordinate{partition, num - partition}
}

// Get a random number between to values
func getRandomIntBetweenRanges(max, min int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn((max)-min) + min
}
