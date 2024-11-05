package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2} // Vertex構造体の初期化
	p := &v           // vのポインタを取得
	// 1e9 は Go では 1000000000 として扱われます
	p.X = 1e9      // ポインタを通じてXフィールドを変更
	fmt.Println(v) // vの内容を出力
}

// X フィールドが 1000000000 に、Y フィールドが 2 のままである
// {1000000000 2}
