package main

import (
	"fmt"
	"time"
)

/*
select用于检测通道是否阻塞 ,用法和switch相似,但每个case必须是IO操作,case默认无break
如果case都可通信则随机执行语句,如果case全都阻塞,并且结尾没有default语句,select语句将会阻塞
 */
func main() {
	// 有时为了避免单个通道阻塞造成整个程序的阻塞,可以通过select为通道设置超时机制 :
	c := make(chan int)		//  被监听的主通道 c
	o := make(chan bool)	//  同步协程的通道 o
	go func() {				//  创建新协程
		for {
			select {		//  监听语句
			case v := <-c :	//  监听c是否阻塞(如果5秒内通信,则执行此case下语句后跳出select)
				fmt.Println(v)	//打印 c接收到的内容
				o <- true		//则向 o发送数据
				break			//跳出 (跳出select语句)
			case <-time.After(time.Second * 5) :  //同步监听延时通道 : 延时5秒
				fmt.Println("已超时")		//打印提示
				o <- true	//发送
				break		//跳出
			}
		}
	}()
	//c <- 666   //传输数据后将不会超时,程序往下执行
	<-o		// 接收通信后
	fmt.Println("程序结束") //程序运行完毕
}
