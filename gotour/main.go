package main

import (
	"fmt"
)

// Vertex は緯度 (Lat) と経度 (Long) の2つの float64 型のフィールドを持つ構造体
type Vertex struct {
	Lat, Long float64
}

// map[string]Vertex 型のマップ m を作成
// キーとして string 型、値として Vertex 型を持つマップ
// mapに渡すトップレベルの型が単純な型名である場合は、
// リテラルの要素から推定できますので、その型名を省略することができます。
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
