package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	// fmt.Printf の %v は、デフォルトのフォーマットで値を表示します。
	// スライスや配列を %v で表示した場合、要素はスペースで区切られた形で表示されます。
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// len=0 cap=0 []
// len=1 cap=1 [0]
// len=2 cap=2 [0 1]
// len=5 cap=6 [0 1 2 3 4]
