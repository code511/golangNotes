package main
/* GO语言有两个系统函数(系统自动执行,无参无返,不能被调用)
	main函数 : 只能用于 package main (main 包),一个项目必须且只能有一个
	init函数 :	能用于所有 package (包),一个项目可以有多个
				init函数在main函数之前执行,如果 package main中导入的包有init函数,则先执行包中的init
	多个init函数 : 先执行导包中的init函数,再按顺序执行
			注 :  init函数一般用于程序执行前的引导初始化工作 !
*/
func init()  { // 先执行
	println("程序引导")
	}
func init() {  // 后执行
	println("程序引导中")
}
func main(){   // 再执行
	println(" 1\n 2\n 3\n程序结束")
}