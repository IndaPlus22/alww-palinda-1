package fibo

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n2 := 0
	n1 := 1
	counter := -1
	return func() int {
		if counter < 1 {
			counter++
			return counter
		}

		n := n2 + n1
		n2 = n1
		n1 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
