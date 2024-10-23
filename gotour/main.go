package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0

// defer へ渡した関数が複数ある場合、
// その呼び出しはスタック( stack )されます。

// 呼び出し元の関数がreturnするとき、
// defer へ渡した関数は LIFO(last-in-first-out) の順番で実行されます。
