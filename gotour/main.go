package main

import (
	"github.com/atotto/go-tour-jp/wc"
)

// WordCount は与えられた文字列内の単語の出現回数をカウントし、map で返します。
func WordCount(s string) map[string]int {
	// 仮の実装として map[string]int{"x": 1} を返していますが、
	// 実際には入力された文字列 s を単語に分割し、
	// それぞれの単語の出現回数をカウントして
	// map[string]int を返すように実装する必要があります。
	return map[string]int{"x": 1}
}

func main() {
	// 関数のテストを実行
	// wc パッケージには、文字列を単語ごとに分割して
	// その出現回数を数える関数をテストするための便利な関数 Test が含まれています。
	wc.Test(WordCount)
}
