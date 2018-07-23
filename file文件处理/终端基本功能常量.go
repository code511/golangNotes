package main

import (
	"fmt"
	"os"
)

func main(){
// 操作文件终端句柄常量(默认开启,可关闭) :
//标准输入 os.Stdin
	fmt.Fprint(os.Stdin,"标准输入")//等同于 fmt.Scan(&a) ,输入 "标准输入"
	//os.Stdin.Close()  关闭功能
//标准输出 os.Stdout
	fmt.Fprint(os.Stdout,"标准输出\n")//等同于 fmt.Print("标准输出")
	//os.Stdout.Close()	关闭功能
//标准错误输出 os.Stderr
	fmt.Fprint(os.Stderr,"标准错误输出")//同上
	//os.Stderr.Close() 关闭功能
}