package pieces

import (
	piece "chess/primitives"
)

type Knight struct {
	piece.Piece
}

func CreateKnight(row, col int) Knight {
	return Knight{piece.Piece{"N", row, col}}
}

func (k Knight) Attacks() bool {
	return false
}

func (k Knight) Row() int { return k.Piece.Row }

func (k Knight) Col() int { return k.Piece.Col }

func (k Knight) Name() string { return k.Piece.Name }
