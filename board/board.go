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
	initializeBoard(board, m, n)
	return Board{m, n, board, make(map[primitives.Attacks]bool)}
}

// CreateBoardWithPieces creates a board with the pieces being @used
func CreateBoardWithPieces(m, n int, used map[primitives.Attacks]bool) Board {
	board := make([][]rune, m)
	for i := range board {
		board[i] = make([]rune, n)
	}
	initializeBoard(board, m, n)
	addUsedPieces(board, used)
	return Board{m, n, board, used}
}

func initializeBoard(runes [][]rune, m,n int) {
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			runes[i][j]='_'
		}
	}
}

// usedPieces pieces to be set in the board
func addUsedPieces(board [][]rune, used map[primitives.Attacks]bool) {
	for piece := range used {
		board[piece.Row()][piece.Col()] = piece.Name()
	}
}

func ShowBoard(board Board) {
	for i := 1; i < board.M; i++ {
		for j := 1; j < board.N; j++ {
			print(string(board.board[i][j]))
		}
		println()
	}
	println("------------------------------")
}

func Place(board Board, piece primitives.Attacks) Board {
	board.used[piece] = true
	return CreateBoardWithPieces(board.M, board.N, board.used)
}

// IsSafe return true if no other piece in the board gets attacked by c and if c is not already placed
func IsSafe(b Board, c primitives.Attacks) bool {
	doesNotAttackOtherPieces := doesNotAttackOtherPieces(b.used, c)
	return doesNotAttackOtherPieces && notYetPlaced(b, c)
}

func doesNotAttackOtherPieces(usedPieces map[primitives.Attacks]bool, c primitives.Attacks) bool {
	for p := range usedPieces {
		if p.Attacks(c) || c.Attacks(p) {
			return false
		}
	}
	return true
}

func notYetPlaced(b Board, c primitives.Attacks) bool {
	return b.board[c.Row()][c.Col()] == '_'
}
