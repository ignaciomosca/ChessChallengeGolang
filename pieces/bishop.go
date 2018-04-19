package pieces

import piece "chess/primitives"

type Bishop struct {
	piece.Piece
}

func CreateBishop(row, col int) Bishop {
	return Bishop{piece.Piece{"B", row, col}}
}

func (bishop Bishop) Attacks() bool {
	return false
}

func (k Bishop) Row() int { return k.Piece.Row }

func (k Bishop) Col() int { return k.Piece.Col }

func (k Bishop) Name() string { return k.Piece.Name }
