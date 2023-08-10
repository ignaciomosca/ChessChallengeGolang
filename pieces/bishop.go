package pieces

import (
	piece "chesschallengegolang/primitives"
	"chesschallengegolang/utils"
	"fmt"
)

// Bishop represents a Bishop in a Chessboard
type Bishop struct {
	piece.Piece
}

func CreateBishop(row, col int) Bishop {
	return Bishop{piece.Piece{'B', row, col}}
}

func (b Bishop) Attacks(dest piece.Attacks) bool {
	return utils.Abs(dest.Row()-b.Row()) == utils.Abs(dest.Col()-b.Col())
}

func (b Bishop) Row() int { return b.Piece.Row }

func (b Bishop) Col() int { return b.Piece.Col }

func (b Bishop) Name() rune { return b.Piece.Name }

func (b Bishop) ToString() string {
	r := string(b.Name())
	return fmt.Sprintf("%s%d%d", r, b.Row(), b.Col())
}
