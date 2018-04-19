package primitives

type Piece struct {
	Name rune
	Row int
	Col int
}

type Attacks interface {
	Attacks(piece Attacks) bool
	Row() int
	Col() int
	Name() rune
}