// defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。

// defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
