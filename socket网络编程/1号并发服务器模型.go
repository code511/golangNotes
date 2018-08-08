package main

/*
并发服务器功能:
实现多个客户端连接处理数据
把从客户端接收到的数据变为大写并返回给客户端
支持客户端发送 exit 断开连接指令

 */

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")  //创建服务器 ,返回描述符 lis
	checker(err)  //检查错误
	fmt.Println("等待客户端连接...")
	for {
		co, err1 := lis.Accept() //等待,连接后得到操作接口 co
		checker(err1)
		go maker(co)	// 每次连接开启一个新协程 maker操作 co
	}
}
func maker(conn net.Conn) {  //形参为 Conn接口
	defer conn.Close()		//延迟执行  断开连接
	addr := conn.RemoteAddr().String() //获得 当前连接的 IP地址和端口
	fmt.Println(addr,"地址已连接")	//打印连接状态
	buf := make([]byte, 1024)	//声明 字节切片
	for {	//使用循环实现 可以传输多次数据
		n, err2 := conn.Read(buf)  //读取数据到 buf 切片 ,得到长度 n
		checker(err2)
		fmt.Printf("接收到%v数据 :%s\n", addr, string(buf[:n]))//打印数据
		if string(buf[:4]) == "exit"{ //如果收到 客户端发送 exit
			fmt.Println(addr,"断开连接") //打印指令
			return	//跳出循环 , 断开连接 , 执行结束
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //转换大写后返回给客户端
	}
}
func checker(err error) { //检查错误
	if err != nil {
		fmt.Println("this err is ", err)
		return
	}
}
