package main

import "fmt"

// https://scrapbox.io/maki-Py/Go%E3%81%AE%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E3%81%AE%E5%AE%B9%E9%87%8F%E3%81%AE%E5%AE%9A%E7%BE%A9

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
