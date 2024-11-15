package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// 実行のたびに異なる順序で結果が表示される可能性がある
// world
// hello
// hello
// world
// hello
// world
// world
// hello
// hello

// 実行の仕組み
// go say("world") は新しい Goroutine を起動しますが、
// メインの Goroutine はその開始を待たずに次の say("hello") に進みます。
//
// これにより、say("world") と say("hello") は並行して実行され、
// 2つの Goroutine が同時に動いている状態になります。
//
// 各 Goroutine 内では time.Sleep(100 * time.Millisecond) によって100ミリ秒の遅延が発生します。
// これにより、fmt.Println が交互に実行されるように見えますが、
// 実際の実行順序は Go のスケジューラによって管理されるため、実行のたびに結果が異なることがあります。

// プログラム内で go キーワードを使うと、
// メインの Goroutine が終了した時点で他の Goroutine も強制的に終了します。
// したがって、もし main() 内で say("hello") が終了してしまうと、その直後にプログラム全体が終了し、
// say("world") が完全に実行されないことがあります。

// 実際に Goroutine の動作を管理したい場合は、sync.WaitGroup などを使って
// メインの Goroutine が他の Goroutine の終了を待機するようにすることが一般的です。
