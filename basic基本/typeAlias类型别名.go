package main

import "fmt"

func main() {
	/*类型别名(给一个数据类型起一个别名)
	格式(使用type关键字):		*/
	type char rune //把rune类型增加一个名字char
	var c char = '字'
	type ( //多行批量添加别名操作
		bigint  int
		gofloat float64
	)
	var (
		i bigint  = 11
		f gofloat = 3.14
	)
	fmt.Printf("c=%c,i=%v,f=%v", c, i, f)
}
