package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounterは同時に使用しても安全です。
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Incは指定されたキーのカウンターをインクリメントします。
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// ロックをかけることで、1度に1つのゴルーチンのみがマップc.vにアクセスできます。
	c.v[key]++
	c.mu.Unlock()
}

// Valueは指定されたキーのカウンターの現在の値を返します。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// ロックをかけることで、1度に1つのゴルーチンのみがマップc.vにアクセスできます。
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
