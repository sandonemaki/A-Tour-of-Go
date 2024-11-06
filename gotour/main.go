package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// ビットシフト演算 1 << uint(i) は、数学的には 2**i（2 の i 乗）と同じ
		// pow = [1, 2, 4, 8, 16, 32, 64, 128, 256, 512]
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
