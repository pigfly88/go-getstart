package goroutine

import (
	"testing"
	"time"
)

// 可能不会有任何输出，或者只输出了部分，因为子goroutine可能还没执行，主goroutine已经执行完了
func TestGroutine0(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
		}()
	}

	t.Log("main go end")
}

// i的值并非连续，因为等到go执行的时候i可能已经累加了很多次了。
func TestGroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
		}()
	}
	time.Sleep(time.Second * 3)
	t.Log("main go end")
}

// 乱序打出0到9
func TestGroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(j int) {
			t.Log(j)
		}(i)
	}
	time.Sleep(time.Second * 3)
	t.Log("main go end")
}
