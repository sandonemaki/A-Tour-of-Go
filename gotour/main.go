package main

import (
	"fmt"
)

// インターフェース I は M() メソッドを持つインターフェース
type I interface {
	M()
}

func main() {
	// インターフェース I は M() メソッドを持つ
	var i I

	describe(i) // 出力：(<nil>, <nil>)
	// M() メソッドが i を nil として受け取る。
	// nil pointer dereference というランタイムエラーが発生
	// panic: runtime error: invalid memory address or nil pointer dereference
	i.M() // nil dereference in dynamic method call
}

// describe 関数はインターフェース I 型の変数 i を引数にとり表示
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
