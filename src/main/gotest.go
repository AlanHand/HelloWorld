package main

import (
	"fmt"
	"time"
)

func main(){
	chanDemo()
}
//10个无缓存的通道
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10 ; i++  {
		channels[i] = createWorker(i)
	}
	//fmt.Println(len(channels))

	bufferedChannel()

	channelClose()
}
//缓存通道
func bufferedChannel() {
	c := make(chan int , 3)
	go worker(0 , c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}
//关闭通道
func channelClose(){
	c := make(chan int)
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
//创建通道
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id , c)
	return c
}
//读取通道数据
func worker(id int, c chan int) {
	for n := range c{
		fmt.Println(id,"--",n)
	}
}

