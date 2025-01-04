// インターフェース自体の中にある具体的な値が nil の場合、メソッドは nil をレシーバーとして呼び出されます。

// いくつかの言語ではこれは null ポインター例外を引き起こしますが、Go では nil をレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的です(この例では M メソッドのように)。

// 具体的な値として nil を保持するインターフェイスの値それ自体は非 nil であることに注意してください。

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
