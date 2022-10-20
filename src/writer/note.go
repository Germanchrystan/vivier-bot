package writer

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/composer"
)

func engraveNote(note composer.Note, isFirstChordNote bool) string {
	var chordTag string
	if !isFirstChordNote {
		chordTag = "<chord/>"
	}
	accidental := composer.GetAccidental(note.Accidental)
	return fmt.Sprintf("<note>%s\n<pitch>\n<step>%s</step>\n<alter>%f</alter>\n<octave>%d</octave>\n</pitch>\n<duration>1</duration>\n<voice>1</voice><type>quarter</type><accidental>%s</accidental><stem>up</stem></note>",
		chordTag,
		note.Step,
		accidental.Alter,
		note.Octave,
		accidental.Accidental,
	)
}

func engraveNotes(notes []composer.Note) string {
	var result string
	for i, note := range notes {
		isFirstChordNote := i == 0
		result += engraveNote(note, isFirstChordNote)
	}
	return result
}
