/*
1) go语言以包作为管理单
2) 每个文件必须先声明包
3) 主程序(可执行程序)必须要有一个main包!
*/
package mian //main包固定格式

import "fmt" //通过import导入非main包

func main() { //通过func声明函数(方法),左大括号必须和函数名同一行

	fmt.Println("hello,word!") //标准打印输出(自动换行)
	/*
	   1) //为行注释,/*为块注释,同java
	   2) go语言语句结尾没有分号(;),不同于java
	*/
}
