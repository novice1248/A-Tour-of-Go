// 下の2つの呼び出しを比べると、ポインタを引数に取る ScaleFunc 関数は、ポインタを渡す必要があることに気がつくでしょう:

// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
// メソッドがポインタレシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます:

// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK
// v.Scale(5) のステートメントでは、 v は変数であり、ポインタではありません。 メソッドでポインタレシーバが自動的に呼びだされます。 Scale メソッドはポインタレシーバを持つ場合、Goは利便性のため、 v.Scale(5) のステートメントを (&v).Scale(5) として解釈します。

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
