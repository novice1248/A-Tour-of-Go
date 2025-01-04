// スライスは長さ( length )と容量( capacity )の両方を持っています。

// スライスの長さは、それに含まれる要素の数です。

// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数です。

// スライス s の長さと容量は len(s) と cap(s) という式を使用して得ることができます。

// 十分な容量を持って提供されているスライスを再スライスすることによって、スライスの長さを伸ばすことができます。 その容量を超えて伸ばしたときに何が起こるかを見るため、プログラム例でのスライスのいずれかの操作を変更してみてください。

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
