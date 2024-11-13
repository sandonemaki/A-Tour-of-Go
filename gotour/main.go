package main

import (
	"fmt"
)

func main() {
	// interface i は 0 個のメソッドを持つインターフェース
	// interface{} は空のインターフェースで、Go では 任意の型 を保持できる特別なインターフェース。
	var i interface{}
	describe(i)

	// 空のインターフェースは、任意の型の値を保持できる。
	// i は int 型の値 42 を保持
	i = 42
	describe(i)

	// i は string 型の値 "hello" を保持
	i = "hello"
	describe(i)
}

// describe 関数の引数に空のインターフェイス i が渡され値を表示
// 空のインターフェースは、未知の型の値を扱うコードで使用される。
// interface{} 型の任意の数の引数を受け取る。
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// (<nil>, <nil>)
// (42, int)
// (hello, string)
