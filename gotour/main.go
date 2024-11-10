package main

import (
	"fmt"
)

// fibonacci 関数が「func(int) int 型の関数」を返す
func fibonacci() func(int) int {
	// 初期化
	a, b := 0, 1
	return func(n int) int {
		if n == 0 {
			return a
		} else if n == 1 {
			return b
		} else {
			// フィボナッチ数列の計算
			a, b = b, a+b
			return b
		}
	}
}

func main() {
	// fibonacci() 関数からクロージャを取得し、変数 f に保存する。
	// これにより、f は状態を保持したまま呼び出すたびに次のフィボナッチ数を計算する関数となる。
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
