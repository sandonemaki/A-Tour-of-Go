package main

import (
	"io"
	"os"
	"strings"
)

// io.Reader は Go の標準インタフェースで、
// データをストリームから読み取るためのメソッド Read を持っています。
// このメソッドは、次のように定義されます：
// func (T) Read(b []byte) (n int, err error)

// b []byte：読み取ったデータを格納するためのバイトスライス（配列のようなもの）。
// n int：実際に読み取ったバイト数。
// err error：エラーがあれば返します。ストリームの終わりに達した場合は io.EOF が返されます。

// rot13Readerの目的
// rot13Reader 型は、io.Reader をラップし、
// 読み取ったデータをROT13という暗号に変換して出力すること
// rot13Reader は単純に io.Reader を内包する構造体
type rot13Reader struct {
	r io.Reader
}

// rot13Readerの目的
// Read メソッドを実装し、ROT13 暗号を適用します。
// ROT13 変換は、アルファベットの各文字を13文字後ろにずらします。
func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

// ROT13 変換を適用するヘルパー関数
func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
