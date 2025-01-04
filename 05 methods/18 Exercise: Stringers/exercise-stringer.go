// IPAddr 型を実装してみましょう IPアドレスをドットで4つに区切った( dotted quad )表現で出力するため、 fmt.Stringer インタフェースを実装してください。

// 例えば、 IPAddr{1, 2, 3, 4} は、 "1.2.3.4" として出力するようにします。

package main

import (
	"fmt"
	"strconv"
	)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	s := strconv.Itoa(int(ip[0])) + "." + strconv.Itoa(int(ip[1])) + "." + strconv.Itoa(int(ip[2])) + "." + strconv.Itoa(int(ip[3]))
	return s

	// これでもいける
	// return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
