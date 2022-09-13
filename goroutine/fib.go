package main

import (
	"fmt"
	"time"
)

/*func TestGoroutine(t *testing.T) {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func(j int) {
			sumPart := 0
			max := j + 10
			for ; j < max; j++ {
				sumPart += j
				t.Log(j, sumPart)
			}
			c <- sumPart
		}(i * 10)
		close(c)
	}

	sum := 0
	//for i := 0; i < 10; i++ {
	if tmp, ok := <-c; ok {
		sum += tmp
	}
	t.Log(sum)
}*/

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
