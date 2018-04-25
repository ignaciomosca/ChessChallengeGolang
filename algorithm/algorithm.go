package algorithm

import (
	b "chess/board"
	p "chess/pieces"
	"reflect"
)

func Solution(board b.Board, pieces []rune, solutions *[]b.Board, testedConfigurations *[]b.Board) *[]b.Board {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					modifiedBoard := b.Place(board, c)
					if len(pieces) != 1 {
						if !contains(&modifiedBoard, testedConfigurations) {
							*testedConfigurations = append(*testedConfigurations, modifiedBoard)
							Solution(modifiedBoard, pieces[1:], solutions, testedConfigurations)
						}
					} else {
						if !contains(&modifiedBoard, solutions) {
							*solutions = append(*solutions, modifiedBoard)
						}
					}
				}
			}
		}
	}
	return solutions
}

func contains(boardContains *b.Board, testedConfigurations *[]b.Board) bool {
	for k := range *testedConfigurations {
		if reflect.DeepEqual(k, boardContains) {
			return true
		}
	}
	return false
}

