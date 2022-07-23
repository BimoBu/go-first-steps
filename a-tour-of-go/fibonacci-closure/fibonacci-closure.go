/*

Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

*/

package fibonacciClosure

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	callNumber := 0

	twoAgo := 0
	oneAgo := 1

	return func() int {
		defer func() {
			callNumber++
		}()

		// if callNumber == 0 || callNumber == 1 {
		// 	return callNumber
		// }

		tmpOneAgo := oneAgo
		oneAgo = twoAgo + oneAgo
		twoAgo = tmpOneAgo

		return oneAgo
	}
}

func Test() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
