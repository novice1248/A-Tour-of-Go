// よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

// 例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

// io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

// rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。

//このコードはネットから拾ってきたものであまり現状理解できていない
//後から復習すべき
package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	encryptionKey = 13
	table         = alphabet + alphabet
)

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := range p {
		p[i] = r.rotate(p[i], encryptionKey)
	}
	return n, nil
}

func (r *rot13Reader) rotate(b byte, key int) byte {
	ch := rune(b)
	if !unicode.IsLetter(ch) {
		return b
	}

	key = key % len(alphabet)

	isUpper := false
	if unicode.IsUpper(ch) {
		isUpper = true
		ch = unicode.ToLower(ch)
	}

	idx := byte(ch) - 'a' + byte(key)
	rotatedCh := rune(table[idx])
	if isUpper {
		rotatedCh = unicode.ToUpper(rotatedCh)
	}
	return byte(rotatedCh)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
