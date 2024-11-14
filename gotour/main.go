package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// ErrNegativeSqrt 型の Error メソッド
func (e ErrNegativeSqrt) Error() string {
	// fmt.Sprint(e) を直接呼び出すと e の Error() メソッドが再帰的に呼び出されるため無限ループになる。
	// これを避けるために float64(e) に型変換し単なる数値として出力する。
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt 関数の修正
func Sqrt(x float64) (float64, error) {
	// Sqrt 関数に正の数と負の数を渡し、エラーメッセージを確認します。
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// 簡易的なニュートン法による近似
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	// 正の数のテスト
	fmt.Println(Sqrt(2))
	// 負の数のテスト
	fmt.Println(Sqrt(-2))
}
