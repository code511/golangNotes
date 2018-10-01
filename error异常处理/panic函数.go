package main
/* panic函数(致命异常处理) :
当遇到不可恢复的错误状态(数组越界,空指针)时,内部自动调用panic函数(空接口形参)报告致命错误
* 一般而言,当panic异常发生时,程序会中断运行,并立即执行goroutina(线程),随后程序崩溃并输出日志信息
日志信息包括panic和函数调用的堆栈跟踪信息.
* 不是所有的panic异常都来自运行时,直接调用内置的panic函数也会引发panic异常
 */
func main(){
testa()
testb()
testc()
}
func testa(){
	println("1")
}
func testb(){
	panic("this is panic pathTest")  //程序中断,并提示错误信息
}
func testc(){
	println("3")
}