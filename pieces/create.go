package pieces

import "chess/primitives"

func CreatePiece(piece primitives.Attacks, row, col int) primitives.Attacks{
	switch piece.Name() {
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
