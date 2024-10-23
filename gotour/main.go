package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	// osは評価される値
	//runtime.GOOS は、Goの組み込みパッケージである runtime パッケージの変数が格納。
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
