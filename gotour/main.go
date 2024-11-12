package main

import (
	"fmt"
	"math"
)

// Vertex という構造体を定義
type Vertex struct {
	X, Y float64
}

// Vertex 型のポインタレシーバ v を持つ Abs メソッドを定義
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Vertex 型の変数 v を引数に Scale 関数を定義
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	// Vertex 型の変数 v を初期化
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	// Vertex 型のポインタ p を初期化
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	// ポインタ変数 p を渡す
	fmt.Println(AbsFunc(*p))
}

// 5
// 5
// 5
// 5
