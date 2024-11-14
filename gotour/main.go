package main

import (
	"fmt"
)

// Person 型の構造体を定義
type Person struct {
	Name string
	Age  int
}

// Person 型のレシーバを持つ String メソッドは、
// fmt.Stringer インターフェースを実装しています。
// このメソッドにより、fmt.Println や fmt.Sprintf などで
// Person 型のインスタンスを出力する際に、自動的にこの String メソッドが呼び出され、
// カスタムフォーマットで出力されます。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// Person 型のインスタンス a と z を作成
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// type Stringer interface {
//   String() string
// }
// Stringer インタフェースは、stringとして表現することができる型です。
// fmt パッケージ(と、多くのパッケージ)では、変数を文字列で出力するために
// このインタフェースがあることを確認します。
