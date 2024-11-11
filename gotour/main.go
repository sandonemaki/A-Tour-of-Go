package main

import (
	"fmt"
	"math"
)

// Vertex 型の構造体を定義する
type Vertex struct {
	X, Y float64
}

// Go にはクラスの概念はないが、構造体や型に対してメソッドを定義することができる。
// このメソッドは特別なレシーバ（receiver）引数を持っている
// レシーバは、func キーワードとメソッド名の間に記述され、メソッドがどの型に属するかを示す。

// Abs というメソッドを Vertex 型に定義。
// (v Vertex) がレシーバで、これにより Abs メソッドは Vertex 型に関連付けられる。
func (v Vertex) Abs() float64 {
	// メソッド内で v.X や v.Y を参照できる
	// v が指している Vertex インスタンスのフィールドを使用して計算。
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	// Vertex 型の変数 v を X=3、Y=4 で初期化
	v := Vertex{3, 4}
	// v.Abs() のようにして Abs メソッドを呼び出す
	fmt.Println(v.Abs())
}

// 5
