// 明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論されます。

// 右側の変数が型を持っている場合、左側の新しい変数は同じ型になります:

// var i int
// j := i // j is an int
// 右側に型を指定しない数値である場合、左側の新しい変数は右側の定数の精度に基いて int, float64, complex128 の型になります:

// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128
// 例のコードにある変数 v の初期値を変えて、型がどのように変化するかを見てみてください。

package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
