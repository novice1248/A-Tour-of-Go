// セミコロン(;)を省略することもできます。つまり、C言語などにある while は、Goでは for だけを使います。

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
