package pieces

import (
	piece "chess/primitives"
	"chess/utils"
)

// Bishop represents a Bishop in a Chessboard
type Bishop struct {
	piece.Piece
}

func CreateBishop(row, col int) Bishop {
	return Bishop{piece.Piece{'B', row, col}}
}

func (b Bishop) Attacks(dest piece.Attacks) bool {
	return utils.Abs(dest.Row() - b.Row()) == utils.Abs(dest.Col() - b.Col());
}

func (b Bishop) Row() int { return b.Piece.Row }

func (b Bishop) Col() int { return b.Piece.Col }

func (b Bishop) Name() rune { return b.Piece.Name }
