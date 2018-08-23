package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

/*
实现功能:
多个客户端可修改昵称
多个客户端上线下线提醒
客户端可以查看当前在线用户
客户端发送消息在全体客户端广播
客户端超过时间未操作将自动退出
 */
type clients struct {
	//声明结构体,保存用户信息
	c    chan string //用户数据通道
	name string      //昵称
	addr string      //网络地址
}

var onlineMap map[string]clients //声明map,保存多个用户信息
var messaage = make(chan string) //程序主通道
func main() {
	list, err := net.Listen("tcp", "127.0.0.1:8000") //创建服务器
	if err != nil {
		fmt.Println("list err is :", err) //检查错误
		return
	}
	defer list.Close() //延迟关闭服务器
	go maneger()       //开协程,操作map
	for {
		conn, err := list.Accept() //阻塞,循环等待连接
		if err != nil { //检查错误
			fmt.Println("conn err is", err)
			continue
		}
		go handleConn(conn)	//新协程,对当前连接用户进行操作(多个连接多个协程)
	}
}
func maneger() {
	onlineMap = make(map[string]clients) //初始化map
	for {
		msg := <-messaage //阻塞,等待接收数据
		for _, cli := range onlineMap { //遍历map
			cli.c <- msg //阻塞,有数据时写入
		}
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close() //延迟关闭连接
	cliAddr := conn.RemoteAddr().String() //返回当前连接网络地址
	cli := clients{make(chan string), cliAddr, cliAddr}//声明初始化结构体,保存当前用户数据
	onlineMap[cliAddr] = cli //把结构体赋值给map
	go writeMsgToClients(cli, conn) //开协程,导入结构体和当前连接处理用户数据
	messaage <- makeMsg(cli, "login") //阻塞,用户上线则提示
	cli.c <- makeMsg(cli, "current")	//提示用户本身信息
	isQuit := make(chan bool)	//检测用户是否退出通道
	hasDate := make(chan bool)	//检测用户超时自动退出中继通道
	go func() {
		buf := make([]byte, 2048)	//初始化 字节切片buf
		for {
			n, err := conn.Read(buf) //读取用户数据到 buf ,返回数据长度
			if n == 0 {			//如果长度为0
				isQuit <- true //触发用户退出通道
				fmt.Println("conn.Read err:", err)
				return
			}
			msg := string(buf[:n-1]) //将用户数据赋值给 msg
			if len(msg) == 3 && msg == "who" {	//如果数据长度等于3并且为who
				conn.Write([]byte("user list:\n"))//向当前连接用户发送数据(当前连接用户信息)
				for _, tmp := range onlineMap {	//遍历map
					word := tmp.addr + "[" + tmp.name + "]" + "\n"//用户地址+昵称
					conn.Write([]byte(word))	//发送给当前连接用户
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" { //如果长度大于等于8并且为rename(修改昵称)
				name := strings.Split(msg, "|")[1]	//提取出用户数据中要修改的名字
				cli.name = name //赋值给结构体
				onlineMap[cliAddr] = cli //再把结构体赋值给map
				conn.Write([]byte("rename ok\n"))//向用户发送提示
			} else {	//其余
				messaage <- makeMsg(cli, msg) //广播用户普通数据
			}
			hasDate <- true //向延时中继通道发送数据
		}
	}()
	for {
		select { //检测通道
		case <-isQuit:  //如果通道不阻塞
			delete(onlineMap, cliAddr) //删除map中当前连接用户信息
			messaage <- makeMsg(cli, "login out") //广播此连接用户退出提示
			return
		case <-hasDate: //如果不阻塞
		case <-time.After(time.Second * 20): //如果阻塞时间达到20秒
			delete(onlineMap, cliAddr)	//删除map中当前连接用户信息
			messaage <- makeMsg(cli, "time leave out")	//广播此连接用户超时退出提示
			return
		}
	}
}
func writeMsgToClients(cli clients, conn net.Conn) { //处理用户数据
	for msg := range cli.c { 	//遍历用户数据
		conn.Write([]byte(msg + "\n"))//向客户端广播
	}
}
func makeMsg(cli clients, msg string) (buf string) {  //导入结构体和字符串,返回字符串(广播前缀)
	buf = "[" + cli.addr + "]" + cli.name + ":" + msg //用户地址 昵称 : 提示信息
	return
}
