package writer

import (
	"fmt"
	"os"
)

func EngraveParts(baseProgression, upperProgression string) string {
	var basePart, upperPart string

	partList := "<part-list>\n<score-part id=\"P1\"><part-name print-object=\"no\">Piano</part-name><score-instrument id=\"P1-I1\">\n<instrument-name>Piano</instrument-name>\n</score-instrument><midi-device id=\"P1-I1\" port=\"1\"></midi-device><midi-instrument id=\"P1-I1\"><midi-channel>1</midi-channel><midi-program>1</midi-program><volume>78.7402</volume><pan>0</pan></midi-instrument></score-part><score-part id=\"P2\"><part-name print-object=\"no\">Piano</part-name>\n<score-instrument id=\"P2-I1\"><instrument-name>Piano</instrument-name></score-instrument><midi-device id=\"P2-I1\" port=\"1\"></midi-device><midi-instrument id=\"P2-I1\"><midi-channel>2</midi-channel><midi-program>1</midi-program><volume>78.7402</volume><pan>0</pan></midi-instrument></score-part></part-list>"
	upperPart = fmt.Sprintf("<part id=\"P1\">%s</part>\n", upperProgression)
	basePart = fmt.Sprintf("<part id=\"P2\">%s</part>\n", baseProgression)

	return partList + upperPart + basePart
}

func EngraveFile(parts string) string {
	return fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE score-partwise PUBLIC \"-//Recordare//DTD MusicXML 3.1 Partwise//EN\" \"http://www.musicxml.org/dtds/partwise.dtd\">\n<score-partwise version=\"3.1\">\n<identification><encoding>\n<software>MuseScore 3.2.3</software>\n<encoding-date>2022-10-22</encoding-date>\n<supports element=\"accidental\" type=\"yes\"/><supports element=\"beam\" type=\"yes\"/><supports element=\"print\" attribute=\"new-page\" type=\"yes\" value=\"yes\"/><supports element=\"print\" attribute=\"new-system\" type=\"yes\" value=\"yes\"/><supports element=\"stem\" type=\"yes\"/></encoding></identification>\n<defaults><scaling>\n<millimeters>7.05556</millimeters>\n<tenths>40</tenths>\n</scaling>\n<page-layout><page-height>1683.78</page-height><page-width>1190.55</page-width><page-margins type=\"even\"><left-margin>56.6929</left-margin><right-margin>56.6929</right-margin><top-margin>56.6929</top-margin><bottom-margin>113.386</bottom-margin></page-margins>\n<page-margins type=\"odd\"><left-margin>56.6929</left-margin><right-margin>56.6929</right-margin><top-margin>56.6929</top-margin><bottom-margin>113.386</bottom-margin></page-margins></page-layout>\n<word-font font-family=\"FreeSerif\" font-size=\"10\"/><lyric-font font-family=\"FreeSerif\" font-size=\"11\"/></defaults><credit page=\"1\">\n<credit-words default-x=\"595.275\" default-y=\"1627.09\" justify=\"center\" valign=\"top\" font-size=\"24\">EXAMPLE</credit-words>\n<credit-words></credit-words></credit>\n<credit page=\"1\"><credit-words default-x=\"1133.86\" default-y=\"1527.09\" justify=\"right\" valign=\"bottom\" font-size=\"12\">VivierBot</credit-words>\n</credit>%s</score-partwise>\n", parts)
}

func CreateFile(file string) {
	f, err := os.Create("progression.musicxml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(file)
}
