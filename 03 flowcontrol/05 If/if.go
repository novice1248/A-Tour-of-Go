// Go言語の if ステートメントは、先ほどの for ループと同様に、括弧 ( ) は不要で、中括弧 { } は必要です。

// (もうおなじみですね！)

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}