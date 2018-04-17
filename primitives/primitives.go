package primitives

type Piece struct {
	Name string
	Row int
	Col int
}

func CreatePiece(name string, row, col int) Piece {
	return Piece {name , row, col}
}

type attacks interface {
	Attacks() bool
}