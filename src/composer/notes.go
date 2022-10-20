package composer

type AccidentalCase int

const (
	natural AccidentalCase = iota
	quarterSharp
	sharp
	quarterFlat
	flat
)

type Note struct {
	NoteNumber int
	Accidental AccidentalCase
	Step       string
	Octave     int
	Freq       float32
}

type Accidental struct {
	Accidental string
	Alter      float32
}

func GetAccidental(acc AccidentalCase) Accidental {
	switch acc {
	case natural:
		return Accidental{
			Alter:      0,
			Accidental: "",
		}
	case quarterSharp:
		return Accidental{
			Alter:      0.5,
			Accidental: "quarter-sharp",
		}
	case sharp:
		return Accidental{
			Alter:      1,
			Accidental: "sharp",
		}
	case quarterFlat:
		return Accidental{
			Alter:      -0.5,
			Accidental: "quarter-flat",
		}
	case flat:
		return Accidental{
			Alter:      -1,
			Accidental: "flat",
		}
	default:
		return Accidental{}
	}
}
