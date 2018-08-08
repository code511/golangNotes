package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	lisener, err := net.Listen("tcp", "127.0.0.1:8000") //监听,等待客户端连接
	checke(err)
	defer lisener.Close()  //延迟 关闭监听器
	fmt.Println("等待客户端连接...") //打印指示
	conn, err := lisener.Accept()	//连接成功返回接口
	checke(err)
	fmt.Println(conn.RemoteAddr().String(),"用户连接成功")  //打印指示
	defer conn.Close() //延迟关闭连接
	buf := make([]byte, 1024)
	n, err := conn.Read(buf) //接收客户端内容 ,缓存到 buf 切片
	checke(err)

	filename := string(buf[:n]) //保存为字符串 filename
	for {
		fmt.Println("准备接收文件,请输入yes/no指令")
		var cmd string
		fmt.Scan(&cmd)
		conn.Write([]byte(cmd)) //向 客户端发送指令
		if cmd == "yes" {
			receiverfile(filename, conn) //接收文件
			fmt.Println("正在接收文件...")
			return
		}
		if cmd == "no" {
			fmt.Println("已取消接收文件")
			return
		} else {
			fmt.Println("输入错误,请重新输入")
		}
	}
}
func receiverfile(filename string, conn net.Conn) { //接收文件处理 函数
	f, err := os.Create(filename)//新建文件,返回文件句柄 f
	checke(err)
	defer f.Close() //延迟关闭文件
	buf := make([]byte, 1024*4) //缓存文件内容
	for {
		n, err := conn.Read(buf) //读取从客户端接收到的内容 到 buf
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err is :", err)
			}
			return
		}
		f.Write(buf[:n])
	}
}
func checke(err error) {
	if err != nil {
		fmt.Println("this err is :", err)
		return
	}
}
