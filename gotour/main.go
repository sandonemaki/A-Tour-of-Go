package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// MyError ポイントレシーバを持つ Error メソッドが string 型を返す
// このメソッドは error インターフェースを実装しています。
// error インターフェースは Error() string というメソッドシグネチャを持ち、
// エラーメッセージを文字列として返すことが求められます。
func (e *MyError) Error() string {
	// MyError 型の When フィールドを文字列に変換して返す
	// e.When は MyError 型の When フィールドで、型は time.Time
	// e.When.String() は time.Time 型の String() メソッドを呼び出し、When を文字列に変換
	// "2024-11-15 15:04:05.123456" のような形式で時間が表現
	return e.When.String() + ": " + e.What
}

// run 関数は error 型を返す
// 戻り値として MyError 型のインスタンスのポインタを返す。
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// 2024-11-15 01:40:24.886132 +0900 JST m=+0.000101126: it didn't work

// run() 関数を呼び出すと、MyError のインスタンスが作成され、そのポインタが error 型として返される。
// main() 関数でこの run() の戻り値を err 変数に代入し、エラーチェックを行う。
// エラーが存在する場合、fmt.Println(err) などで出力すると、
// MyError 型の Error() メソッドが自動的に呼ばれ、カスタムメッセージが表示される。
