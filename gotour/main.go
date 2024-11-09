package main

import (
	"fmt"
)

// adder 関数が「func(int) int 型の関数」を返す
func adder() func(int) int {
	sum := 0
	// adder は func(x int) int { sum += x; return sum } という関数を返す
	return func(x int) int {
		// x を受け取って sum に加え、更新された sum の値を返す
		sum += x
		return sum
	}
}

func main() {
	// pos や neg という変数には、adder によって返されたクロージャ（無名関数）が格納。
	// この時点で adder は終了。
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			// pos や neg に格納されたクロージャが実行されるが、adder 関数は再度実行されない。
			//これらのクロージャは、sum の状態を保持し続ける性質を持っている
			pos(i),    // pos に格納されたクロージャが i を受け取り pos の sum を更新し、新しい値を返す。
			neg(-2*i), // neg に格納されたクロージャが -2*i を受け取り neg の sum を更新して返す。
		)
	}
}
