package writer

import "math"

func FindNearestNote(freq float32, notes []Note) int {
	notesLen := len(notes)                                // Length of notes slice
	notesMiddle := int(math.Floor(float64(notesLen) / 2)) // Half way point of the array (flooring)

	if freq > notes[notesMiddle].freq { // If freq is bigger than the halfway point freq
		if freq < notes[notesMiddle+1].freq { // If freq is smaller than halfway point + 1 freq
			// Get smaller difference
			if freq-notes[notesMiddle].freq < freq-notes[notesMiddle+1].freq {
				return notes[notesMiddle].noteNumber
			}
			return notes[notesMiddle+1].noteNumber
		}
		// Else, if freq is not smaller than halfway point + 1 freq,
		// return recursive function for the upper half
		return FindNearestNote(freq, notes[notesMiddle:])
	} else {
		if freq > notes[notesMiddle-1].freq {
			if freq-notes[notesMiddle].freq < freq-notes[notesMiddle-1].freq {
				return notes[notesMiddle].noteNumber
			}
			return notes[notesMiddle-1].noteNumber
		}
		return FindNearestNote(freq, notes[:notesMiddle])
	}
}
