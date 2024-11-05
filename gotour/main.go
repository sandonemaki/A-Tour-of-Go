package main

import (
	"fmt"
)

func main() {
	// aという名前の文字列型配列を宣言。
	//[2]stringは「2要素を持つ文字列型配列」を意味
	// この配列は、string型の要素を2つ保持することができる。

	// varキーワードによって、aはゼロ値で初期化
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	// 出力はGoの配列のデフォルトの表示形式で表示
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
