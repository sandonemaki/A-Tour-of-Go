package main

import (
	"fmt"
	"math"
)

// Abser は、Abs() メソッドを持つインターフェース
// Abs() メソッドは、float64 型の値を返す
// これを実装する任意の型は、Abs() メソッドを定義しなければならない。
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// MyFloat 型の値を a に代入。
	// f は Abser インターフェースを実装している。
	a = f

	// Vertex 型のポインタを a に代入。
	// ポインタがインターフェースを実装している。
	a = &v

	// 次の行では、v はポインタではありません。
	// エラー。Vertex 型の値そのものは Abser を実装していません。

	// Abs() メソッドは *Vertex (ポインタレシーバ) にのみ定義されているため、
	// ポインタである &v を Abser に代入する必要があります。
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloat は Abs() メソッドを持ち、Abser インターフェースを実装している
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Vertex 型も Abs() メソッドを持ち、
// *Vertex (ポインタレシーバ) が Abser インターフェースを実装している。
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 5

// インターフェース型に代入するには、その型がインターフェースを実装している必要がある。
// ポインタレシーバ (*Vertex) を持つメソッドは、その型のポインタにのみ関連づけられる。
