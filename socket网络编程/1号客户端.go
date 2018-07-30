package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")  //主动连接服务器地址
	checkt(err)
	defer conn.Close() //执行完毕 ,关闭连接
	go func() { //开新协程,用来向服务器发送数据
		str := make([]byte, 1024)
		for {
			i, err := os.Stdin.Read(str) //键盘输入,写入 str 切片 ,返回长度
			checkt(err)
			conn.Write(str[:i]) //将 str 切片 发送到服务器
		}
	}()
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf) //从服务器接收到的数据写入 buf 切片
			if err == io.EOF{ //如果 读取完毕
				fmt.Println("断开连接")
				return
			}
			fmt.Println("服务器回复数据 :", string(buf[:n])) //打印读取到的数据
		}
}
func checkt(err error) {
	if err != nil {
		fmt.Println("this err is :", err)
		return
	}
}
