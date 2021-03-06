package main

import "fmt"

/* 指针类型格式	:
*数据类型 保存 数据类型 的指针,**数据类型 保存 *数据类型 的指针,以此类推
指针默认值为 : nil 或 0
			 & 取目标指针, * 取指针的 值
			 . 访问目标成员
			 可以作为函数形参
注 :
指针类型 带有 该地址 的值,可以操作 该地址 的值(不支持指针运算)
		不要操作没有合法指向(nil)的内存!
	使用格式 :				*/
func main() {
	a := 666
	b := &a //自动推导指针类型
	/*手动方式 :
	var b *int   已知a为int类型,所以声明*int指针类型
	b = &a		 把 a的地址赋值给 b			*/
	*b = 888                               //指针类型 *b 进行赋值操作
	fmt.Printf("a的指针为%v,b的指针为%v\n", &a, b) //打印a,b的指针
	fmt.Printf("a的值为%v,b的值为%v\n", a, *b)   //打印a,b的 值
	var c *int
	fmt.Printf("c的指针为%v\n", c) //打印指针默认值
}
