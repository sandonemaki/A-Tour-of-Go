package main

import "golang.org/x/tour/pic"

// https://scrapbox.io/maki-Py/Go_Pic_%E9%96%A2%E6%95%B0

func Pic(dx, dy int) [][]uint8 {

}
func main() {
	pic.Show(Pic)
}

// uint8(intValue) は型変換を行っており、
// int 型の値を uint8 型に変換します。
// これにより、計算結果を8ビットの範囲に収めることができます。
