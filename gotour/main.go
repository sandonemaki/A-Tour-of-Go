package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // 型は Vertex
	v2 = Vertex{X: 1} // Yは暗黙的に0になる
	v3 = Vertex{}     // XとYがともに0になる
	// & 演算子を使って、Vertex{1, 2} のアドレスを取得し、そのポインタを p に代入しています。
	// つまり、p は *Vertex 型であり、Vertex のメモリアドレスを指しています。
	// Vertex{1, 2} のアドレスが p に代入されるため、p の値は &{1 2} と表示されます。
	p = &Vertex{1, 2} // 型は *Vertex（Vertexのポインタ型）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
