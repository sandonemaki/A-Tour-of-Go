package main

import (
	"fmt"
	"math"
)

// Vertex という構造体を定義
type Vertex struct {
	X, Y float64
}

// Abs 関数は Vertex型の変数 v に対して実行される
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 関数はポインタVertex型の変数 v と変数 f に対して実行される。
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// Vertex 型の変数 v を初期化
	v := Vertex{3, 4}
	// ポインタ v　と 10 を引数に Scale 関数を呼び出し
	Scale(&v, 10)
	// 変数 v を引数に Abs 関数を呼び出し
	fmt.Println(Abs(v))
}

// 50

// もし Scale 関数の引数を *Vertex から Vertex に変更すると、
// 関数内で v のコピーが渡されるため、Scale 関数内の変更は元の v に反映されません。

// ポインタ引数 (*Vertex):
// ・Scale 内の操作が元の v に反映される。
// 値引数 (Vertex):
// ・Scale 内の操作はコピーに対して行われ、元の v には影響しない。
