package primitives

type Piece struct {
	Name string
	Row int
	Col int
}

type Attacks interface {
	Attacks() bool
	Row() int
	Col() int
	Name() string
}