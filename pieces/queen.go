package pieces

import piece "chess/primitives"

type Queen struct {
	piece.Piece
}

func CreateQueen(row, col int) Queen {
	return Queen{piece.Piece{'Q', row, col}}
}

func (queen Queen) Attacks(dest piece.Attacks) bool {
	return false
}

func (q Queen) Row() int { return q.Piece.Row }

func (q Queen) Col() int { return q.Piece.Col }

func (q Queen) Name() rune { return q.Piece.Name }
