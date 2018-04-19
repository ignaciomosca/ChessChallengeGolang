package pieces

import piece "chess/primitives"

type King struct {
	piece.Piece
}

func CreateKing(row, col int) King {
	return King{piece.Piece{'K', row, col}}
}

func (k King) Attacks(p piece.Attacks) bool {
	row := k.Row()
	col := k.Col()
	xMoves := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	yMoves := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	possibleMoves := make(map[piece.Attacks]bool)
	for i := 0; i < 8; i++ {
		destRow := row + xMoves[i]
		destCol := col + yMoves[i]
		if destRow > 0 && destCol > 0 {
			createPiece := CreatePiece(p, destRow, destCol)
			possibleMoves[createPiece] = true
		}
	}

	return possibleMoves[p]
}

func (k King) Row() int { return k.Piece.Row }

func (k King) Col() int { return k.Piece.Col }

func (k King) Name() rune { return k.Piece.Name }
