package goroutine

import (
	"sync"
	"testing"
)

func TestGoroutineSum(t *testing.T) {
	c := make(chan int, 10)
	var wg sync.WaitGroup

	// 计算1+2+3...+99的总和
	for i := 0; i < 10; i++ {
		t.Log("i =", i)
		// 第一个goroutine计算1+2+3...+9，第二个goroutine计算10+11+12...+19，以此类推
		wg.Add(1)
		go func(j int) {
			//t.Log("j =", j)
			sumPart := 0
			max := j + 10
			for ; j < max; j++ {
				sumPart += j
				//t.Log(j, sumPart)
			}
			c <- sumPart
			//t.Log(j, sumPart)
			wg.Done()
		}(i * 10)
	}

	wg.Wait()

	close(c) // 必须调用close，否则range遍历通道将无法结束，会出现死锁
	sum := 0
	for data := range c {
		//t.Log(data)
		sum += data
	}
	t.Log(sum)
}
