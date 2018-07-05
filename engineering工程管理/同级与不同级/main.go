package main //程序必须有一个 main包
/*	第一步 :
	创建 GO PATH 环境变量到 : 存放 src 文件夹的路径, 代表这是一个工作空间(用来存放自己和别人的库,系统通过GO PATH调用)
*/
import "./test02" //不同级目录的导入包 调用函数

func main() {
	test()               //同一个目录可以直接调用其他文件的函数
	a := testt.Add(6, 8) //调用test02文件夹内的tsett包中的Add函数
	println(a)
}
