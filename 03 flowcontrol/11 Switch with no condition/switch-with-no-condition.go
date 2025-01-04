// 条件のないswitchは、 switch true と書くことと同じです。

// このswitchの構造は、長くなりがちな "if-then-else" のつながりをシンプルに表現できます。

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
