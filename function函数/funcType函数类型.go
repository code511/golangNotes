package main

/* 函数类型 :
函数也是一种数据类型,通过 type 给一个函数类型起名
函数的参数返回值类型和定义的函数类型对应时才可以给函数类型赋值		*/

type lx func(int, int) int //定义了一个叫 lx 的函数类型
// 关键字 命名 关键字     形参	  返回值 (参数要一一对应)
func main() {
	var jjj lx          //声明一个函数类型的变量 jjj
	jjj = jia           //向 jjj 赋值 jia 函数  ,赋予了jjj和jia同等的功能
	j := jjj(1, 2)      //使用 jjj 函数类型
	println("1+2 =", j) //打印使用 jjj 的结果
	var xxx lx = jian   //声明一个函数类型 xxx ,并赋值
	x := xxx(1, 2)      //使用 xxx
	println("1-2 =", x) //打印结果
}
func jia(a, b int) int { //相加函数
	return a + b
}
func jian(a, b int) int { //相减函数
	return a - b
}
