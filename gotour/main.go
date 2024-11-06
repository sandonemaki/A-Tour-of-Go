package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// board[0] の要素は []string{"X", "_", "X"}。
// strings.Join(board[1], " ") を実行すると、"X _ X" になります。
// fmt.Printf("%s\n", "X _ X") は、この文字列を出力し、改行します。
