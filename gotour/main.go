package main

import "fmt"

// int 型のチャネルを受け付ける fibonacci 関数
func fibonacci(c, quit chan int) {
	// フィボナッチ数列の値を準備
	x, y := 0, 1
	for {
		// select 文内の2つの case は、どちらのチャネルが準備できたかによって選択される。
		select {
		// c チャネルに x の値を送信できる場合に実行
		// 受信者が c からの受信を待っていると、この case が実行され、x の値が送信される。
		// つまり、受信側が c からの値を受け取れる状態でなければ、この case はブロックされる。
		case c <- x:
			// x, y = y, x+y で次のフィボナッチ数を準備
			x, y = y, x+y
		// quit チャネルから値が受信できる場合に実行
		// quit チャネルに値が入っている（誰かが送信した）ときに実行。
		// quit チャネルが空である場合、この case はブロックされる。
		// 受信する値そのものは必要なく、quit チャネルから信号を受け取ったことを確認したら
		// fmt.Println("quit") を出力し、関数を終了させる
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// c はフィボナッチ数列を送信するためのチャネル
	c := make(chan int)
	// quit はフィボナッチ数列の生成を停止させるためのチャネル
	quit := make(chan int)

	go func() {
		// 新しい Goroutine で、c チャネルから値を10回受信し、その値を出力
		for i := 0; i < 10; i++ {
			// c チャネルから値を受信するごとに、fibonacci 関数内の select 文の case c <- x: が実行される
			// これにより、フィボナッチ数が c チャネルに送信される
			// 受信側（main のループ）が <-c でデータを取り出すことで、fibonacci 関数の送信が続行される
			fmt.Println(<-c)
		}
		// 10回のループが終わった後、quit <- 0 が実行され、quit チャネルに値を送信
		// これにより、fibonacci 関数内の select 文で case <-quit: が実行される
		// この case が実行されると、「quit」と表示され、fibonacci 関数は return して終了
		quit <- 0
	}()
	fibonacci(c, quit)
}

// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// quit
