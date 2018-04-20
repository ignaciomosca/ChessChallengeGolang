package algorithm

import (
	"chess/board"
	p "chess/pieces"
)

func Solution(board *board.Board, pieces []rune, solutions map[*board.Board]bool, testedConfigurations map[*board.Board]bool) map[*board.Board]bool {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if board.IsSafe(c) {
					b := board.Place(c)
					if len(pieces) != 1 {
						if !testedConfigurations[b] {
							testedConfigurations[b] = true
							Solution(b, pieces[1:], solutions, testedConfigurations)
						}
					} else {
						if !solutions[b] {
							solutions[b] = true
						}
					}
				}
			}
		}
	}
	return solutions
}
