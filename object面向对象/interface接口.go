package main

import "fmt"

/* 接口类型(多态) :
接口interface类型是一种自定义抽象类型,描述了一系列方法的集合
接口只能展示自己的方法,不能自我实例化
接口可以嵌入其他接口,或结构中
 * 接口定义 : 名字结尾加上 er 增加辨识度				*/
 type ducker interface { //定义语法
 	property()	// 先定义方法,再另外实现
 }
/* * 接口的实现 :
接口定义的行为不由接口直接实现,而是由(自定义类型)实现,实现了这些方法行为的类型为此接口的实例 		*/
type name struct {
	x string			// 声明自定义结构体类型...
	m string
}
type age struct {
	int
}
type color struct {
	string
}
func (n *name)property()  {		//为自定义类型添加进方法...
	fmt.Printf("它的名字叫%s\n",*n)
}
func (a *age) property()  {
	fmt.Printf("今年%v岁了\n",*a)
}
func (c *color)property()  {
	fmt.Printf("是一只%v色的小鸭子\n",*c)
}
var d ducker // 定义接口类型全局变量 d
func main(){

n := &name{"丑","小鸭"}
d = n
d.property()
a := &age{1}
d = a
d.property()
c := &color{"灰"}
d = c
d.property()
// 接口放入切片调用 :
x := make([]ducker ,3)  // 创建 接口类型切片
x[0] = n	// 赋值...
x[1] = a
x[2] = c
	for _,i := range x {  // 迭代调用
		i.property()
	}
// 接口作为函数形参(多态) :
	duck(n)   //传入 接口内方法中类型变量 调用...
	duck(a)
	duck(c)
}
func duck(d ducker)  {	// 接口作为函数形参
	d.property()
}