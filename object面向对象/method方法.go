package main

import "fmt"

/*
方法: 带有绑定类型(自定义类型)的函数
	 原始自定义类型(不能是指针或接口)作为第一实参
方法集: 一种类型能够调用的多种方法
	语法: func (name type)func name(parameters)(results){}
		 func (变量名 别名类型(实参)) 方法名 (形参) (返回值) {方法体}
	* 普通类型 :			*/
type i int //创建普通类型别名 i
// 1) 相加 :
func (a1 i) add(a2 i) i { //创建绑定 i类型 方法
	return a1 + a2
}

// 2) 打印显示 :
func (a i) show() { // i类型 show
	fmt.Printf("a=%v\n", a)
}

//  * 结构体类型 :
type persons struct {
	name string
	age  int
}

// 1) 调用指针,修改 :
func (x *persons) setInfo(n string, a int) {
	x.name = n
	x.age = a
}

// 2) 打印显示 :
func (x persons) show() { // persons类型 show
	fmt.Printf("x=%+v\n", x)
}

func main() {
	// 普通类型调用 :
	var a i = 1                        //初始化一个 i 类型变量 a (类型名必须严格一致)
	a = a.add(1)                       // b 接收 方法返回值(只要类型(i)一致就可以调用 add方法)
	b := a.add(2)                      // 新变量 自动推导为 i 类型
	fmt.Printf("b的类型为%T,值为%v\n", b, b) //打印结果
	// 结构体类型指针调用(自动转换) :
	var s1 persons         //定义一个persons类型 s1
	(&s1).setInfo("小明", 8) //使用 setInfo 方法对 s1的指针值 赋值修改
	// * 同名不同类型的方法使用 :
	// 		只要接受类型不一样,方法名可以重复,无冲突使用
	a.show()  //使用i类型 show方法打印显示 a
	s1.show() //使用persons类型 show方法打印显示 s1
}
