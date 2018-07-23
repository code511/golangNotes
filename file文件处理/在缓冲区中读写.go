package main

import (
	"bufio"
	"os"
	"fmt"
)

/* 缓冲区(内存) :
在缓冲区中读写是为了提高程序运行速度,最后如果需要才把需要的数据写入磁盘中
终端,文件,网络数据 都可以在缓冲区中读写*/
func main(){
		fmt.Println("请输入字符串: ")		// 提示
		inputReader := bufio.NewReader(os.Stdin)  // 初始化为 从标准输入 读取(也可设置为从文件,网络读取)
		input, err := inputReader.ReadString('\n')	 //从标准输入读取字符串类型赋值给 input
		if err == nil {
			fmt.Printf("输入的是: %s\n", input)	//如果无错误,打印结果 
		}

}