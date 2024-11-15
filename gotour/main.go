package main

import "fmt"

func main() {
	ch := make(chan int, 3) // バッファサイズ3のチャネル

	// バッファにデータを送信
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("3つのデータをバッファに送信しました")

	// バッファがいっぱいなので、次の送信でブロックされる
	go func() {
		ch <- 4
		fmt.Println("4をチャネルに送信しました")
	}()

	fmt.Println("受信を開始します")

	// チャネルからデータを受信
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// これで4つ目のデータが送信され、ブロックが解除される
	fmt.Println(<-ch)
}
