// switch ステートメントは if - else ステートメントのシーケンスを短く書く方法です。

// Go の switch は C や C++、Java、JavaScript、PHP の switch と似ていますが、 Go では選択された case だけを実行してそれに続く全ての case は実行されません。 これらの言語の各 case の最後に必要な break ステートメントが Go では自動的に提供されます。 もう一つの重要な違いは Go の switch の case は定数である必要はなく、 関係する値は整数である必要はないということです。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
