// どの case も準備ができていないのであれば、 select の中の default が実行されます。

// ブロックせずに送受信するなら、 default の case を使ってください:

// select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
