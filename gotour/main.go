package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// i が float64 型でないため、ok は false になり、
	// f にはその型のゼロ値（0）が代入される。
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 型アサーションをチェックなしで実行しているため、
	// i が float64 型でない場合は panic（ランタイムエラー）が発生
	f = i.(float64) // panic
	fmt.Println(f)
}

// hello
// hello true
// 0 false
// panic: interface conversion: interface {} is string, not float64

// まとめ
// 型アサーションを使用する際には、チェック付きの形式 (s, ok := i.(type)) を使うことで、
// 安全に型変換の成否を確認できる。
// チェックなしの型アサーション (s := i.(type)) は、
// 指定された型でない場合に panic を引き起こすため、注意が必要。
