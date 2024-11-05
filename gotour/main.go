package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// a と b の両方が同じ配列 names の異なる部分を指していることです。
	// さらに、a と b は names の同じメモリ領域（インデックス 1 の要素 "Paul"）を指しています。
	// このため、スライス b でインデックス 0 の要素を変更すると、スライス a にも影響が出ます。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// [John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
