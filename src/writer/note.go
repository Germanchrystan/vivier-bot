package writer

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/composer"
)

func generateNote(note composer.Note) string {
	return fmt.Sprintf("<note>\n<pitch>\n<step>%s</step>\n<alter>%f</alter>\n<octave>%d</octave>\n</pitch>\n<duration>1</duration>\n<voice>1</voice><type>quarter</type><accidental>%s</accidental><stem>up</stem></note>",
		note.Step,
		note.Accidental.Alter,
		note.Octave,
		note.Accidental.Accidental,
	)
}
