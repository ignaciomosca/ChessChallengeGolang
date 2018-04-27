package board

import "chess/primitives"

// Board Represents and MxN Chess Board
type Board struct {
	M     int
	N     int
	Board [][]rune
	Used  []primitives.Attacks
}

// CreateBoard creates an empty Board of MxN
func CreateBoard(m, n int) Board {
	board := make([][]rune, m)
	for i := range board {
		board[i] = make([]rune, n)
	}
	initializeBoard(board, m, n)
	var used []primitives.Attacks
	return Board{m, n, board, used}
}

// CreateBoardWithPieces creates a Board with the pieces being @Used
func CreateBoardWithPieces(m, n int, used []primitives.Attacks) Board {
	board := make([][]rune, m)
	for i := range board {
		board[i] = make([]rune, n)
	}
	initializeBoard(board, m, n)
	addUsedPieces(board, used)
	return Board{m, n, board, used}
}

func initializeBoard(runes [][]rune, m, n int) {
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			runes[i][j] = '_'
		}
	}
}

// usedPieces pieces to be set in the Board
func addUsedPieces(board [][]rune, used []primitives.Attacks) {
	for _, piece := range used {
		board[piece.Row()][piece.Col()] = piece.Name()
	}
}

func ShowBoard(board Board) {
	for i := 1; i < board.M; i++ {
		for j := 1; j < board.N; j++ {
			print(string(board.Board[i][j]))
		}
		println()
	}
	println("------------------------------")
}

func Place(board Board, piece primitives.Attacks) Board {
	return CreateBoardWithPieces(board.M, board.N, append(board.Used, piece))
}

// IsSafe return true if no other piece in the Board gets attacked by c and if c is not already placed
func IsSafe(b Board, c primitives.Attacks) bool {
	doesNotAttack := doesNotAttackOtherPieces(b.Used, c)
	return doesNotAttack && notYetPlaced(b, c)
}

func doesNotAttackOtherPieces(usedPieces []primitives.Attacks, c primitives.Attacks) bool {
	for _, p := range usedPieces {
		if p.Attacks(c) || c.Attacks(p) {
			return false
		}
	}
	return true
}

func notYetPlaced(b Board, c primitives.Attacks) bool {
	return b.Board[c.Row()][c.Col()] == '_'
}
