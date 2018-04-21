package algorithm

import (
	b "chess/board"
	p "chess/pieces"
	"reflect"
)

func Solution(board b.Board, pieces []rune, solutions map[*b.Board]bool, testedConfigurations map[*b.Board]bool) map[*b.Board]bool {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					modifiedBoard := b.Place(board, c)
					if len(pieces) != 1 {
						if !contains(&modifiedBoard, testedConfigurations) {
							testedConfigurations[&modifiedBoard] = true
							Solution(modifiedBoard, removeFirstElement(pieces), solutions, testedConfigurations)
						}
					} else {
						if !contains(&modifiedBoard, solutions) {
							solutions[&modifiedBoard] = true
						}
					}
				}
			}
		}
	}
	return solutions
}

func contains(board *b.Board, testedConfigurations map[*b.Board]bool) bool {
	for k := range testedConfigurations {
		if reflect.DeepEqual(k, board) {
			return true
		}
	}
	return false
}

func removeFirstElement(pieces []rune) []rune {
	r := pieces
	remove := r[1:]
	return remove
}
