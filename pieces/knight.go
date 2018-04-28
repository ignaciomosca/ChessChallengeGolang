package pieces

import (
	piece "chess/primitives"
	"fmt"
)

type Knight struct {
	piece.Piece
}

func CreateKnight(row, col int) Knight {
	return Knight{piece.Piece{'N', row, col}}
}

func (k Knight) Attacks(dest piece.Attacks) bool {
	row := k.Row()
	col := k.Col()
	xMoves := []int{1, 2, 2, 1, -1, -2, -2, -1}
	yMoves := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	possibleMoves := make(map[piece.Attacks]bool)
	for i := 0; i < 8; i++ {
		destRow := row + xMoves[i]
		destCol := col + yMoves[i]
		if destRow > 0 && destCol > 0 {
			createPiece := CreatePiece(dest.Name(), destRow, destCol)
			possibleMoves[createPiece] = true
		}
	}
	return possibleMoves[dest]
}

func (k Knight) Row() int { return k.Piece.Row }

func (k Knight) Col() int { return k.Piece.Col }

func (k Knight) Name() rune { return k.Piece.Name }

func (k Knight) ToString() string {
	r := string(k.Name())
	return fmt.Sprintf("%s%d%d", r, k.Row(), k.Col())
}
