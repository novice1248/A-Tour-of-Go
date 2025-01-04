// Sqrt 関数を 以前の演習 からコピーし、 error の値を返すように修正してみてください。

// Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

// 新しい型:

// type ErrNegativeSqrt float64
// を作成してください。

// そして、 ErrNegativeSqrt(-2).Error() で、 "cannot Sqrt negative number: -2" を返すような:

// func (e ErrNegativeSqrt) Error() string
// メソッドを実装し、 error インタフェースを満たすようにします。

// 注意: Error メソッドの中で、 fmt.Sprint(e) を呼び出すことは、無限ループのプログラムになることでしょう。 最初に fmt.Sprint(float64(e)) として e を変換しておくことで、これを避けることができます。 なぜでしょうか？

// 負の値が与えられたとき、 ErrNegativeSqrt の値を返すように Sqrt 関数を修正してみてください。

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

var eps float64 = math.Nextafter(1, 2) - 1

func Sqrt(x float64) (float64, error) {
	//この書き方は実務的ではない if err := ?; err != nil {}の方がいい
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev_z := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= eps {
			break
		}
		prev_z = z
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
