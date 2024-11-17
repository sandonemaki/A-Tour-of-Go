package main

import (
	"fmt"
	"sync"
)

// Fetcherインターフェースは、URLのボディとそのページで見つかったURLのスライスを返します。
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// VisitedMapは、訪れたURLを記録するスレッドセーフな構造体です。
type VisitedMap struct {
	mu      sync.Mutex
	visited map[string]bool
}

// NewVisitedMapは、VisitedMapの新しいインスタンスを生成します。
func NewVisitedMap() *VisitedMap {
	return &VisitedMap{visited: make(map[string]bool)}
}

// MarkVisitedは、指定されたURLを訪問済みに設定。
// sync.Mutex によってこの操作はスレッドセーフに行われる
func (v *VisitedMap) MarkVisited(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.visited[url] = true
}

// IsVisitedは、指定されたURLが訪問済みかどうかを返します。
func (v *VisitedMap) IsVisited(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.visited[url]
}

// Crawlはfetcherを使用して、urlから始まるページを再帰的にクロールし、最大でdepthの深さまで探索します。
func Crawl(url string, depth int, fetcher Fetcher, visited *VisitedMap, wg *sync.WaitGroup) {
	defer wg.Done() // goroutine の完了を WaitGroup に通知します
	// 深さが 0 以下か、その URL がすでに訪問済みかどうかを確認
	// 初回は深さが 4 で、まだ訪問していない URL なので処理が続行される
	if depth <= 0 || visited.IsVisited(url) {
		return
	}

	visited.MarkVisited(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// 子の URL を並行してクロールします。
	for _, u := range urls {
		wg.Add(1) // カウンタを 1 増やす
		// 新しい goroutine を起動して各リンクを再帰的にクロール
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}

func main() {
	// visited という VisitedMap のインスタンスを作成
	// URL がすでに訪問されたかを記録するため
	visited := NewVisitedMap()
	// URL をクロールする複数の goroutine を管理
	var wg sync.WaitGroup

	// 初回クロールの開始
	// WaitGroup のカウンタを 1 増やす
	wg.Add(1)
	// Crawl 関数が並行して実行される goroutine を起動
	go Crawl("https://golang.org/", 4, fetcher, visited, &wg)

	// カウンタが 0 になるまでプログラムをブロック（待機）
	// Crawl 関数が終了してカウンタが 0 になると、wg.Wait() が解除され、main 関数が続行
	wg.Wait()
	fmt.Println("Crawling complete")
}

// fakeFetcherは、事前に用意された結果を返すFetcherです。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcherは、データが設定されたfakeFetcherです。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// found: https://golang.org/ "The Go Programming Language"
// not found: https://golang.org/cmd/
// found: https://golang.org/pkg/ "Packages"
// found: https://golang.org/pkg/os/ "Package os"
// found: https://golang.org/pkg/fmt/ "Package fmt"
// Crawling complete
