package main

import "fmt"

// 斐波那切数列 :	 1,1,2,3,5,8,13
func main(){
send := make(chan int)  // 传送数字通道 send
spring := make(chan bool)// 触发通道 spring
go func() {		//新协程 匿名函数
	for i:=0;i<8 ;i++  {   // send 循环接收 8次数据
		number := <-send	//每次循环接收一次数据
		fmt.Println(number)// 打印数据
	}
	spring <- true  // 8 次发送完毕 ,发送数据触发通信
}()
add(send,spring)	//调用函数
}
func add (send chan<-int,spring <-chan bool){ //运算函数 ,形参转换为 单向管道
	x,y:=1,1	//原始数字
	for  {
		select {
		case send<-x : //如果管道接收数据
			x,y = y,x+y //运行运算
		case b := <-spring: //如果管道发送数据
			fmt.Println("执行完毕",b) //打印结果
		return //结束循环
		}
	}
}