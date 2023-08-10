package pieces

import (
	piece "chesschallengegolang/primitives"
	"fmt"
)

type Rook struct {
	piece.Piece
}

func CreateRook(row, col int) Rook {
	return Rook{piece.Piece{'R', row, col}}
}

func (r Rook) Attacks(dest piece.Attacks) bool {
	sameRow := r.Row() == dest.Row()
	sameCol := r.Col() == dest.Col()
	return sameRow || sameCol
}

func (r Rook) Row() int { return r.Piece.Row }

func (r Rook) Col() int { return r.Piece.Col }

func (r Rook) Name() rune { return r.Piece.Name }

func (r Rook) ToString() string {
	t := string(r.Name())
	return fmt.Sprintf("%s%d%d", t, r.Row(), r.Col())
}
