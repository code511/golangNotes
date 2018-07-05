package main

import (
	"fmt"
	"math"
)

/*函数(也叫方法,功能) :
规则 :
1)	函数可以声明在 main主函数 上面 或者 下面
	一个函数内的量不能重名
2)	组成 : 	func  函数名 (形参) (返回值) {函数体}
	2.1)	声明 :	函数由关键字 func 开始声明
	2.2)	命名 :	函数名首字母小写为 private(私有) ,首字母大写为 public(公共)
	2.3)	形参 :	接收数据的参数
		2.3.1) 只可以单向接收参数
		2.3.2) 可以无形参,或有多个形参(通过,分隔)
	2.4)	返回值:	返回传入形参后经过函数体运算的结果
		2.4.1) 可以无返回值,或有多个返回值(通过,分隔)
		2.4.2) 无返回值或单返回值时,()括号可以省略
		2.4.3) 有返回值时,函数体内必须有 return返回语句
		2.4.4) 可以只有类型,没有变量名
3)	参数类型格式 (可以是函数类型) :
	3.1) 多个同类型参数可以只在一段参数后加类型		: (a , b int)
	3.2) 多个不同类型参数		: (a , b int , c , d float , e rune)
	3.3)	不定参数 (不固定参数数量 ,可以是 0 或 n 个参数)
		3.3.1)	只能用在形参
		3.3.2)	格式 :  ...type(数据类型)
		3.3.3)	只能是参数中的最后一个类型
		3.3.4)	传递 : 	函数 ( 参数名[:2]...)	   : 传递不定参数 2 角标之前的参数	(不包尾)
						函数 ( 参数名[2:]...)    : 传递不定参数 2 角标之后的参数	(包头)
示例 :
1)	无参无返回值 :			*/
func funcOne() {
	println("无参无返回值 funcOne : 示例") //打印示例
}

// 2) 有参无返回值 :
func funcTow(a, b int) {
	c := a + b
	println("多参无返回值 funTow : 实参相加 =", c) //做加法,打印结果
}

// 3) 有不定参数无返回值 :			接收不定参传递的函数
func funcThree1(a ...int) {
	var b int
	for _, data := range a {
		b += data
	}
	println("不定参数无返回值 funcThree : 不定参相加 =", b)
}

// 3.1)	不定参数传递 :				传递不定参的函数
func funcThree(a ...int) {
	funcThree1(a[:3]...) //向 funcThree1传递了接收到的数据a前3个角标的值(a[0],a[1],a[2])
}

// 4) 无参数有单返回值 :  可以省略量名和(括号)
func funcFour() int {
	return 666
}

/* 完整写法(推荐) :		给返回值一个命名
func funcFour()(a int){
	a = 666
	return
}	*/
// 5) 无参数多返回值 :
func funcFive() (a, b, c int) { //返回值不命名也可用( int, int, int )方式
	a, b, c = 1, 2, 3
	return
}

// 6) 多参多返回值 :
func funcSix(a, b float64) (min, max float64) {
	min = math.Min(a, b)
	max = math.Max(a, b)
	return
}

// 函数调用 :
func main() {
	funcOne()                //无参无返函数调用
	funcTow(1, 2)            //多参无返函数调用
	funcThree(1, 3, 5, 7, 8) //不定参无返函数调用
	a := funcFour()          //无参单返回值函数调用
	println("无参数有返回值 funcFour : a =", a)
	b, c, d := funcFive() //无参多返回值函数调用
	fmt.Printf("无参数多返回值 funcFive : b=%d,c=%d,d=%d\n", b, c, d)
	_, max := funcSix(66, 99) //多参多返函数调用 (用匿名函数跳过不想要的返回值)
	fmt.Printf("多参多返回值 funcSix : 最大值为%v\n", max)

}
