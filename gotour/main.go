package main

import "fmt"

// Vertex という構造体を定義
type Vertex struct {
	X, Y float64
}

// Vertex 型のポインタレシーバ v を持つ Scale 関数は引数で受け取った変数 f に対して実行する。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex 型のポインタ変数 v 、f を引数に Scale 関数を定義
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	// Vertex 型の変数 v を初期化
	v := Vertex{3, 4}
	// ポインタレシーバ v を持つ Scale メソッドを呼び出し
	// Scale メソッドはポインタレシーバ (v *Vertex) を持っているため、
	// v.Scale(2) の呼び出し時に v が自動的にポインタとして渡され v 自体が変更
	v.Scale(2)
	// ScaleFunc はポインタを引数として受け取る関数なので、
	// 呼び出し側では & を付けて v のアドレスを渡す。
	// これにより、関数内で v の内容が変更され、main 内の v に反映される
	ScaleFunc(&v, 10)

	// Vertex 型のポインタ p を初期化
	p := &Vertex{4, 3}
	// ポインタレシーバ p に対して Scale メソッドを呼び出し
	// p はもともと &Vertex{4, 3} というポインタ型なので、
	// Scale メソッドの呼び出しでは p がそのまま渡される
	p.Scale(3)
	// ScaleFunc もポインタを引数として受け取る関数なので、p（ポインタ型）はそのまま渡される。
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// {60 80} &{96 72}

// 呼び出し側で & をつける理由
// ・関数やメソッドがポインタ型を引数として受け取る場合、
// 呼び出し側で & を使って変数のアドレスを渡します。

// まとめ
// メソッドや関数がポインタ型のレシーバや引数を持つ場合、
// 呼び出し側で & をつけてアドレスを渡します。
