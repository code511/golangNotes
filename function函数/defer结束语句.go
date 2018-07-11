package main

/* defer 关键字用于函数或方法的延迟执行(defer语句只能出现在函数或方法内部)
defer语句在main函数结束前调用
defer经常被用于成对的操作,如:开关锁,开关文件,释放资源类的操作
通过defer机制,资源一定会被释放
多个defer语句从main函数开始倒序执行(LIFO先进先出)
*/
func main() {
	defer println("程序结束") //最后执行的的defer语句
	a, b, c := 1, 2, 3
	defer println(a)
	defer println(b) //有传参,先传递参数,后执行
	defer println(c)
	//与匿名函数结合使用
	defer func(a, b, c int) { //后调用
		println(a)
		println(b) //打印提前传递进去的参数
		println(c)
	}(a, b, c) //有传参,会先把1,2,3参数传递进去
	a, b, c = 5, 2, 0
	println(a, b, c)
}
