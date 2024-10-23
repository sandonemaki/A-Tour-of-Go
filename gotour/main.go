package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// 初期値の設定
	z := 1.0
	// 10回までループ
	for i := 0; i < 10; i++ {
		// 更新式
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %f\n", i+1, z)
	}
	return z
}

func main() {
	// Sqrt関数を呼び出して結果を表示
	fmt.Println("Approximation of sqrt(2):", Sqrt(2))
	// 標準ライブラリの math.Sqrt での確認
	fmt.Println("math.Sqrt(2):", math.Sqrt(2))
}

// Iteration 1: z = 1.500000
// Iteration 2: z = 1.416667
// Iteration 3: z = 1.414216
// Iteration 4: z = 1.414214
// Iteration 5: z = 1.414214
// Iteration 6: z = 1.414214
// Iteration 7: z = 1.414214
// Iteration 8: z = 1.414214
// Iteration 9: z = 1.414214
// Iteration 10: z = 1.414214
// Approximation of sqrt(2): 1.4142135623746899
// math.Sqrt(2): 1.4142135623730951
