package board

import "chess/primitives"

// Board Represents and MxN Chess board
type Board struct {
	M     int
	N     int
	board [][]string
	used  map[primitives.Attacks]bool
}

// CreateBoard creates an empty board of MxN
func CreateBoard(m, n int) Board {
	board := make([][]string, m)
	for i := range board {
		board[i] = make([]string, n)
	}
	return Board{m, n, board, make(map[primitives.Attacks]bool)}
}

// CreateBoardWithPieces creates a board with the pieces being @used
func CreateBoardWithPieces(m, n int, used map[primitives.Attacks]bool) Board {
	board := make([][]string, m)
	for i := range board {
		board[i] = make([]string, n)
	}
	return Board{m, n, board, used}
}

// usedPieces pieces to be set in the board
func addUsedPieces(board [][]string, used map[primitives.Attacks]bool) {
	for piece := range used {
		board[piece.Row()][piece.Col()] = piece.Name()
	}
}

func (board Board) ShowBoard() {
	for i := 0; i < board.M; i++ {
		for j := 0; j < board.N; j++ {
			println(board.board[i][j])
		}
		println()
	}
	println()
}

func (board Board) place(piece primitives.Attacks) Board {
	board.used[piece] = true
	return CreateBoardWithPieces(board.M, board.N, board.used)
}
