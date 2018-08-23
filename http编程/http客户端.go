package main

import (
	"net/http"
	"fmt"
)

// http客户端 :
func main(){
resp,err := http.Get("http://www.baidu.com") //用get方法请求百度服务器
if err != nil{		//检查错误
	fmt.Println("get err:",err)
	return
}
defer resp.Body.Close() //延迟关闭请求
fmt.Println("Status =",resp.Status)	// 打印响应包 响应状态
fmt.Println("StatusCode =",resp.StatusCode)	//	 响应码
fmt.Println("Header =",resp.Header)		 	//   头部信息
// 打印响应包body(包体) :
buf := make([]byte,2048*2) //声明 字节切片 buf
var tmp string			//声明 字符串 tmp
for {
	n,err := resp.Body.Read(buf) //将包体读取到 buf ,返回读取长度 n
	if n == 0{		//如果读取完毕
		fmt.Println("resp.Body.Read err :",err)
		break
	}
	tmp += string(buf[:n])//将 buf 拼接为 字符串 tmp
}
fmt.Println("tmp :",tmp)//打印包体
}