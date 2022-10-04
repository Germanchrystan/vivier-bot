package composer

type Matrix struct {
	XNote Note
	YNote Note
}

func newMatrix(x, y Note) Matrix {
	matrix := Matrix{
		XNote: x,
		YNote: y,
	}

	return matrix
}

type Coordinate [2]int

type Chord struct {
	Coors []Coordinate
	Notes []Note
}
