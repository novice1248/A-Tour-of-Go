// 同じ数列を保持するような、形の異なる二分木( binary tree )は多く存在し得ます。 例えば、ここに数列 1, 1, 2, 3, 5, 8, 13 を保持する2つの二分木があります。


// 2つの二分木が同じ数列を保持しているかどうかを確認する機能は、多くの言語においてかなり複雑です。

// シンプルな解決方法を記述するために、Goの並行性( concurrency )とチャネルを利用してみます。

// 例では、型を以下のように定義している tree パッケージを利用します:

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// 1. Walk 関数を実装してください。

// 2. Walk 関数をテストしてください。

// 関数 tree.New(k) は、値( k, 2k, 3k, ..., 10k )をもつ、ランダムに構造化 (しかし常にソートされています) した二分木を生成します。

// 新しいチャネル ch を生成し、 Walk を動かしましょう:

// go Walk(tree.New(1), ch)
// そして、そのチャネルから値を読み出し、10個の値を表示してください。 それは、 1, 2, 3, ..., 10 という表示になるでしょう。

// 3. Same 関数を実装してください。 t1 と t2 が同じ値を保存しているどうかを判断するため、 Walk を使ってください。

// 4. Same 関数をテストしてください。

// Same(tree.New(1), tree.New(1)) は、 true を返すように、 Same(tree.New(1), tree.New(2)) は、 false を返すようにします。

// Tree のドキュメントは こちら です。

// まだ並列処理の理解があまり足りていない
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, c chan int) {
	// returnまでdelay
	defer close(c)
	walk(t, c)
}

func walk(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}
	walk(t.Left, c)
	c <- t.Value
	walk(t.Right, c)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(3)))
	fmt.Println(Same(tree.New(1), tree.New(4)))

	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(3), tree.New(1)))
	fmt.Println(Same(tree.New(4), tree.New(1)))
}
