// WordCount 関数を実装してみましょう。string s で渡される文章の、各単語の出現回数のmapを返す必要があります。 wc.Test 関数は、引数に渡した関数に対しテストスイートを実行し、成功か失敗かを結果に表示します。

// strings.Fields で、何かヒントを得ることができるはずです。

// Note: このテストスイートで何を入力とし、何を期待しているかについては、golang.org/x/tour/wcを見てみてください。

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}


func main() {
	wc.Test(WordCount)
}
