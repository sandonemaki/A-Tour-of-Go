package main

import (
	"fmt"
)

func sum(nums []int, ch chan int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	ch <- total // 合計をチャネルに送信
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// チャネルを生成
	ch := make(chan int)

	// 配列を半分に分けて2つのGoroutineで合計を計算
	// 2つの Goroutine が ch に結果を送信
	// 最初の半分のスライス（1, 2, 3, 4, 5）を合計して ch に送信。
	go sum(nums[:len(nums)/2], ch)
	// 後半のスライス（6, 7, 8, 9, 10）を合計して ch に送信。
	go sum(nums[len(nums)/2:], ch)

	// チャネルから合計を受信
	// これにより、2つの Goroutine が同じ ch にそれぞれの合計を送信。
	// x := <-ch と y := <-ch で ch から2回受信することで、
	// 2つの Goroutine の結果をそれぞれ x と y に割り当てている。
	x := <-ch
	y := <-ch

	fmt.Println("Total Sum:", x+y) // 合計を表示
}

// Total Sum: 55
