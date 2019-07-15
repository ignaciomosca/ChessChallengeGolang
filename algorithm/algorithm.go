package algorithm

import (
	b "ChessChallengeGolang/board"
	p "ChessChallengeGolang/pieces"
	"ChessChallengeGolang/primitives"
	"sort"
)

// Solution returns a list of solutions of type Board
func Solution(board b.Board, pieces []rune, solutions *[]b.Board, testedConfigurations map[int][]primitives.Attacks) *[]b.Board {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					board.Board[i][j] = pieces[0]
					board.Used = append(board.Used, c)
					if len(pieces) != 1 {
						if !contains(board.Used, testedConfigurations) {
							testedConfigurations[len(board.Used)] = append(testedConfigurations[len(board.Used)], board.Used...)
							Solution(board, pieces[1:], solutions, testedConfigurations)
							board.Board[i][j] = '_'
							board.Used = board.Used[:len(board.Used)-1]
						}
					} else {
						*solutions = append(*solutions, b.CreateBoardWithPieces(board.M, board.N, board.Used))
					}
				}
			}
		}
	}
	return solutions
}

func contains(usedContains []primitives.Attacks, testedConfiguration map[int][]primitives.Attacks) bool {
	for _, k := range testedConfiguration {
		if samePieces(usedContains, k) {
			return true
		}
		return false
	}
	return false
}

func samePieces(used1, used2 []primitives.Attacks) bool {
	u1 := append([]primitives.Attacks{}, used1...)
	u2 := append([]primitives.Attacks{}, used2...)
	for _, u := range [][]primitives.Attacks{u1, u2} {
		sort.Slice(u, func(i, j int) bool {
			x, y := u[i], u[j]
			if x.Row() != y.Row() {
				return x.Row() < y.Row()
			}
			if x.Col() != y.Col() {
				return x.Col() != y.Col()
			}
			return x.Name() < y.Name()
		})
	}
	if len(u1) != len(u2) {
		return false
	} else {
		for i := range u1 {
			if u1[i].Name() != u2[i].Name() || u1[i].Row() != u2[i].Row() || u1[i].Col() != u2[i].Col() {
				return false
			}
		}
		return true
	}
}
