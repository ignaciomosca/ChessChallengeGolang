package board

import "chess/primitives"

// Board Represents and MxN Chess board
type Board struct {
	M     int
	N     int
	board [][]rune
	used  map[primitives.Attacks]bool
}

// CreateBoard creates an empty board of MxN
func CreateBoard(m, n int) Board {
	board := make([][]rune, m)
	for i := range board {
		board[i] = make([]rune, n)
	}
	return Board{m, n, board, make(map[primitives.Attacks]bool)}
}

// CreateBoardWithPieces creates a board with the pieces being @used
func CreateBoardWithPieces(m, n int, used map[primitives.Attacks]bool) Board {
	board := make([][]rune, m)
	for i := range board {
		board[i] = make([]rune, n)
	}
	addUsedPieces(board, used)
	return Board{m, n, board, used}
}

// usedPieces pieces to be set in the board
func addUsedPieces(board [][]rune, used map[primitives.Attacks]bool) {
	for piece := range used {
		board[piece.Row()][piece.Col()] = piece.Name()
	}
}

func (board Board) ShowBoard() {
	for i := 0; i < board.M; i++ {
		for j := 0; j < board.N; j++ {
			println(string(board.board[i][j]))
		}
		println()
	}
	println()
}

func (board *Board) place(piece primitives.Attacks) Board {
	board.used[piece] = true
	return CreateBoardWithPieces(board.M, board.N, board.used)
}

// FindPiece locates a piece in a board with coordinates row,col
func (b *Board) FindPiece(row, col int) rune {
	result := b.board[row][col]
	if result == '\u0000' {
		return '_'
	} else {
		return result
	}
}

// IsSafe return true if no other piece in the board gets attacked by c and if c is not already placed
func (b *Board) IsSafe(c primitives.Attacks) bool {
	doesNotAttackOtherPieces := doesNotAttackOtherPieces(b.used, c)
	return doesNotAttackOtherPieces && b.notYetPlaced(c)
}

func doesNotAttackOtherPieces(usedPieces map[primitives.Attacks]bool, c primitives.Attacks) bool {
	for p := range usedPieces {
		if p.Attacks(c) || c.Attacks(p) {
			return false
		}
	}
	return true
}

func (b *Board) notYetPlaced(c primitives.Attacks) bool {
	return b.board[c.Row()][c.Col()] == '\u0000'
}
