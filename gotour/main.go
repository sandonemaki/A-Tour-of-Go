package main

import (
	"fmt"
	"image"
)

func main() {
	// 画像の色モデル（例えばRGBなど）を返すメソッド
	// 100x100 ピクセルの新しい画像を作成します。
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// 画像のサイズ（範囲）を表す矩形を返すメソッド
	// Bounds：戻り値は Rectangle 型で、画像の左上（開始座標）と右下（終了座標）の位置を持っている
	// 結果: (0,0)-(100,100)
	fmt.Println(m.Bounds())
	// 指定した座標のピクセルの色を返すメソッド
	// At(x, y int): 指定した座標 (x, y) の位置にあるピクセルの色を返す。
	// 返り値は color.Color 型
	// (0,0) の位置にあるピクセルの RGBA 値を出力します。
	// 結果: 0 0 0 0　初期状態ではそのピクセルが透明な黒（R, G, Bが0でアルファも0）
	fmt.Println(m.At(0, 0).RGBA())
}

// (0,0)-(100,100)
// 0 0 0 0
