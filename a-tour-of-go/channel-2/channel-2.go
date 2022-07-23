// No excercise, just seeing how the channel values are populated and consumed

package channel2

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("new")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Test() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("reading")
			fmt.Println(<-c)
		}
		fmt.Println("quitting")
		quit <- 0
	}()
	fibonacci(c, quit)
}
