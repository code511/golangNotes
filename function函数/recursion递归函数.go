package main

/* 递归函数 (函数调用自身) :
调用流程 : 先调用,后返回
*/
func main() {
	descend := func1(100) //实现从100递减加到1,得出结果
	println("descend 100递减到1 =", descend)
	add := func2(1)
	println("add 1递加到100 =", add)
}
func func1(a int) int { //递减函数{func1(a int)-->func1(a-1)-->func1(a)-->func1(a-1)-->...-->func(a int)-->return}
	if a == 1 { //		a +		100			  100-1		   99		  99-1				 1	结束递减递归(返回参数)
		return 1
	}
	return a + func1(a-1)
}
func func2(a int) int { //递加函数
	if a == 100 {
		return 100
	}
	return a + func2(a+1)
}
