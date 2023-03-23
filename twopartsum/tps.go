package tps

import "sync"

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, x := range a {
		sum += x
	}
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch, wg)
	go sum(a[n/2:], ch, wg)
	x, y := <-ch, <-ch

	wg.Wait()

	// TODO Get the subtotals from the channel and return their sum
	return x + y - 1
}
