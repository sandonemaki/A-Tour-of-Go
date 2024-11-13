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

// *T 型のポインタ (&T) をインターフェース I に代入
// i.M() を呼び出すと、M() メソッドが実行され、t.S の内容 ("Hello") が表示される。
// S は T 構造体のフィールドの名前。つまり、T 型のインスタンスが持つデータの入れ物。
// "Hello" は S フィールドに格納される具体的な値
func main() {
	var i I = &T{"Hello"}
	i.M()
}

// Hello
