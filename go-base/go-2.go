package main

import (
	"fmt"
	//"time"
	"runtime"
)




var ch = make(chan int)
func test(){
	for i:=0; i<100 ;i++{
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d  ", i)
	}
	ch<-0  //方案2
}

func main()  {
	//使用两个核
	//runtime.GOMAXPROCS(2)   //使用最多2核
	go test()
	go test()
	//time.Sleep(time.Second) //注释掉此行 则不会打印数据到控制台,main比test()先退出了
	 <-ch  //方案2
	 <-ch //方案2
}


//2.使用信道:channel
var quit chan int = make(chan int)
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}
/*
func main() {
	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()
	//保证goroutine都执行完，主线程才结束
	for i := 0; i < 2; i++ {
		<- quit
	}
}*/
