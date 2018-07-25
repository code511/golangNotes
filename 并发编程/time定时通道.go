package main

import (
	"time"
	"fmt"
)

/*
定时通道分为两种 :
1) 单次定时 Timer or After  : 定时后产生一次事件
2) 循环定时 Ticker :	定时后循环产生事件
 */
func main(){
// 单次定时:
// 1) 普通定时 :
time.Sleep(time.Second*2)	//延时 2秒
fmt.Println("开始计时")		//执行事件
// 2) Timer 和 After :
timers := time.NewTimer(time.Second*2)  //创建定时 2秒通道,两秒后timer通道接收内容(当前时间)
timers1 :=time.After(time.Second*5)		//与Timer效果相同
now := time.Now()	//获得当前时间
fmt.Println(now)	//打印当前时间
a := <- timers.C	//执行到创建管道 2秒后接收数据(当前时间)赋值给 a ,通道接收内容前后位阻塞状态
fmt.Println(a)		//打印结果
a1 := <-timers1		//执行到创建管道 5秒后接收数据 赋值给 a1
fmt.Println(a1)		//打印结果
//	循环定时 Ticker :
timers2 := time.NewTicker(time.Second*2)  //创建一个循环定时管道 timers2 ,开始计时
		i := 0 //创建计数器 i
	for{
		<- timers2.C  //计时两秒后 接收到数据,丢弃掉
		i++			//每次循环计数器 +1
		fmt.Println(i) //打印每次结果
		if i == 5 {		//如果计数器 i 达到 5
		timers2.Stop()  // 循环通道 停止
		break		//跳出循环
		}
	}
// 重置定时管道: 重置后单独重新计时,前后互不影响
timers.Reset(time.Second*3)  //重置为 3秒
a2 := <- timers.C
fmt.Println(a2)
// 关闭定时管道: 关闭后管道将失效
timers.Stop()
}