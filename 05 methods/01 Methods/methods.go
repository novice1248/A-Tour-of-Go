// Goには、クラス( class )のしくみはありませんが、型にメソッド( method )を定義できます。

// メソッドは、特別なレシーバ( receiver )引数を関数に取ります。

// レシーバは、 func キーワードとメソッド名の間に自身の引数リストで表現します。

// この例では、 Abs メソッドは v という名前の Vertex 型のレシーバを持つことを意味しています。

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}