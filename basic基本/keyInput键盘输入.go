package mian

import "fmt"

func main() {
	/*键盘输入(使用内建函数fmt中的scan键盘监听系列方法)*/
	var a int          //声明一个int变量
	print("请输入变量a的值:") //打印输入提示
	fmt.Scan(&a)       //等待键入 ,键入后赋值给a
	print("a=", a)     //打印结果
}
