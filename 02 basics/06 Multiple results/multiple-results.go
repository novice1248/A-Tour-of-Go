// 関数は複数の戻り値を返すことができます。

// この swap 関数は2つの string を返します。 とても簡単に交換できますね！

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
