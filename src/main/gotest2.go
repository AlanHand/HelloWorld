package main

import (
	"sync"
	"fmt"
)

type worker2 struct{
	in 		chan int
	done 	func()
}

func main() {
	chanDemo2()

}
func chanDemo2() {
	var wg sync.WaitGroup
	var workers [2]worker2

	//创建10个协程
	for i := 0 ; i < 2 ; i++{
		workers[i] = createWorker2(i,&wg)
	}

	//添加一个计数
	wg.Add(4)

	//往通道里面传入数据
	for i,worker := range workers{
		worker.in <- 'a' + i
 	}
	for i,worker := range workers{
		worker.in <- 'A' + i
	}
	//阻塞直到所有协程任务完成
	wg.Wait()
}
//创建协程
func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in : make(chan int),
		done : func() {
				wg.Done()
			},
	}
	go doWork(id , w)
	return w
}
//协程开始工作,等待接收数据
func doWork(id int, w worker2) {
	for n := range w.in{
		fmt.Println(id,"--",n)
		//解锁
		w.done()
	}
}
