package main

import "fmt"

/*	匿名函数(闭包) :
匿名函数是定义在函数内部的没有名字的函数,闭包通过匿名函数来实现,它可以调用并改变同一作用域的量和量值
匿名函数也是一种函数类型!
 */
func main(){
	a := 10
	b := 20
	println(a,b)	//打印结果10,20
	//定义匿名函数,直接调用
	// 无参无返 :
	func(){
		a =66	//匿名函数改变了外部的变量
		b =88
	}()
	println(a,b)	//打印结果66,88
	//有参有返回值 :
	min,max := func(a,b int)(min,max int) { //有返回值需要有变量接收
		if a<b {
			min =a
			max =b
		}else {	min =b
			max =a}
			return
	}(a,b)			//直接输入参数并调用
	fmt.Printf("min=%d,max=%d\n",min,max)//打印结果
	//可以把匿名函数赋值给一个变量,或者声明一个匿名函数类型,通过赋值的方式调用
	f1 := func() {	//把匿名函数自动推导赋值给 f1
		println("匿名函数")
	}
	f1()			//调用 f1
	type nmhs func()//声明一个匿名函数类型
	var f2 nmhs		//声明一个匿名函数类型变量 f2
	f2 =f1			//把 f1 赋值给 f2
	f2()			//调用 f2 =调用 f1


}