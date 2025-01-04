// ポインタレシーバでメソッドを宣言できます。

// これはレシーバの型が、ある型 T への構文 *T があることを意味します。 （なお、 T は *int のようなポインタ自身を取ることはできません）

// 例では *Vertex に Scale メソッドが定義されています。

// ポインタレシーバを持つメソッド(ここでは Scale )は、レシーバが指す変数を変更できます。 レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的です。

// Scale の宣言(line 16)から * を消し、プログラムの振る舞いがどう変わるのかを確認してみましょう。

// 変数レシーバでは、 Scale メソッドの操作は元の Vertex 変数のコピーを操作します。 （これは関数の引数としての振るまいと同じです）。 つまり main 関数で宣言した Vertex 変数を変更するためには、Scale メソッドはポインタレシーバにする必要があるのです。

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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
