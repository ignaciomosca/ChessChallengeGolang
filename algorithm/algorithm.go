package algorithm

import (
	b "chess/board"
	p "chess/pieces"
)

func Solution(board *b.Board, pieces []rune, solutions map[*b.Board]bool, testedConfigurations map[*b.Board]bool) map[*b.Board]bool {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					b := b.Place(board, c)
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
