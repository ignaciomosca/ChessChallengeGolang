package algorithm

import (
	b "chesschallengegolang/board"
	p "chesschallengegolang/pieces"
	"sort"
	"strings"
)

func Solution(board b.Board, pieces []rune, solutions *map[string]bool, testedConfigurations *map[string]bool) *[]string {
	if len(pieces) != 0 {
		for i := 1; i < board.M; i++ {
			for j := 1; j < board.N; j++ {
				c := p.CreatePiece(pieces[0], i, j)
				if b.IsSafe(board, c) {
					modifiedBoard := b.Place(board, c)
					configString := ConfigurationToString(modifiedBoard)
					if !modifiedBoard.ConfigLookup[configString] {
						modifiedBoard.ConfigLookup[configString] = true
						if len(pieces) != 1 {
							if !(*testedConfigurations)[configString] {
								(*testedConfigurations)[configString] = true
								Solution(modifiedBoard, pieces[1:], solutions, testedConfigurations)
							}
						} else {
							if !(*solutions)[configString] {
								(*solutions)[configString] = true
								println("Found Solutions", len(*solutions))
							}
						}
					}
				}
			}
		}
	}
	var keys []string
	for key := range *solutions {
		keys = append(keys, key)
	}
	return &keys
}

func ConfigurationToString(board b.Board) string {
	var pieces []string
	for _, piece := range board.Used {
		pieces = append(pieces, piece.ToString())
	}
	sort.Strings(pieces)
	return strings.Join(pieces, "|")
}
