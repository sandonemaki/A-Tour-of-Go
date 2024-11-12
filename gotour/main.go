package main

import (
	"fmt"
	"math"
)

// Vertex という構造体を定義
type Vertex struct {
	X, Y float64
}

// Vertex 型の変数 v に対して Abs というメソッドを定義
// Abs メソッドは変数レシーバ
// これはレシーバのコピーを渡します。メソッド内での操作は元の変数に影響を与えません。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Vertex 型のポインタ v に対して Scale というメソッドを定義
// Scale メソッドはポインタレシーバ
// 元の変数へのポインタを渡します。メソッド内での操作は元の変数自体を変更できます。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// Vertex 型の変数 v を初期化
	v := Vertex{3, 4}
	// v の Scale メソッドを呼び出し
	// ポインタレシーバであるため、v 自体が変更される
	v.Scale(10)
	// v の Abs メソッドを呼び出し
	// 元の v の X と Y が変更された後の値を出力
	fmt.Println(v.Abs())
}

// Abs メソッドは 変数レシーバ なので、
// 元の変数 v のコピーが渡されてメソッド内の変更は元の v に影響を与えない。
// 一方、Scale メソッドは ポインタレシーバ なので、
// 元の変数 v へのポインタが渡され、メソッド内での操作が元の変数 v 自体に反映される。

// 要点
// ・変数レシーバ: コピーが渡され、元の値は変更されない。
// ・ポインタレシーバ: ポインタが渡され、元の値が変更される。

// 50
