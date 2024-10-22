package main

import (
	"fmt"
	"math"
)

// 関数の引数が float64 型なので、これらの整数を float64 型に変換して処理。
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
