package board

import piece "chess/primitives"

type Board struct {
	M    int
	N    int
	board [][]string
	used map[piece.Piece]bool
}

func CreateBoardWithPieces(m, n int, used map[piece.Piece]bool) Board{
	board := make([][]string, m)
	for i := range board {
		board[i] = make([]string, n)
	}
	return Board{m, n, board, used}
}
