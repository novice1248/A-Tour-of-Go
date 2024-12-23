// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。

// この例では、

// x int, y int
// を

// x, y int
// へ省略できます。

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
