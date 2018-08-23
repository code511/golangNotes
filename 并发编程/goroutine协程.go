package main

import (
	"time"
	"fmt"
)

/*
创建一个协程,只需要一个 go 关键字
一个cup核心开启一个线程,一个线程可以开启N个协程(通过时间片轮转执行)
goroutine协程是GO语言并行设计的核心,比线程更小,更高效,更易用,更方便.
普通协程包含在 main协程(主协程)内,随 main协程开始结束.
 */
func newTask() {
	for {
		fmt.Println("newTask routine")
		time.Sleep(time.Second)	//用时间延迟器阻塞该协程 1秒钟 ,达到交替执行的目的
	}
}
func main() {
	go newTask() //遇到go 启动新协程运行这个函数

	for {		//同时运行主协程
		fmt.Println("main routine")
		time.Sleep(time.Second) // 用时间延迟器阻塞该协程 1秒钟
	}
}
