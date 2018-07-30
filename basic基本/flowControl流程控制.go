package main

import (
	"fmt"
	"time"
)

func main() {
	/*流程控制
	go语言支持最基本的三种程序运行结构 : 顺序结构 , 选择结构 , 循环结构 .
		1)	顺序结构 (程序按顺序执行,不发生跳转)
		2)	选择结构
		  2.1) if语句 :
			if		如果   	(每个if都会执行判断)
			else if	否则如果	(已经有语句为true则不执行)
			else	否则		(都不满足则执行else)
			支持初始化单语句
			{执行体}
	*/
	if a := 10; a == 10 { //if 初始化a ; 条件 {执行体} , (以下所有条件全部为true)
		println("a = 10") //执行(每个if语句只要为true即执行)
	}
	if a := 10; a >= 10 {
		println("a >= 10") //执行
	} else if a := 10; a <= 10 { //因为上面有条件为true,所以下面都不执行(多选用if,if,if;单选用elseif,elseif)
		println("a <= 10") //不执行(只要以上一个为true,则else if语句不执行)
	} else if a := 10; a != 11 {
		println("a != 11") //不执行
	} else {
		println("以上条件都不满足") //不执行,以上条件都不满足时else才执行
	} /*
		2.2) switch语句 :
		简洁,可读性强
		支持一个初始化语句
		{判断体}
		默认执行后跳出
	*/
	var num int
	println("请按下楼层:")
	fmt.Scan(&num) //键盘监听赋值给num
	switch num {
	case 1: //判断num是否为1(num == 1可以省略为1)
		println("按下了1层")
		//break (每一个条件体结尾都默认有一个break,执行即跳出switch语句)
	case 2:
		println("按下了2层")
		fallthrough //条件体结尾加上fallthrough,条件为true则余下无条件执行
	case 3:
		println("按下了3层")
	default: //都不满足执行
		println("不坐电梯")
	} /* 其他特性 :
	 */
	switch score := 88; { //初始化语句中条件可以省略
	case score > 90:
		println("优秀")
	case score > 70, score > 80: //可以多条件判断,逗号隔开
		println("良好")
	case score > 60:
		println("一般")
	default:
		println("不及格")
	}
	/*
		3) 循环结构 (for循环) :
			3.1)	标准格式 :
		for 初始化条件 ; 判断条件 ; 条件变化 { 执行体 } */
	var a2 int                    //声明for语句体外变量a2
	for a1 := 1; a1 <= 10; a1++ { //	for 初始化条件 ; 判断条件 ; 条件变化 { 执行体 } */
		a2 += a1           //执行体
		println("a2=", a2) //打印相加过程
	}
	println("最后a2=", a2) //打印最后结果
	/*
		 	3.2)	迭代 (range) :
				获取每一个角标的元素
		 	a)	传统方法 :      */
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	} /*
		3.2.1)	range方法 :
			迭代打印每个元素,返回两个值:值的角标,和角标的元素		*/
	for i, data := range str { // for(关键字) i(角标) ,data(数据) := range(关键字) str(要遍历的量)
		fmt.Printf("str[%d] = %c\n", i, data) // data可以省略,或用_代替 (保留角标,元素丢弃)
	} //结果同上
	/*	3.3)	死循环 :				*/
	b1 := 0
	for { //for后面不加条件,代表条件永远为真(死循环)
		b1++
		time.Sleep(time.Second) //利用time函数延迟1秒
		/*	3.3.1)	continue和break控制方式(用在loop,switch,select) :
			continue : 跳过单次 (含本次)
			break	 : 跳出本次循环 (含本次) , 如果嵌套多个循环,则跳出本次内循环
		*/
		if b1 == 5 {
			continue //如果b1 == 5 则跳过本次循环,等待下次循环
		}
		if b1 == 10 {
			break //如果b1 == 10 则跳出本次循环
		}
		println("b1 = ", b1) //无break时,会打印出无限++的b1
	}
	/*	4)	goto语句跳转	(用在一个函数内的任何地方)	*/
	println("第一个") //执行第一个
	goto perform   //增加跳转,标签为perform(直接跳转到 perform: 处执行)
	println("第二个") //第二个被跳转
perform: //使用了perform标签
	println("第三个") //执行第三个
}
