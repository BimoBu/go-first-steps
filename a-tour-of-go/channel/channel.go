// No excercise, just seeing how the channel values are populated and consumed

package channel

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("Putting x into channel, x = %v\n", x)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Test() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Printf("Printing i to console, i = %v\n", i)
		fmt.Println(i)
	}
}
