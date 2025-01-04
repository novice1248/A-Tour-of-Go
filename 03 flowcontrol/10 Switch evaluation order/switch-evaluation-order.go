// switch caseは、上から下へcaseを評価します。 caseの条件が一致すれば、そこで停止(自動的にbreak)します。

// (例えば、

// switch i {
// case 0:
// case f():
// }
// では、 i==0 であれば、 case 0 でbreakされるため f は呼び出されません)

// Note: Go playground上の時間は、いつも "2009-11-10 23:00:00 UTC" です。 この値の意味は、読者の楽しみのために残しておきます(^^)

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
