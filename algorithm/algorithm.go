package algorithm

import (
	b "ChessChallengeGolang/board"
	p "ChessChallengeGolang/pieces"
	"ChessChallengeGolang/primitives"
	"sort"
)

func Solution(board b.Board, pieces []rune, solutions *[]b.Board, testedConfigurations *[]b.Board) *[]b.Board {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					modifiedBoard := b.Place(board, c)
					if len(pieces) != 1 {
						if !contains(modifiedBoard, *testedConfigurations) {
							*testedConfigurations = append(*testedConfigurations, modifiedBoard)
							Solution(modifiedBoard, pieces[1:], solutions, testedConfigurations)
						}
					} else {
						if !contains(modifiedBoard, *solutions) {
							*solutions = append(*solutions, modifiedBoard)
							println("Found Solutions", len(*solutions))
						}
					}
				}
			}
		}
	}
	return solutions
}

func contains(boardContains b.Board, testedConfigurations []b.Board) bool {
	for _, k := range testedConfigurations {
		if len(k.Used) == len(boardContains.Used) && EqualBoards(k.Board, boardContains.Board, k.M, k.N) {
			return true
		}
	}
	return false
}

func SamePieces(board1, board2 []primitives.Attacks) bool {
	pieces1 := piecesToString(board1)
	pieces2 := piecesToString(board2)
	sort.Strings(pieces1)
	sort.Strings(pieces2)
	for i := 0; i < len(pieces1); i++ {
		if pieces1[i] != pieces2[i] {
			return false
		}
	}
	return true
}

func piecesToString(board []primitives.Attacks) []string {
	var results []string
	for _, p := range board {
		results = append(results, p.ToString())
	}
	return results
}

func EqualBoards(board1 [][]rune, board2 [][]rune, M, N int) bool {
	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}
