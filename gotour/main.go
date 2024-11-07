package main

import (
	"fmt"
	"math"
)

// compute 関数は、引数として２つの float64 型の値をとり、float64 型の結果を返す関数を受け取ります。
func compute(fn func(float64, float64) float64) float64 {
	// 引数として渡された fn を 3 と 4 という引数で実行し結果を返します。
	return fn(3, 4)
}

func main() {
	// hypot は匿名関数で2つの float64 型の引数 x と y を受け取り、その平方根を返します。
	// これは三角形の斜辺（hypotenuse）を求めるための計算
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// hypot 関数を実行
	fmt.Println(hypot(5, 12))

	// compute 関数に hypot を渡し、3 と 4 を引数として計算した結果を出力する
	fmt.Println(compute(hypot))
	// compute 関数に math.Pow を渡し、3^4 を計算して出力
	fmt.Println(compute(math.Pow))
}

// 13
// 5
// 81
