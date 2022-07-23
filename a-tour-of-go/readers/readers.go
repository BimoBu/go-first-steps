/*

Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

*/

package readers

import (
	"fmt"
	// "golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	a, n := byte('A'), 0

	for i := range b {
		b[i] = a
		n++
	}

	return n, nil
}

func Test() {
	fmt.Println("Readers needs to be executed in the Go Playground")
	// reader.Validate(MyReader{})
}
