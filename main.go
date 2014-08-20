package main

import (
	"fmt"
)

func main() {
	// Set up the pipeline.
	c := gen(2, 3) // series of numbers
	out := sq(c)   // pass the channel 'c' into the sq function

	// Consume the output.
	fmt.Println(<-out) // 4.  get 1st result out of channel 'out'
	fmt.Println(<-out) // 9.  get 2nd result out of channel 'out'
}

func gen(nums ...int) <-chan int { // returns a channel of type int
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out // returns channel 'out'
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() { // call goroutine
		for n := range in { // iterate over range
			out <- n * n // put square of n in buffer
		}
		close(out) //
	}() // immediate invoke anonymous function
	return out
}
