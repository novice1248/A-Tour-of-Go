// インデックスや値は、 " _ "(アンダーバー) へ代入することで捨てることができます。

// for i, _ := range pow
// for _, value := range pow
// もしインデックスだけが必要なのであれば、2つ目の値を省略します。

// for i := range pow

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
