package main

import (
	"fmt"
	"io"
	"strings"
)

// io.Reader インタフェースを規定しています。

// Goの標準ライブラリには、ファイル、ネットワーク接続、圧縮、暗号化などで、
// このインタフェースの 多くの実装 があります。
// io.Reader インタフェースは Read メソッドを持ちます:
// func (T) Read(b []byte) (n int, err error)
// Read は、データを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返します。
// ストリームの終端は、 io.EOF のエラーで返します。

func main() {
	// strings.NewReader("Hello, Reader!") は、io.Reader インタフェースを実装した
	// strings.Reader を作成し、
	// "Hello, Reader!" という文字列を読むことができるようにします。
	r := strings.NewReader("Hello, Reader!")

	// バイトスライス b を長さ 8 の空のバイト配列として作成
	b := make([]byte, 8)

	for {
		// r.Read(b) によって、r（strings.Reader オブジェクト）が b の長さまでデータを読み取り、
		// その読み取ったバイト数 n とエラー err を返します。
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// b[:n] は、実際に読み取った n バイト分のスライスを文字列として出力
		fmt.Printf("b[:n] = %q\n", b[:n])
		// if err == io.EOF の条件が成立すると、ストリームの終端に達したことを意味しループが終了
		if err == io.EOF {
			break
		}
	}
}

// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
// b[:n] = ""

// b は8バイトのブロックでデータを読み取るため、
// strings.Reader から「Hello, Reader!」の文字列が b に読み込まれます。
// 読み取りが進むと、文字列の最後で io.EOF が返され、ループが終了します。
