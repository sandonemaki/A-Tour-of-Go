package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read メソッドを実装し、b スライスに 'A' を埋め込みます。
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // バイトスライスの各要素に 'A' を設定
	}
	return len(b), nil // 埋め込まれたバイト数と nil エラーを返す
}

func main() {
	reader.Validate(MyReader{}) // 実装をテスト
}
