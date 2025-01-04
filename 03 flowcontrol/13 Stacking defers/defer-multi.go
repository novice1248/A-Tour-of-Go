// defer へ渡した関数が複数ある場合、その呼び出しはスタック( stack )されます。 呼び出し元の関数がreturnするとき、 defer へ渡した関数は LIFO(last-in-first-out) の順番で実行されます。

// defer ステートメントについてさらに学ぶには、 こちら(blog post)を読んでみてください。

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
