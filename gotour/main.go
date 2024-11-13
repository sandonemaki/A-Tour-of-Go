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
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	// インターフェース I は M() メソッドを持つ
	var i I

	// 型 T: 構造体 T は string 型のフィールド S を持つ
	var t *T

	// t は nil。インターフェース i は
	// *T 型として nil を保持している。

	// この段階で i は nil の具体的な値（*T 型）を持っていますが、
	// i 自体は nil ではありません
	i = t
	// インターフェースが nil の具体的な値を持っていることを示す。
	describe(i) // 出力: (<nil>, *main.T)
	// M() メソッドが t を nil として受け取り、fmt.Println("<nil>") が出力。
	i.M()

	i = &T{"hello"}
	describe(i) // 出力: (&{hello}, *main.T)
	// M() メソッドが t を受け取り、"hello" を出力する。
	i.M()
}

// describe 関数はインターフェース I 型の変数 i を引数にとり表示
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// (<nil>, *main.T)
// <nil>
// (&{hello}, *main.T)
// hello

// 解説
// Go では、インターフェース I に具体的な値（アンダーライイング・バリュー）が
// nil の場合でも、インターフェース自体は nil ではない状態として扱われます。
// つまり、インターフェースは「nil の具体的な値」を保持していますが、
// インターフェース型 I そのものは非 nil です。

// i.M() が呼び出されるとき、i が保持している具体的な値が nil であれば、
// M() メソッドのレシーバは nil として渡されます。
// この場合、M() メソッド内で適切に nil チェックをして処理を行うのが一般的です。

// Go では、nil レシーバを適切に扱えるようなメソッドを書くことで、
// null ポインター例外のようなエラーを回避できます。
