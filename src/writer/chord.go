package writer

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/composer"
)

func engraveChord(chord composer.Chord) (baseNotes, upperNotes string) {
	return engraveNotes([]composer.Note{chord.BaseNotes[0], chord.BaseNotes[1]}), // Base notes
		engraveNotes(chord.Notes) // Upper notes
}

func engraveProgression(progression []composer.Chord) (string, string) {
	var basePart, upperPart string
	measure := 1
	for index, chord := range progression {
		var measureSeparator string
		if index == 0 {
			measureSeparator = fmt.Sprintf("<measure number=\"%d\">", measure)
		} else if index%4 == 0 {
			measure++
			measureSeparator = fmt.Sprintf("</measure>><measure number=\"%d\">", measure)
		} else {
			measureSeparator = ""
		}
		bp, up := engraveChord(chord)
		basePart += measureSeparator + bp
		upperPart += measureSeparator + up

	}
	return basePart, upperPart
}
