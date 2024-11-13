package main

import (
	"fmt"
)

// 型switchの宣言は、型アサーション i.(T) と同じ構文を持ちますが、
// 特定の型 T はキーワード type に置き換えられます。
func do(i interface{}) {
	// インターフェースの値 i によって保持される値を保持
	// i が保持している具体的な型をチェック
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// do(true) の場合、i は bool 型なので、どの case にも一致せず、
		// default ブロックが実行されます。
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

// Twice 21 is 42
// "hello" is 5 bytes long
// I don't know about type bool!
