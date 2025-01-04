// もっともよく使われているinterfaceの一つに fmt パッケージ に定義されている Stringer があります:

// type Stringer interface {
//     String() string
// }
// Stringer インタフェースは、stringとして表現することができる型です。 fmt パッケージ(と、多くのパッケージ)では、変数を文字列で出力するためにこのインタフェースがあることを確認します。

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
