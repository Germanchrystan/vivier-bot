package notes

type accidentalCase int

const (
	natural accidentalCase = iota
	quarterSharp
	sharp
	quarterFlat
	flat
)

type Note struct {
	noteNumber int
	accidental Accidental
	step       string
	octave     int
	freq       float32
}

type Accidental struct {
	accidental string
	alter      float32
}

func GetAccidental(accidental accidentalCase) Accidental {
	switch accidental {
	case natural:
		return Accidental{
			alter:      0,
			accidental: "",
		}
	case quarterSharp:
		return Accidental{
			alter:      0.5,
			accidental: "quarter-sharp",
		}
	case sharp:
		return Accidental{
			alter:      1,
			accidental: "sharp",
		}
	case quarterFlat:
		return Accidental{
			alter:      -0.5,
			accidental: "quarter-flat",
		}
	case flat:
		return Accidental{
			alter:      -1,
			accidental: "flat",
		}
	default:
		return Accidental{}
	}
}

var notesArray = []Note{
	{noteNumber: 3, octave: 2, step: "C", freq: 65.406, accidental: GetAccidental(natural)},
	{noteNumber: 4, octave: 2, step: "C", freq: 67.323, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 5, octave: 2, step: "C", freq: 69.296, accidental: GetAccidental(sharp)},
	{noteNumber: 6, octave: 2, step: "D", freq: 71.326, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 7, octave: 2, step: "D", freq: 73.416, accidental: GetAccidental(natural)},
	{noteNumber: 8, octave: 2, step: "D", freq: 75.567, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 9, octave: 2, step: "D", freq: 77.782, accidental: GetAccidental(sharp)},
	{noteNumber: 10, octave: 2, step: "E", freq: 80.061, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 11, octave: 2, step: "E", freq: 82.407, accidental: GetAccidental(natural)},
	{noteNumber: 12, octave: 2, step: "E", freq: 84.822, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 13, octave: 2, step: "F", freq: 87.307, accidental: GetAccidental(natural)},
	{noteNumber: 14, octave: 2, step: "F", freq: 89.865, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 15, octave: 2, step: "F", freq: 92.499, accidental: GetAccidental(sharp)},
	{noteNumber: 16, octave: 2, step: "G", freq: 95.209, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 17, octave: 2, step: "G", freq: 97.999, accidental: GetAccidental(natural)},
	{noteNumber: 18, octave: 2, step: "G", freq: 100.870, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 19, octave: 2, step: "G", freq: 103.826, accidental: GetAccidental(sharp)},
	{noteNumber: 20, octave: 2, step: "A", freq: 106.869, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 21, octave: 2, step: "A", freq: 110.000, accidental: GetAccidental(natural)},
	{noteNumber: 22, octave: 2, step: "A", freq: 113.223, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 23, octave: 2, step: "A", freq: 116.541, accidental: GetAccidental(sharp)},
	{noteNumber: 24, octave: 2, step: "B", freq: 119.956, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 25, octave: 2, step: "B", freq: 123.471, accidental: GetAccidental(natural)},
	{noteNumber: 26, octave: 2, step: "B", freq: 127.089, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 27, octave: 3, step: "C", freq: 130.813, accidental: GetAccidental(natural)},
	{noteNumber: 28, octave: 3, step: "C", freq: 134.646, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 29, octave: 3, step: "C", freq: 138.591, accidental: GetAccidental(sharp)},
	{noteNumber: 30, octave: 3, step: "D", freq: 142.652, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 31, octave: 3, step: "D", freq: 146.832, accidental: GetAccidental(natural)},
	{noteNumber: 32, octave: 3, step: "D", freq: 151.135, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 33, octave: 3, step: "D", freq: 155.563, accidental: GetAccidental(sharp)},
	{noteNumber: 34, octave: 3, step: "E", freq: 160.122, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 35, octave: 3, step: "E", freq: 164.814, accidental: GetAccidental(natural)},
	{noteNumber: 36, octave: 3, step: "E", freq: 169.643, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 37, octave: 3, step: "F", freq: 174.614, accidental: GetAccidental(natural)},
	{noteNumber: 38, octave: 3, step: "F", freq: 179.731, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 39, octave: 3, step: "F", freq: 184.997, accidental: GetAccidental(sharp)},
	{noteNumber: 40, octave: 3, step: "G", freq: 190.418, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 41, octave: 3, step: "G", freq: 195.998, accidental: GetAccidental(natural)},
	{noteNumber: 42, octave: 3, step: "G", freq: 201.741, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 43, octave: 3, step: "G", freq: 207.652, accidental: GetAccidental(sharp)},
	{noteNumber: 44, octave: 3, step: "A", freq: 213.737, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 45, octave: 3, step: "A", freq: 220.000, accidental: GetAccidental(natural)},
	{noteNumber: 46, octave: 3, step: "A", freq: 226.446, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 47, octave: 3, step: "A", freq: 233.082, accidental: GetAccidental(sharp)},
	{noteNumber: 48, octave: 3, step: "B", freq: 239.912, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 49, octave: 3, step: "B", freq: 246.942, accidental: GetAccidental(natural)},
	{noteNumber: 50, octave: 3, step: "B", freq: 254.178, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 51, octave: 4, step: "C", freq: 261.626, accidental: GetAccidental(natural)},
	{noteNumber: 52, octave: 4, step: "C", freq: 269.292, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 53, octave: 4, step: "C", freq: 277.183, accidental: GetAccidental(sharp)},
	{noteNumber: 54, octave: 4, step: "D", freq: 285.305, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 55, octave: 4, step: "D", freq: 293.665, accidental: GetAccidental(natural)},
	{noteNumber: 56, octave: 4, step: "D", freq: 302.270, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 57, octave: 4, step: "D", freq: 311.127, accidental: GetAccidental(sharp)},
	{noteNumber: 58, octave: 4, step: "E", freq: 320.244, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 59, octave: 4, step: "E", freq: 329.628, accidental: GetAccidental(natural)},
	{noteNumber: 60, octave: 4, step: "E", freq: 339.286, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 61, octave: 4, step: "F", freq: 349.228, accidental: GetAccidental(natural)},
	{noteNumber: 62, octave: 4, step: "F", freq: 359.461, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 63, octave: 4, step: "F", freq: 369.994, accidental: GetAccidental(sharp)},
	{noteNumber: 64, octave: 4, step: "G", freq: 380.836, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 65, octave: 4, step: "G", freq: 391.995, accidental: GetAccidental(natural)},
	{noteNumber: 66, octave: 4, step: "G", freq: 403.482, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 67, octave: 4, step: "G", freq: 415.305, accidental: GetAccidental(sharp)},
	{noteNumber: 68, octave: 4, step: "A", freq: 427.474, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 69, octave: 4, step: "A", freq: 440.000, accidental: GetAccidental(natural)},
	{noteNumber: 70, octave: 4, step: "A", freq: 452.893, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 71, octave: 4, step: "A", freq: 466.164, accidental: GetAccidental(sharp)},
	{noteNumber: 72, octave: 4, step: "B", freq: 479.823, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 73, octave: 4, step: "B", freq: 493.883, accidental: GetAccidental(natural)},
	{noteNumber: 74, octave: 4, step: "B", freq: 508.355, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 75, octave: 5, step: "C", freq: 523.251, accidental: GetAccidental(natural)},
	{noteNumber: 76, octave: 5, step: "C", freq: 538.584, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 77, octave: 5, step: "C", freq: 554.365, accidental: GetAccidental(sharp)},
	{noteNumber: 78, octave: 5, step: "D", freq: 570.609, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 79, octave: 5, step: "D", freq: 587.330, accidental: GetAccidental(natural)},
	{noteNumber: 80, octave: 5, step: "D", freq: 604.540, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 81, octave: 5, step: "D", freq: 622.254, accidental: GetAccidental(sharp)},
	{noteNumber: 82, octave: 5, step: "E", freq: 640.487, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 83, octave: 5, step: "E", freq: 659.255, accidental: GetAccidental(natural)},
	{noteNumber: 84, octave: 5, step: "E", freq: 678.573, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 85, octave: 5, step: "F", freq: 698.456, accidental: GetAccidental(natural)},
	{noteNumber: 86, octave: 5, step: "F", freq: 718.923, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 87, octave: 5, step: "F", freq: 739.989, accidental: GetAccidental(sharp)},
	{noteNumber: 88, octave: 5, step: "G", freq: 761.672, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 89, octave: 5, step: "G", freq: 783.991, accidental: GetAccidental(natural)},
	{noteNumber: 90, octave: 5, step: "G", freq: 806.964, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 91, octave: 5, step: "G", freq: 830.609, accidental: GetAccidental(sharp)},
	{noteNumber: 92, octave: 5, step: "A", freq: 854.948, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 93, octave: 5, step: "A", freq: 880.000, accidental: GetAccidental(natural)},
	{noteNumber: 94, octave: 5, step: "A", freq: 905.786, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 95, octave: 5, step: "A", freq: 932.328, accidental: GetAccidental(sharp)},
	{noteNumber: 96, octave: 5, step: "B", freq: 959.647, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 97, octave: 5, step: "B", freq: 987.767, accidental: GetAccidental(natural)},
	{noteNumber: 98, octave: 5, step: "B", freq: 1016.710, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 99, octave: 6, step: "C", freq: 1046.502, accidental: GetAccidental(natural)},
	{noteNumber: 100, octave: 6, step: "C", freq: 1077.167, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 101, octave: 6, step: "C", freq: 1108.731, accidental: GetAccidental(sharp)},
	{noteNumber: 102, octave: 6, step: "D", freq: 1141.219, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 103, octave: 6, step: "D", freq: 1174.659, accidental: GetAccidental(natural)},
	{noteNumber: 104, octave: 6, step: "D", freq: 1209.079, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 105, octave: 6, step: "D", freq: 1244.508, accidental: GetAccidental(sharp)},
	{noteNumber: 106, octave: 6, step: "E", freq: 1280.975, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 107, octave: 6, step: "E", freq: 1318.510, accidental: GetAccidental(natural)},
	{noteNumber: 108, octave: 6, step: "E", freq: 1357.146, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 109, octave: 6, step: "F", freq: 1396.913, accidental: GetAccidental(natural)},
	{noteNumber: 110, octave: 6, step: "F", freq: 1437.846, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 111, octave: 6, step: "F", freq: 1479.978, accidental: GetAccidental(sharp)},
	{noteNumber: 112, octave: 6, step: "G", freq: 1523.344, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 113, octave: 6, step: "G", freq: 1567.982, accidental: GetAccidental(natural)},
	{noteNumber: 114, octave: 6, step: "G", freq: 1613.927, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 115, octave: 6, step: "G", freq: 1661.219, accidental: GetAccidental(sharp)},
	{noteNumber: 116, octave: 6, step: "A", freq: 1709.896, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 117, octave: 6, step: "A", freq: 1760.000, accidental: GetAccidental(natural)},
	{noteNumber: 118, octave: 6, step: "A", freq: 1811.572, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 119, octave: 6, step: "A", freq: 1864.655, accidental: GetAccidental(sharp)},
	{noteNumber: 120, octave: 6, step: "B", freq: 1919.294, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 121, octave: 6, step: "B", freq: 1975.533, accidental: GetAccidental(natural)},
	{noteNumber: 122, octave: 6, step: "B", freq: 2033.421, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 123, octave: 7, step: "C", freq: 2093.005, accidental: GetAccidental(natural)},
	{noteNumber: 124, octave: 7, step: "C", freq: 2154.334, accidental: GetAccidental(quarterSharp)},
	{noteNumber: 125, octave: 7, step: "C", freq: 2217.461, accidental: GetAccidental(sharp)},
	{noteNumber: 126, octave: 7, step: "D", freq: 2282.438, accidental: GetAccidental(quarterFlat)},
	{noteNumber: 127, octave: 7, step: "D", freq: 2349.318, accidental: GetAccidental(natural)},
}
