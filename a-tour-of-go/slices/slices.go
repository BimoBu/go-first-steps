/*

Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

*/

package slices

import "fmt"

//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := [][]uint8{}

	for y := 0; y < dy; y++ {
		row := []uint8{}

		for x := 0; x < dx; x++ {
			value := uint8(x + y)
			row = append(row, value)
		}

		picture = append(picture, row)
	}

	return picture
}

func Test() {
	fmt.Println("Slices needs to be executed in the Go Playground")
	// pic.Show(Pic)
}
