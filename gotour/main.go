package main

import (
	"fmt"
	"math"
)

// Vertex 型の構造体を定義する
type Vertex struct {
	X, Y float64
}

// Absは先ほどと機能を変えずにメソッドとして定義されていた Abs を
// 通常の関数として定義している

// 一つ前：メソッドから関数への変更: 以前の Abs メソッドは Vertex 型のレシーバ (v Vertex) を持っていた。
// これにより、Vertex 型のインスタンス（例: v）に対して v.Abs() としてメソッドを呼び出せるようになっていた。

// 今回の例：Abs を通常の関数として定義し、Vertex 型のインスタンスを引数として受け取っている
func Abs(v Vertex) float64 {
	// メソッド内で v.X や v.Y を参照できる
	// v が指している Vertex インスタンスのフィールドを使用して計算。
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	// Vertex 型の変数 v を X=3、Y=4 で初期化
	v := Vertex{3, 4}
	// v.Abs() のようにして Abs メソッドを呼び出す
	fmt.Println(Abs(v))
}

// 5

// メソッドと関数の違い
// メソッドは特定の型に関連付けられているため、型.メソッド() という呼び出し方ができる。
// 型に関係する動作をまとめて表現することが可能です。
// 通常の関数は特定の型に関連付けられていないため、引数として型のインスタンスを渡して呼び出す。
// 書き方の違いは型に関連付けているかどうかだけ
