package composer

import "math"

func FindNearestNote(Freq float32, notes []Note) int {
	notesLen := len(notes)                                // Length of notes slice
	notesMiddle := int(math.Floor(float64(notesLen) / 2)) // Half way point of the array (flooring)

	if Freq > notes[notesMiddle].Freq { // If Freq is bigger than the halfway point Freq
		if Freq < notes[notesMiddle+1].Freq { // If Freq is smaller than halfway point + 1 Freq
			// Get smaller difference
			if Freq-notes[notesMiddle].Freq < Freq-notes[notesMiddle+1].Freq {
				return notes[notesMiddle].NoteNumber
			}
			return notes[notesMiddle+1].NoteNumber
		}
		// Else, if Freq is not smaller than halfway point + 1 Freq,
		// return recursive function for the upper half
		return FindNearestNote(Freq, notes[notesMiddle:])
	} else {
		if Freq > notes[notesMiddle-1].Freq {
			if Freq-notes[notesMiddle].Freq < Freq-notes[notesMiddle-1].Freq {
				return notes[notesMiddle].NoteNumber
			}
			return notes[notesMiddle-1].NoteNumber
		}
		return FindNearestNote(Freq, notes[:notesMiddle])
	}
}
