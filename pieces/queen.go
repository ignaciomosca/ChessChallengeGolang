package pieces

import (
	piece "ChessChallengeGolang/primitives"
	"ChessChallengeGolang/utils"
	"fmt"
)

type Queen struct {
	piece.Piece
}

func CreateQueen(row, col int) Queen {
	return Queen{piece.Piece{'Q', row, col}}
}

func (q Queen) Attacks(dest piece.Attacks) bool {
	sameRow := q.Row() == dest.Row()
	sameCol := q.Col() == dest.Col()
	sameDiag := utils.Abs(dest.Row()-q.Row()) == utils.Abs(dest.Col()-q.Col())
	return sameRow || sameCol || sameDiag
}

func (q Queen) Row() int { return q.Piece.Row }

func (q Queen) Col() int { return q.Piece.Col }

func (q Queen) Name() rune { return q.Piece.Name }

func (q Queen) ToString() string {
	r := string(q.Name())
	return fmt.Sprintf("%s%d%d", r, q.Row(), q.Col())
}
