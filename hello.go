package main

import (
	"fmt"
	"chess/pieces"
)

func main() {
	fmt.Println(pieces.CreateKnight(1, 1))
	fmt.Println(pieces.CreateBishop(1, 1))
	fmt.Println(pieces.CreateKing(1, 1))
	fmt.Println(pieces.CreateQueen(1, 1))
	fmt.Println(pieces.CreateRook(1, 1))
}
