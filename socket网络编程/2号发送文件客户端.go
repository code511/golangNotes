package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func main() {
	fmt.Println("请输入要发送的文件path(+文件名) :") //启动提示
	var path string                      //声明变量 path
	fmt.Scan(&path)                      // 键盘输入后 赋值给 path
	info, err := os.Stat(path)           //获取文件属性结构体 info
	if err != nil { //检查错误
		fmt.Println("文件path输入错误 :", err)
		return
	}
	conn, err := net.Dial("tcp", "127.0.0.1:8000", ) //主动连接 服务器地址
	checks(err)
	defer conn.Close()                       //延迟关闭 连接
	_, err = conn.Write([]byte(info.Name())) //向服务器发送文件名字
	checks(err)
	fmt.Println("等待服务器指令...")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf) //接收服务器回复,读取到 buf 切片中, 舍弃读取长度
	checks(err)
	if string(buf[:n]) == "yes" { //如果服务器回复 "ok"
		sendfile(path, conn) //开始发送文件
		fmt.Println("开始发送...")
	}
	if string(buf[:n]) == "no"{
		fmt.Println("服务器已拒绝接收文件")
	}
}
func sendfile(path string, conn net.Conn) { //发送文件函数 sendfile
	f, err := os.Open(path)  	//只读方式打开文件
	checks(err)
	defer f.Close()		//延迟 关闭文件
	buf := make([]byte, 1024*4)	//创建数组,缓存文件内容
	for {
		n, err := f.Read(buf)	//读取文件内容
		if err != nil {			//检查如果读取完毕或错误
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err :", err)
			}
			return
		}
		conn.Write(buf[:n]) // 向服务器发送文件内容
	}
}
func checks(err error) {	//检查错误函数
	if err != nil {
		fmt.Println("this err is :", err)
	}
}
