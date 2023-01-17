package goroutine

import (
	"testing"
	//"time"
)

func TestWaitByChannel(t *testing.T) {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			t.Log(i)
			sign <- struct{}{} // 空结构体不占用内存空间（0字节），这个值在整个 Go 程序中永远都只会存在一份。
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
