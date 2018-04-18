package primitives

type Piece struct {
	Name string
	Row int
	Col int
}

type attacks interface {
	Attacks() bool
}