package main

import "fmt"

func main() {
	/*打印输出
	1 标准打印(一般用于字符串)*/
	print("goodbye java")
	println("hello go") //标准打印
	/* 2 格式化打印(应用占位符) */
	a := 1
	b := 2.2
	c := '字'
	d := "字符串"
	e := true
	/* 1) 打印值所属格式(用%T表示)*/
	fmt.Printf("%T,%T,%T,%T,%T\n", a, b, c, d, e) //打印值所属格式
	/* 2) 打印出量的值
	手动:
	%d = 整型
	%f = 浮点型
	%c = 字符型
	%s = 字符串型
	%t = 布尔型
	自动(字符型自动默认为十进制数字):
	%v = 所有默认类型(自动推导类型)
	%+v= 自动推导详细内容
	*/
	fmt.Printf("%d,%f,%c,%s,%t\n", a, b, c, d, e) //手动打印值
	fmt.Printf("%v,%v,%v,%v,%v\n", a, b, c, d, e) //自动打印值
}
