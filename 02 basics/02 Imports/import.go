// このコードでは、括弧でパッケージのインポートをグループ化し、factoredインポートステートメント( factored import statement )としています。

// もちろん、複数のインポートステートメントで書くこともできます:

// import "fmt"
// import "math"
// ですが、先に示したfactoredインポートステートメントの方がより良いスタイルです。

// 訳注: ここの factored の意味は、「要素化、グループ化、整理済み」ということです。良い日本語募集中。

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
