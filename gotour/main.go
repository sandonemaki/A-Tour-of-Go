package main

import (
	"fmt"
	"math"
)

// MyFloat という新しい型を float64 を基に定義
type MyFloat float64

// MyFloat 型の変数 f に対して Abs というメソッドを定義
func (f MyFloat) Abs() float64 {
	if f < 0 {
		// f が負ならば 正の値を返す
		return float64(-f)
	}
	return float64(f)
}
func main() {
	// MyFloat 型の変数 f を -math.Sqrt2 で初期化
	f := MyFloat(-math.Sqrt2)
	// f の Abs メソッドを呼び出して、絶対値を表示
	fmt.Println(f.Abs())
}

// 1.4142135623730951
