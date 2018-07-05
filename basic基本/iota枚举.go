package mian

import "fmt"

func main() {
	/*
		iota枚举
		iota在const关键字出现时将被重置为0(const内部的第一行之前)，
		const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
			1) 枚举只能在常量的表达式中使用 (如果同一行,值相同) */
	const (
		a1     = iota       //0
		a2, a3 = iota, iota //1,1,1
		a4     = iota       //2
	)
	/*
		2) 每次const出现,都会让iota初始为零  */
	const b1 = iota //0
	const (         //0
		b2 = iota //1
		b3        //2
	)
	/*
		3)	可用下划线_或插队的方式跳过某个值
	*/const (
		c1 = iota //0
		_         //跳过了1
		c2 = 3.14 //赋值3.14,并跳过2
		c3        //3
	)
	fmt.Printf("a1=%d,a2=%d,a3=%d,a4=%d,b1=%d,b2=%d,b3=%d,c1=%d,c2=%f,c3=%f\n", a1, a2, a3, a4, b1, b2, b3, c1, c2, c3)
}
