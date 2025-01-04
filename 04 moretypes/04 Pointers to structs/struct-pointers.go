// structのフィールドは、structのポインタを通してアクセスすることもできます。

// フィールド X を持つstructのポインタ p がある場合、フィールド X にアクセスするには (*p).X のように書くことができます。 しかし、この表記法は大変面倒ですので、Goでは代わりに p.X と書くこともできます。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
