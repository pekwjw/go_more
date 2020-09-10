package main

import (
	"fmt"
	"time"
)

func ProduceNoBuf(ch  chan <- int) {
	for i:=0; i<10; i++ {
		ch <- i
		fmt.Println("send: ", i)
	}
}

func ConsumerNoBuf(ch <- chan int) {
	for i:= 0; i<10; i++ {
		v := <- ch
		fmt.Println("receive: ", v)
	}
}

// NoBuf 无缓冲生产者消费者模型
func NoBuf() {
	ch := make(chan int)
	go ProduceNoBuf(ch)
	go ConsumerNoBuf(ch)
	time.Sleep(time.Second*1)
}

func ProduceQ(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send To Queue:", i)
	}
}

func ConsumerQ(ch <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Queue Receive:", v)
	}
}

// Queue 有缓冲队列生产者消费者模型
func Queue() {
	ch := make(chan int, 10)
	go ProduceQ(ch)
	go ConsumerQ(ch)
	time.Sleep(1 * time.Second)
}

func main() {
	NoBuf()
	Queue()
}
