package main

import "fmt"

/*	new函数	:
 表达式 new(T)将利用指针类型 创建一个 T类型的匿名变量
	利用new(T) 为指针类型分配一个新的合法空间
	只需使用new()函数,不需清理内存,GO会自动清理
*/
func main() {
	a := new(int)                         // 为 a申请一个新的合法指针,并自动推导为指针类型
	*a = 666                              // 为 a的指针赋值
	fmt.Printf("a的指针为%v,a的值为%v\n", a, *a) //打印a的指针和值
}
