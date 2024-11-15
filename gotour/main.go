package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image 型の定義
type Image struct {
	width, height int
}

// Bounds メソッドの実装
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel メソッドの実装
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At メソッドの実装
func (img Image) At(x, y int) color.Color {
	// ピクセルの色の値を計算
	v := uint8((x + y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// Image インスタンスを生成し、画像を表示
	m := Image{width: 256, height: 256}
	pic.ShowImage(m)
}
