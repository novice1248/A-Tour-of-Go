// 型アサーション は、インターフェースの値の基になる具体的な値を利用する手段を提供します。

// t := i.(T)
// この文は、インターフェースの値 i が具体的な型 T を保持し、基になる T の値を変数 t に代入することを主張します。

// i が T を保持していない場合、この文は panic を引き起こします。

// インターフェースの値が特定の型を保持しているかどうかを テスト するために、型アサーションは2つの値(基になる値とアサーションが成功したかどうかを報告するブール値)を返すことができます。

// t, ok := i.(T)
// i が T を保持していれば、 t は基になる値になり、 ok は真(true)になります。

// そうでなければ、 ok は偽(false)になり、 t は型 T のゼロ値になり panic は起きません。

// この構文と map から読み取る構文との類似点に注意してください。

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
