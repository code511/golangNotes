package main

import "fmt"

/*
channel :
chan(引用类型)是协程之间用来接收和发送数据的通道,用来协调和同步协程之间的执行和数据传输,可以用make创建,close()关闭
类型分为 : 双向(默认), 单向(双向可转单向,单向不可转双向)
无缓存(默认,无数据传输时协程阻塞),有缓存(容量满发送和空接收将阻塞)
属性有 : 数据类型,长度,容量
chan 零值为 nil , 通过 <- 操作符发送和接收 ,通过 range来迭代遍历
 */
func main(){
// 语法 :
// 1) 双向无缓存 :
xx := make(chan int)
// 2) 双向有缓存 :
xx1 := make(chan int,5) //缓存容量 5
// 3) 单向无缓存 :
var dx chan <- int  //只能发送
var dx1 <- chan int	//只能接收
// 4) 双向转单向 :
var dx2 chan <- int =xx  // 转单向发送
var dx3 <- chan int = xx1 //转单向接收
// 传输 :
// 1) 发送 :
xx <- 666
dx <- 888
dx2 <- 555
// 2) 接收 :
a := <- xx1
b := <- dx3
fmt.Println(a,b)
// 3) 丢弃(无变量接收) :
<-xx
<-dx1
// 操作 :
// 1) 关闭通道 :
close(xx1)
// 2) 检查通道是否关闭或为空 :
if i,ok := <- xx1 ; ok == true{
	fmt.Println(i)
}
}