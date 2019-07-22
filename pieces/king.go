package pieces

import (
	piece "ChessChallengeGolang/primitives"
	"fmt"
)

type King struct {
	piece.Piece
}

func CreateKing(row, col int) King {
	return King{piece.Piece{'K', row, col}}
}

func (k King) Attacks(dest piece.Attacks) bool {
	row := k.Row()
	col := k.Col()
	xMoves := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	yMoves := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	for i := 0; i < 8; i++ {
		destRow := row + xMoves[i]
		destCol := col + yMoves[i]
		if dest.Row() == destRow && dest.Col() == destCol {
			return true
		}
	}

	return false
}

func (k King) Row() int { return k.Piece.Row }

func (k King) Col() int { return k.Piece.Col }

func (k King) Name() rune { return k.Piece.Name }

func (k King) ToString() string {
	r := string(k.Name())
	return fmt.Sprintf("%s%d%d", r, k.Row(), k.Col())
}
