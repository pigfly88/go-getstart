package goroutine
import (
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {
	ch1 := make(chan int, 1)
	ch1 <- 10
	select {
	case <- ch1:
		// 从有缓冲chan中读取数据，由于缓冲区没有数据且没有发送者，该分支会阻塞
		fmt.Println("Received from ch")
	default:
		fmt.Println("this is default")
	}
}
