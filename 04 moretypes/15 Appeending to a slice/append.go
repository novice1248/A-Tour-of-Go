// スライスへ新しい要素を追加するには、Goの組み込みの append を使います。 append についての詳細は documentation を参照してみてください。

// func append(s []T, vs ...T) []T
// 上の定義を見てみましょう。 append への最初のパラメータ s は、追加元となる T 型のスライスです。 残りの vs は、追加する T 型の変数群です。

// append の戻り値は、追加元のスライスと追加する変数群を合わせたスライスとなります。

// もし、元の配列 s が、変数群を追加する際に容量が小さい場合は、より大きいサイズの配列を割り当て直します。 その場合、戻り値となるスライスは、新しい割当先を示すようになります。

// (スライスについてより詳しく学ぶには、Slices: usage and internalsを読んでみてください)

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
