package main

import (
	"net/http"
	"fmt"
)
// http服务器 :
func main(){
//注册处理函数,当用户连接,自动调用指定的处理函数
http.HandleFunc("/",handConn)
//监听绑定
http.ListenAndServe(":8000",nil)
}
// w : 给客户端发送数据
// r : 读取客户端相关的数据
func handConn(w http.ResponseWriter,r *http.Request)  {
w.Write([]byte("hello go")) //向客户端发送数据
fmt.Println("r.Method =",r.Method)//打印: 客户端的请求方法
fmt.Println("r.URL =",r.URL)			   //客户端的地址处理函数
fmt.Println("r.Header =",r.Header) 	   //客户端请求包的头部信息(map)
fmt.Println("r.Body =",r.Body)		   //请求包主体
}