package main

import (
	"fmt"
)

type IPAddr [4]byte

// IPAddr型に対してString()メソッドを追加
// fmt.Stringerインターフェースを実装し、IPAddrを文字列に変換
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	// hostsというマップを定義。キーがstring型、値がIPAddr型
	// loopbackとgoogleDNSのエントリをhostsマップに追加。それぞれの値はIPAddr型のリテラル
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// forループを使ってhostsマップをイテレートし、fmt.Printfでキーと値を出力
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// loopback: 127.0.0.1
// googleDNS: 8.8.8.8
