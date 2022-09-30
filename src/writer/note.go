package notes

import "fmt"

func generateNote(note Note) string {
	return fmt.Sprintf("<note>\n<pitch>\n<step>%s</step>\n<alter>%f</alter>\n<octave>%d</octave>\n</pitch>\n<duration>1</duration>\n<voice>1</voice><type>quarter</type><accidental>%s</accidental><stem>up</stem></note>", note.step, note.accidental.alter, note.octave, note.accidental.accidental)
}
