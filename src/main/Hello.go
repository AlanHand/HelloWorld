package main

import (
	"fmt"
)

//定义全局变量
var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d /%s"
)

type Shaper interface {
	Area() float32
}

type Square struct{
	side float32
}
type Rectangle struct{
	length,width float32
}
//类型Square实现了Shaper的Area方法
func (sq *Square) Area() float32{
	return sq.side * sq.side
}
//类型Rectangle实现了Shaper的Area方法
func (r *Rectangle) Area() float32{
	return r.length * r.width
}

func main() {
	//fmt.Println("Please enter you full name:")
	//fmt.Scanln(&firstName, &lastName)
	//
	//fmt.Print("Hi %s %s!\n", firstName, lastName)

	////协程的使用
	//ch:=make(chan string)
	//go sendData(ch)
	//go getData(ch)
	//time.Sleep(1e9 * 2)
	//fmt.Println("退出")r := &Rectangle{5,3}
	//	//q := &Square{5}
	//	//
	//	////类型转换为接口
	//	//shapes := []Shaper{r, q}
	//	//for n :=range shapes{
	//	//	fmt.Println(shapes[n])
	//	//	fmt.Println(shapes[n].Area())
	//	//}

	//接口的使用,前提是类型实现了接口的方法,按照这个思路就是说一切的类型只要实现了接口的方法都可以调用接口的方法
	//

	out := make(chan int)
	//对于通道来讲,发送数据的通道必须提前有接收数据的通道进行数据的消耗,否则发送数据的通道就会一直等待(死锁),因为这里创建的通道默认是没有缓存的
	//因此若是给通道设置一个容量值,那么只要通道中有数据就不会造成阻塞了
	//out := make(chan int , value)
	fmt.Println("-----")
	go f1(out)
	out <- 2

}
func f1(in chan int){
	fmt.Println(<- in)
}
/**
 * 协程中接收数据
 */
func getData(ch chan string) {
	var input string
	//time.Sleep(1e9 * 2) // 这里如果设置睡眠2秒则和main中同时退出,下面的ch协程就随着main退出了
	for{
		input = <- ch
		fmt.Printf("%s",input)
	}
}
/**
 * 协程中发送数据
 */
func sendData(ch chan string){
	ch <- "Alan\t\n"
	ch <- "Tom\t\n"
	ch <- "Tony\t\n"
}

