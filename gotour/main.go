package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")

	fmt.Println("hello")
}

// hello
// World

// defer へ渡した関数の引数は、すぐに評価されますが、
// その関数自体は呼び出し元の関数がreturnするまで実行されません。
