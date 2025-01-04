// ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
