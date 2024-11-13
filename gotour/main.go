package main

import (
	"fmt"
	"math"
)

// Vertex という構造体を定義
type Vertex struct {
	X, Y float64
}

// Vertex 型のポインタレシーバ v を持つ Scale メソッドを定義
// v の値を変更
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex 型の変数 v を引数に Scale 関数を定義
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// Vertex 型のポインタ v を初期化
	v := &Vertex{4, 3}
	// v.Abs() は (*v).Abs() として解釈される
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	// v.Abs() は (*v).Abs() として解釈される
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

// Before scaling: &{X:4 Y:3}, Abs: 5
// After scaling: &{X:20 Y:15}, Abs: 25

// %+v を使うと、構造体の各フィールド名が出力に含まれるため、
// フィールド名とその値がペアで表示される。
// 例えば、&{X:4 Y:3} のように X や Y のフィールド名が表示される。
