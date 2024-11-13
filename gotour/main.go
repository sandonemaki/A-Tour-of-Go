package main

import (
	"fmt"
)

// インターフェース I は M() メソッドを持つインターフェース
type I interface {
	M()
}

// 型 T: 構造体 T は string 型のフィールド S を持つ
type T struct {
	S string
}

// *T 型のポインタが M() メソッドを持つ
// これにより、*T 型がインターフェース I を実装していることになる
func (t *T) M() {
	fmt.Println(t.S)
}

// &T (ポインタ型) は M() メソッドを持つため、インターフェース I に代入できる
func main() {
	var i I = &T{"Hello"}
	i.M()
}
