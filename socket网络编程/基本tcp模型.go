package main

import (
	"net"
	"fmt"
)

/*
tcp的两种模型 :
c/s 模型 :
(client/server)客户端/服务器
b/s 模型 :
(browser/server)浏览器/服务器
 */
 //通过net包来实现 :
func main(){
// 1) 服务器 :
// 监听						使用tcp协议模式			端口
listener,err := net.Listen("tcp","127.0.0.1:8000") //返回文件描述符 ls
check(err)
defer listener.Close()
// 阻塞等待用户连接
	conn,err1 := listener.Accept()  //连接后得到一个接口 conn
check(err1)
// 接收用户数据
buf := make([]byte,1024)
n,err2 := conn.Read(buf)  // 将数据读取到字节切片,返回读取到的长度(无数据时将阻塞)
check(err2)
fmt.Println("buf=",string(buf[:n])) //打印读取到的内容
defer conn.Close()//关闭当前用户连接
}
func check(err error){  //检查错误函数
	if err != nil{
		fmt.Println("this err is :",err)
		return
	}
	go client()  // 开协程运行客户端 连接服务器并发送数据
}
// 2) 客户端 :
func client(){
	//主动连接服务器
	conner,err := net.Dial("tcp","127.0.0.1:8000") //连接后得到接口 conner
check(err)
	//发送数据
	conner.Write([]byte("are you ok ?"))
	defer conner.Close()  //关闭连接
}
