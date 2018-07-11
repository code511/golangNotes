package main

type egg int
func (i egg)eggs(){
	println(i,"个鸡蛋")
}
func main(){
e := egg(5) //初始化 egg类型变量 e
e.eggs()	//调用 eggs方法
// 1) (把一套数据和方法封装为)方法值 :
// 对变量和方法进行操作
theEggs := e.eggs	//创建方法值 theEggs
theEggs()			//调用
// 2) 方法表达式 :
// 对类型和方法进行操作(如有指针,则跟随)
te := egg.eggs  //封装的方法表达式
te(e) //显式传递(显示变量名)
}