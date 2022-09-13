package channel

import "fmt"

/**
receiveChan <-chan int 只能收
sendChan chan<- int 只能发
把<想象成一个张开口的嘴巴，<-chan可以很形象地想象成是别人张开嘴巴吃了chan，所以只能从chan里面收数据。
而chan<-是chan自己张开嘴巴把数据吃进去，所以只能发数据给chan。
*/
func counter(sendChan chan<- int) {
	for x := 0; x < 100; x++ {
		sendChan <- x
	}
	close(sendChan)
}

func squarer(receiveChan <-chan int, sendChan chan<- int) {
	for v := range receiveChan {
		sendChan <- v * v
	}
	close(sendChan)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
