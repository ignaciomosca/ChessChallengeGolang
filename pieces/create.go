package pieces

import "chesschallengegolang/primitives"

func CreatePiece(piece rune, row, col int) primitives.Attacks {
	switch piece {
	case 'Q':
		return CreateQueen(row, col)
	case 'N':
		return CreateKnight(row, col)
	case 'K':
		return CreateKing(row, col)
	case 'R':
		return CreateRook(row, col)
	case 'B':
		return CreateBishop(row, col)
	default:
		panic("Unknown Piece")
	}
}
