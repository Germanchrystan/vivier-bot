package composer

import (
	"math"
	"math/rand"
	"time"
)

func FindNearestNote(Freq float32, notes []Note) Note {
	notesLen := len(notes)                                // Length of notes slice
	notesMiddle := int(math.Floor(float64(notesLen) / 2)) // Half way point of the array (flooring)

	if Freq > notes[notesMiddle].Freq { // If Freq is bigger than the halfway point Freq
		// Is Freq overflows the notes catalog, reduce it by one octave
		if Freq > notes[len(notes)-1].Freq {
			return FindNearestNote(Freq/2, notes)
		}

		if Freq < notes[notesMiddle+1].Freq { // If Freq is smaller than halfway point + 1 Freq
			// Get smaller difference
			if Freq-notes[notesMiddle].Freq < Freq-notes[notesMiddle+1].Freq {
				return notes[notesMiddle]
			}
			return notes[notesMiddle+1]
		}
		// Else, if Freq is not smaller than halfway point + 1 Freq,
		// return recursive function for the upper half
		return FindNearestNote(Freq, notes[notesMiddle:])
	} else {
		if Freq > notes[notesMiddle-1].Freq {
			if Freq-notes[notesMiddle].Freq < Freq-notes[notesMiddle-1].Freq {
				return notes[notesMiddle]
			}
			return notes[notesMiddle-1]
		}
		return FindNearestNote(Freq, notes[:notesMiddle])
	}
}

// Get a random number between to values
func getRandomIntBetweenRanges(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn((max)-min) + min
}
