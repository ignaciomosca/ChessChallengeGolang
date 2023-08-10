package algorithm

import (
	b "chesschallengegolang/board"
	p "chesschallengegolang/pieces"
	"chesschallengegolang/primitives"
	"sort"
	"strings"
)

func Solution(M, N int, pieceCounts map[rune]int) []b.Board {
	solutions := []b.Board{}
	visitedConfigs := make(map[string]bool)
	initialBoard := b.CreateBoard(M, N)
	solve(initialBoard, pieceCounts, &solutions, &visitedConfigs)
	return solutions
}

func popPieceType(pieceCounts map[rune]int) (rune, bool) {
	for k, v := range pieceCounts {
		if v > 0 {
			return k, true
		}
	}
	return 0, false
}

func solve(board b.Board, remainingPieces map[rune]int, solutions *[]b.Board, visitedConfigs *map[string]bool) {
	configString := configurationToString(board)
	if (*visitedConfigs)[configString] {
		return // Skip already visited configurations
	}

	nextPieceType, hasRemainingPieces := popPieceType(remainingPieces)
	if !hasRemainingPieces {
		solutionBoard := copyBoard(board)
		configStringSolutionBoard := configurationToString(solutionBoard)
		(*visitedConfigs)[configStringSolutionBoard] = true
		*solutions = append(*solutions, solutionBoard)
		return
	}

	remainingCounts := make(map[rune]int)
	for k, count := range remainingPieces {
		remainingCounts[k] = count
	}

	for i := 1; i < board.M; i++ {
		for j := 1; j < board.N; j++ {
			piece := p.CreatePiece(nextPieceType, i, j)
			if board.Board[i][j] == '_' && b.IsSafe(board, piece) {
				modifiedBoard := b.Place(board, piece)
				remainingCounts[nextPieceType]--
				solve(modifiedBoard, remainingCounts, solutions, visitedConfigs)
				remainingCounts[nextPieceType]++
			}
		}
	}
}

func copyBoard(original b.Board) b.Board {
	copiedBoard := b.CreateBoard(original.M, original.N)
	for i := 1; i < original.M; i++ {
		for j := 1; j < original.N; j++ {
			copiedBoard.Board[i][j] = original.Board[i][j]
		}
	}
	copiedBoard.Used = append([]primitives.Attacks{}, original.Used...)
	return copiedBoard
}

func configurationToString(board b.Board) string {
	var pieces []string
	for _, piece := range board.Used {
		pieces = append(pieces, piece.ToString())
	}
	sort.Strings(pieces)
	return strings.Join(pieces, "|")
}
