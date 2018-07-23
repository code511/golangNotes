package main

import "fmt"

/* recover函数 :
 为了不使程序中断运行,可以使用recover函数跳过并返回panic异常,继续往下执行
注: recover只能在defer语句执行,默认recover == nil			*/
func main(){
test1()
test2(15) //数组越界异常触发
test3()
}
func test1()  {
println("1")
}
func test2(i int)  {
	defer func() {  //defer内建函数
		if rec :=recover(); rec != nil{ //如果出现异常 ,则返回异常信息,程序继续执行
			fmt.Println(rec)
		}
	}()
	var a [10]int
	a[i] = 444
	fmt.Println(a[i])
}
func test3(){
	println("3")
}