package main

import (
	"fmt"
)

// Vertex は緯度 (Lat) と経度 (Long) の2つの float64 型のフィールドを持つ構造体
type Vertex struct {
	Lat, Long float64
}

func main() {
	// map[string]Vertex 型のマップ m を作成
	// キーとして string 型、値として Vertex 型を持つマップ
	var m map[string]Vertex
	m = make(map[string]Vertex)
	// "Bell Labs" というキーに Vertex 型の値 40.68433, -74.39967 を追加
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
