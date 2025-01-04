// 下記のように、インターフェースの値は、値と具体的な型のタプルのように考えることができます:

// (value, type)
// インターフェースの値は、特定の基底になる具体的な型の値を保持します。

// インターフェースの値のメソッドを呼び出すと、その基底型の同じ名前のメソッドが実行されます。

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
