package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestUnbufferedChannel(t *testing.T) {
	c := make(chan int, 0) //创建无缓冲的通道 c
	//内置函数 len 返回未被读取的缓冲元素数量，cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
		time.Sleep(2 * time.Second)
		c <- 3
	}()


	time.Sleep(5 * time.Second)
	for i := 0; i < 4; i++ {
		//t.Log("sleep 2 seconds")
		//time.Sleep(2 * time.Second) //延时2s
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main进程结束")
}

func send(ch chan int) {
	value := 0
	for {
		//var value = rand.Intn(100)
		ch <- value
		fmt.Printf("send %d\n", value)
		value++
	}
}

func recv(ch chan int) {
	for {
		value := <- ch
		fmt.Printf("recv %d\n", value)
		time.Sleep(time.Second)
	}
}

func TestS(t *testing.T) {
	var ch = make(chan int)
	// 子协程循环读
	go recv(ch)
	// 主协程循环写
	send(ch)
}

func TestF(t *testing.T) {
	var ball = make(chan string, 2)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	//go kickBall("李四")
	//go kickBall("王二麻子")
	//go kickBall("刘大")
	ball <- "裁判"   // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}


