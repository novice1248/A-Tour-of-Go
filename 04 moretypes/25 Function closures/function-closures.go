// Goの関数は クロージャ( closure ) です。 クロージャは、それ自身の外部から変数を参照する関数値です。 この関数は、参照された変数へアクセスして変えることができ、その意味では、その関数は変数へ"バインド"( bind )されています。

// 例えば、 adder 関数はクロージャを返しています。 各クロージャは、それ自身の sum 変数へバインドされます。

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
