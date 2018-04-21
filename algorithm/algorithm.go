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
							Solution(modifiedBoard, removeFirstElement(pieces), solutions, add(testedConfigurations, modifiedBoard))
						}
					} else {
						if !contains(&modifiedBoard, solutions) {
							add(solutions, modifiedBoard)
						}
					}
				}
			}
		}
	}
	return solutions
}
func add(boardList *[]b.Board, newBoard b.Board) *[]b.Board {
	newCollection := append(*boardList, newBoard)
	return &newCollection
}


func contains(board *b.Board, testedConfigurations *[]b.Board) bool {
	for k := range *testedConfigurations {
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
