package goroutine

import (
	"sync/atomic"
	"testing"
	"time"
)

/**
按顺序执行goroutine，下面的代码按顺序输出了0到9
利用原子操作累加一个计数器，通过计数器和当前循环的i做比较来判断是否该当前的goroutine执行了
*/

var count uint32

func TestInorder(t *testing.T) {
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			trigger(i, func() {
				t.Log(i)
			})
		}(i)
	}
	trigger(10, func() {})
}

func trigger(i uint32, fn func()) {
	for {
		if n := atomic.LoadUint32(&count); n == i { // 计数器和i相等则执行，否则陷入for循环等待直到 n==i
			fn()
			atomic.AddUint32(&count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}
