package main

import "fmt"

/* 空接口 :
空接口(interface{})不包含任何的方法,所以所有类型都实现了空接口,可以粗存任意类型的数值
空接口作为函数形参时,表示可以接收任意类型数量的实参,如 ...interface{} (万能类型)
 */
func main(){
// * 普通使用 :
	var v interface{} //声明 空接口(万能类型) v
	v = 123
	v = "贝多芬"		//可以储存任何类型
	fmt.Println(v)  //打印
// * 作为形参 :
kjk("字符串",'字',123,3.14)  //调用传实参
}
func kjk(i...interface{})  {  //空接口形参 函数
	fmt.Println(i)
}