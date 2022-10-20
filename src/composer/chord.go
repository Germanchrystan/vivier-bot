package composer

import (
	"github.com/wesovilabs/koazee"
)

type Vector [2]int

type Chord struct {
	BaseNotes [2]Note
	Color     []Vector
	Notes     []Note
}

// Create a slice of Vectors for chords. Using koazee lib to use "contains" method
func NewColor(vectorPoints, vectorQuantity int) []Vector {
	var color []Vector
	stream := koazee.StreamOf(color)

	for len(color) != vectorQuantity {
		max := getRandomIntBetweenRanges(2, vectorPoints)
		coor := NewVector(max)
		contains, _ := stream.Contains(coor)
		if !contains {
			color = append(color, coor)
			stream = koazee.StreamOf(color)
		}
	}
	return color
}

// Generates a random coord, using a num argument as the sum of both numbers
func NewVector(num int) Vector {
	partition := getRandomIntBetweenRanges(1, num)
	return Vector{partition, num - partition}
}

func NewChord(color []Vector, baseNotes [2]Note) Chord {
	var notes []Note
	for _, v := range color {
		freq := baseNotes[0].Freq*float32(v[0]) + baseNotes[1].Freq*float32(v[1])
		note := FindNearestNote(freq, NotesCatalog)
		notes = append(notes, note)
	}
	return Chord{
		BaseNotes: baseNotes,
		Color:     color,
		Notes:     notes,
	}
}
