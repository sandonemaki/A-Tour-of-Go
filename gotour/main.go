package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// スライスの長さを 0 に変更しました。
	// これにより s は空になりますが、元の配列全体を参照しているため、容量は変わらず 6 のままです。
	s = s[:0]
	printSlice(s)

	// スライスの長さを 4 に再設定しています。
	// s の先頭から4つの要素（[2, 3, 5, 7]）が含まれます。
	// 容量は元の配列全体を参照しているので 6 のままです。
	s = s[:4]
	printSlice(s)

	// s = s[2:] を実行すると、スライスは s[2] から始まるように再設定されます。
	// 元の配列の 2 番目の要素（5）から開始し、元の配列の終わりまでを含むビューに変わります。
	// 元のスライスの 2 番目以降、つまり [5, 7] の2つの要素のみを指すため、長さが 2 になります。
	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	// 容量 (cap): スライスの開始位置から元の配列の終わりまでの要素数。
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
