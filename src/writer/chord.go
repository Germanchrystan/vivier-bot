package writer

type Matrix struct {
	XNote  Note
	YNote  Note
	Matrix [][]Note
}

func newMatrix(x, y Note) Matrix {
	matrix := Matrix{
		XNote:  x,
		YNote:  y,
		Matrix: [][]Note{},
	}

	return matrix
}

type Coordinate [2]int

type Chord struct {
	Coors []Coordinate
	Notes []Note
}
