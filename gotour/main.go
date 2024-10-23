package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("When's Saturday?")
	// 現在の曜日を time.Weekday 型で返す。
	// 曜日に対応する整数値（0: 日曜日, 1: 月曜日, …, 6: 土曜日）。
	today := time.Now().Weekday()
	// time.Weekday 型の定数で、6 という整数値に対応
	// 6（土曜日）に対する switch 文が実行
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
