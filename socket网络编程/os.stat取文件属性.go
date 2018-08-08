package main

import (
	"os"
	"fmt"
)

/*
os.stat 得到文件路径后,用 os.stat方法获取文件属性:
Name()  :文件名字 (不含扩展名)
Size()	:文件大小
SYS()	:文件底层数据来源
Mode()	:文件的模式
Modtime():文件的修改时间
Isdir() :判断是否是目录
 */

func main(){
// 示例为从命令行参数获取,也可从fmt.Scan()获取获取文件path
filer := os.Args  //获取命令行数据 (输入的文件路径)
if len(filer) != 2{  // filer[0]返回执行本身,filer[1]返回输入的命令行数据
	fmt.Println("命令输入错误")
}
filename := filer[1]  //得到命令行数据
info,err := os.Stat(filename)
	if err != nil {
		fmt.Println("os.stat err :",err)
		return
	}
	fmt.Println(info)			 //文件属性集合结构体
	fmt.Println(info.Name())	 //文件的名字
	fmt.Println(info.Size())	 //文件的大小
	fmt.Println(info.Sys())		 //文件的底层数据来源
	fmt.Println(info.Mode())	 //文件的权限模式
	fmt.Println(info.ModTime())  //返回最后修改时间
	fmt.Println(info.IsDir())	 //是否是目录(文件夹)
}