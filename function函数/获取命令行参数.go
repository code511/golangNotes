package main

import (
	"os"
	"fmt"
)

/* 获取命令行参数 :
	1) 通过OS包来实现
	2) 以字符串切片的方式传递,输入的参数以空格分割(参数从1开始)
 */
func main(){
	date := os.Args	//用date 接收用户输入的参数(字符串切片)
	for i,s:=range date {
		fmt.Printf("输入的第%d个,内容:%s\n",i,s)
	}
}		//执行命令 : go run 获取命令行参数.go  i love you

