package writer

import (
	"fmt"
	"strings"

	"github.com/Germanchrystan/vivier-forest/src/composer"
)

func engraveChord(chord composer.Chord) (baseNotes, upperNotes string) {
	return engraveNotes([]composer.Note{chord.BaseNotes[0], chord.BaseNotes[1]}), // Base notes
		engraveNotes(chord.Notes) // Upper notes
}

func EngraveProgression(progression []composer.Chord) (string, string) {
	var basePart, upperPart string
	measure := 1
	for index, chord := range progression {
		var measureSeparator string
		var measureSeparatorBase, measureSeparatorUpper string
		if index == 0 {
			measureSeparator = fmt.Sprintf("<measure number=\"%d\">", measure)

			headerUpper := "<print><system-layout><system-margins><left-margin>0.00</left-margin><right-margin>-0.00</right-margin></system-margins><top-system-distance>170.00</top-system-distance></system-layout></print><attributes><divisions>1</divisions><key><fifths>0</fifths></key><time><beats>4</beats><beat-type>4</beat-type></time><clef><sign>G</sign><line>2</line></clef></attributes>\n"
			headerBase := "<print><staff-layout number=\"1\"><staff-distance>65.00</staff-distance></staff-layout></print><attributes><divisions>1</divisions><key><fifths>0</fifths></key><time><beats>4</beats><beat-type>4</beat-type></time><clef><sign>F</sign><line>4</line></clef></attributes>\n"

			measureSeparatorBase, measureSeparatorUpper = measureSeparator+headerBase, measureSeparator+headerUpper
		} else if index%4 == 0 {
			measure++
			measureSeparator = fmt.Sprintf("</measure>><measure number=\"%d\">", measure)
			measureSeparatorBase, measureSeparatorUpper = measureSeparator, measureSeparator
		} else {
			measureSeparatorBase, measureSeparatorUpper = "", ""
		}

		bp, up := engraveChord(chord)
		basePart += measureSeparatorBase + bp
		upperPart += measureSeparatorUpper + up

		// End of progression
		if index == len(progression)-1 {
			remainingBeatQuantity := 4 - index%4
			remainingBeats := strings.Repeat("<note><rest/><duration>1</duration><voice>1</voice><type>quarter</type></note>\n", remainingBeatQuantity)

			basePart += remainingBeats + "</measure>"
			upperPart += remainingBeats + "</measure>"
		}
	}
	return basePart, upperPart
}
