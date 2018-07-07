package main

import "fmt"

/* 内存和地址 :
没一个变量都有两个属性 :值 和 内存地址(也叫指针,用 &变量名 取地址)
 */
func main(){
	a := 10
	fmt.Printf("a的值为%d\n",a)		//打印值
	fmt.Printf("a的内存地址为%v\n",&a) //打印地址(指针)
}