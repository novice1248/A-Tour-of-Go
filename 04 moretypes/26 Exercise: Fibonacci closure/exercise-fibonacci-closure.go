// 関数を用いた面白い例を見てみましょう。

// fibonacci (フィボナッチ)関数を実装しましょう。この関数は、連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返します。

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	m1, m2 := 1, 0
	return func() int {
		m1, m2 = m2, m1 + m2
		return m1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
