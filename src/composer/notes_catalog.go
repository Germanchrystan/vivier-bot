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
	Accidental Accidental
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

var NotesCatalog = []Note{
	{NoteNumber: 3, Octave: 2, Step: "C", Freq: 65.406, Accidental: GetAccidental(natural)},
	{NoteNumber: 4, Octave: 2, Step: "C", Freq: 67.323, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 5, Octave: 2, Step: "C", Freq: 69.296, Accidental: GetAccidental(sharp)},
	{NoteNumber: 6, Octave: 2, Step: "D", Freq: 71.326, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 7, Octave: 2, Step: "D", Freq: 73.416, Accidental: GetAccidental(natural)},
	{NoteNumber: 8, Octave: 2, Step: "D", Freq: 75.567, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 9, Octave: 2, Step: "D", Freq: 77.782, Accidental: GetAccidental(sharp)},
	{NoteNumber: 10, Octave: 2, Step: "E", Freq: 80.061, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 11, Octave: 2, Step: "E", Freq: 82.407, Accidental: GetAccidental(natural)},
	{NoteNumber: 12, Octave: 2, Step: "E", Freq: 84.822, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 13, Octave: 2, Step: "F", Freq: 87.307, Accidental: GetAccidental(natural)},
	{NoteNumber: 14, Octave: 2, Step: "F", Freq: 89.865, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 15, Octave: 2, Step: "F", Freq: 92.499, Accidental: GetAccidental(sharp)},
	{NoteNumber: 16, Octave: 2, Step: "G", Freq: 95.209, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 17, Octave: 2, Step: "G", Freq: 97.999, Accidental: GetAccidental(natural)},
	{NoteNumber: 18, Octave: 2, Step: "G", Freq: 100.870, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 19, Octave: 2, Step: "G", Freq: 103.826, Accidental: GetAccidental(sharp)},
	{NoteNumber: 20, Octave: 2, Step: "A", Freq: 106.869, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 21, Octave: 2, Step: "A", Freq: 110.000, Accidental: GetAccidental(natural)},
	{NoteNumber: 22, Octave: 2, Step: "A", Freq: 113.223, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 23, Octave: 2, Step: "A", Freq: 116.541, Accidental: GetAccidental(sharp)},
	{NoteNumber: 24, Octave: 2, Step: "B", Freq: 119.956, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 25, Octave: 2, Step: "B", Freq: 123.471, Accidental: GetAccidental(natural)},
	{NoteNumber: 26, Octave: 2, Step: "B", Freq: 127.089, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 27, Octave: 3, Step: "C", Freq: 130.813, Accidental: GetAccidental(natural)},
	{NoteNumber: 28, Octave: 3, Step: "C", Freq: 134.646, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 29, Octave: 3, Step: "C", Freq: 138.591, Accidental: GetAccidental(sharp)},
	{NoteNumber: 30, Octave: 3, Step: "D", Freq: 142.652, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 31, Octave: 3, Step: "D", Freq: 146.832, Accidental: GetAccidental(natural)},
	{NoteNumber: 32, Octave: 3, Step: "D", Freq: 151.135, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 33, Octave: 3, Step: "D", Freq: 155.563, Accidental: GetAccidental(sharp)},
	{NoteNumber: 34, Octave: 3, Step: "E", Freq: 160.122, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 35, Octave: 3, Step: "E", Freq: 164.814, Accidental: GetAccidental(natural)},
	{NoteNumber: 36, Octave: 3, Step: "E", Freq: 169.643, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 37, Octave: 3, Step: "F", Freq: 174.614, Accidental: GetAccidental(natural)},
	{NoteNumber: 38, Octave: 3, Step: "F", Freq: 179.731, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 39, Octave: 3, Step: "F", Freq: 184.997, Accidental: GetAccidental(sharp)},
	{NoteNumber: 40, Octave: 3, Step: "G", Freq: 190.418, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 41, Octave: 3, Step: "G", Freq: 195.998, Accidental: GetAccidental(natural)},
	{NoteNumber: 42, Octave: 3, Step: "G", Freq: 201.741, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 43, Octave: 3, Step: "G", Freq: 207.652, Accidental: GetAccidental(sharp)},
	{NoteNumber: 44, Octave: 3, Step: "A", Freq: 213.737, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 45, Octave: 3, Step: "A", Freq: 220.000, Accidental: GetAccidental(natural)},
	{NoteNumber: 46, Octave: 3, Step: "A", Freq: 226.446, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 47, Octave: 3, Step: "A", Freq: 233.082, Accidental: GetAccidental(sharp)},
	{NoteNumber: 48, Octave: 3, Step: "B", Freq: 239.912, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 49, Octave: 3, Step: "B", Freq: 246.942, Accidental: GetAccidental(natural)},
	{NoteNumber: 50, Octave: 3, Step: "B", Freq: 254.178, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 51, Octave: 4, Step: "C", Freq: 261.626, Accidental: GetAccidental(natural)},
	{NoteNumber: 52, Octave: 4, Step: "C", Freq: 269.292, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 53, Octave: 4, Step: "C", Freq: 277.183, Accidental: GetAccidental(sharp)},
	{NoteNumber: 54, Octave: 4, Step: "D", Freq: 285.305, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 55, Octave: 4, Step: "D", Freq: 293.665, Accidental: GetAccidental(natural)},
	{NoteNumber: 56, Octave: 4, Step: "D", Freq: 302.270, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 57, Octave: 4, Step: "D", Freq: 311.127, Accidental: GetAccidental(sharp)},
	{NoteNumber: 58, Octave: 4, Step: "E", Freq: 320.244, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 59, Octave: 4, Step: "E", Freq: 329.628, Accidental: GetAccidental(natural)},
	{NoteNumber: 60, Octave: 4, Step: "E", Freq: 339.286, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 61, Octave: 4, Step: "F", Freq: 349.228, Accidental: GetAccidental(natural)},
	{NoteNumber: 62, Octave: 4, Step: "F", Freq: 359.461, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 63, Octave: 4, Step: "F", Freq: 369.994, Accidental: GetAccidental(sharp)},
	{NoteNumber: 64, Octave: 4, Step: "G", Freq: 380.836, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 65, Octave: 4, Step: "G", Freq: 391.995, Accidental: GetAccidental(natural)},
	{NoteNumber: 66, Octave: 4, Step: "G", Freq: 403.482, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 67, Octave: 4, Step: "G", Freq: 415.305, Accidental: GetAccidental(sharp)},
	{NoteNumber: 68, Octave: 4, Step: "A", Freq: 427.474, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 69, Octave: 4, Step: "A", Freq: 440.000, Accidental: GetAccidental(natural)},
	{NoteNumber: 70, Octave: 4, Step: "A", Freq: 452.893, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 71, Octave: 4, Step: "A", Freq: 466.164, Accidental: GetAccidental(sharp)},
	{NoteNumber: 72, Octave: 4, Step: "B", Freq: 479.823, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 73, Octave: 4, Step: "B", Freq: 493.883, Accidental: GetAccidental(natural)},
	{NoteNumber: 74, Octave: 4, Step: "B", Freq: 508.355, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 75, Octave: 5, Step: "C", Freq: 523.251, Accidental: GetAccidental(natural)},
	{NoteNumber: 76, Octave: 5, Step: "C", Freq: 538.584, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 77, Octave: 5, Step: "C", Freq: 554.365, Accidental: GetAccidental(sharp)},
	{NoteNumber: 78, Octave: 5, Step: "D", Freq: 570.609, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 79, Octave: 5, Step: "D", Freq: 587.330, Accidental: GetAccidental(natural)},
	{NoteNumber: 80, Octave: 5, Step: "D", Freq: 604.540, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 81, Octave: 5, Step: "D", Freq: 622.254, Accidental: GetAccidental(sharp)},
	{NoteNumber: 82, Octave: 5, Step: "E", Freq: 640.487, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 83, Octave: 5, Step: "E", Freq: 659.255, Accidental: GetAccidental(natural)},
	{NoteNumber: 84, Octave: 5, Step: "E", Freq: 678.573, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 85, Octave: 5, Step: "F", Freq: 698.456, Accidental: GetAccidental(natural)},
	{NoteNumber: 86, Octave: 5, Step: "F", Freq: 718.923, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 87, Octave: 5, Step: "F", Freq: 739.989, Accidental: GetAccidental(sharp)},
	{NoteNumber: 88, Octave: 5, Step: "G", Freq: 761.672, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 89, Octave: 5, Step: "G", Freq: 783.991, Accidental: GetAccidental(natural)},
	{NoteNumber: 90, Octave: 5, Step: "G", Freq: 806.964, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 91, Octave: 5, Step: "G", Freq: 830.609, Accidental: GetAccidental(sharp)},
	{NoteNumber: 92, Octave: 5, Step: "A", Freq: 854.948, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 93, Octave: 5, Step: "A", Freq: 880.000, Accidental: GetAccidental(natural)},
	{NoteNumber: 94, Octave: 5, Step: "A", Freq: 905.786, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 95, Octave: 5, Step: "A", Freq: 932.328, Accidental: GetAccidental(sharp)},
	{NoteNumber: 96, Octave: 5, Step: "B", Freq: 959.647, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 97, Octave: 5, Step: "B", Freq: 987.767, Accidental: GetAccidental(natural)},
	{NoteNumber: 98, Octave: 5, Step: "B", Freq: 1016.710, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 99, Octave: 6, Step: "C", Freq: 1046.502, Accidental: GetAccidental(natural)},
	{NoteNumber: 100, Octave: 6, Step: "C", Freq: 1077.167, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 101, Octave: 6, Step: "C", Freq: 1108.731, Accidental: GetAccidental(sharp)},
	{NoteNumber: 102, Octave: 6, Step: "D", Freq: 1141.219, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 103, Octave: 6, Step: "D", Freq: 1174.659, Accidental: GetAccidental(natural)},
	{NoteNumber: 104, Octave: 6, Step: "D", Freq: 1209.079, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 105, Octave: 6, Step: "D", Freq: 1244.508, Accidental: GetAccidental(sharp)},
	{NoteNumber: 106, Octave: 6, Step: "E", Freq: 1280.975, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 107, Octave: 6, Step: "E", Freq: 1318.510, Accidental: GetAccidental(natural)},
	{NoteNumber: 108, Octave: 6, Step: "E", Freq: 1357.146, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 109, Octave: 6, Step: "F", Freq: 1396.913, Accidental: GetAccidental(natural)},
	{NoteNumber: 110, Octave: 6, Step: "F", Freq: 1437.846, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 111, Octave: 6, Step: "F", Freq: 1479.978, Accidental: GetAccidental(sharp)},
	{NoteNumber: 112, Octave: 6, Step: "G", Freq: 1523.344, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 113, Octave: 6, Step: "G", Freq: 1567.982, Accidental: GetAccidental(natural)},
	{NoteNumber: 114, Octave: 6, Step: "G", Freq: 1613.927, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 115, Octave: 6, Step: "G", Freq: 1661.219, Accidental: GetAccidental(sharp)},
	{NoteNumber: 116, Octave: 6, Step: "A", Freq: 1709.896, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 117, Octave: 6, Step: "A", Freq: 1760.000, Accidental: GetAccidental(natural)},
	{NoteNumber: 118, Octave: 6, Step: "A", Freq: 1811.572, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 119, Octave: 6, Step: "A", Freq: 1864.655, Accidental: GetAccidental(sharp)},
	{NoteNumber: 120, Octave: 6, Step: "B", Freq: 1919.294, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 121, Octave: 6, Step: "B", Freq: 1975.533, Accidental: GetAccidental(natural)},
	{NoteNumber: 122, Octave: 6, Step: "B", Freq: 2033.421, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 123, Octave: 7, Step: "C", Freq: 2093.005, Accidental: GetAccidental(natural)},
	{NoteNumber: 124, Octave: 7, Step: "C", Freq: 2154.334, Accidental: GetAccidental(quarterSharp)},
	{NoteNumber: 125, Octave: 7, Step: "C", Freq: 2217.461, Accidental: GetAccidental(sharp)},
	{NoteNumber: 126, Octave: 7, Step: "D", Freq: 2282.438, Accidental: GetAccidental(quarterFlat)},
	{NoteNumber: 127, Octave: 7, Step: "D", Freq: 2349.318, Accidental: GetAccidental(natural)},
}
