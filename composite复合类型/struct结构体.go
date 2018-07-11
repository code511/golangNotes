package main

import "fmt"

/* 结构体 :
结构体是 多种类型数据的集合
传递为值传递,全程支持指针操作
可以定义在函数体外:		*/
type student struct { //定义了一个学生属性的 结构体类型
	name  string //名字
	age   int    //年龄
	class string //班级
}

func main() {
	// * 初始化 :
	// 1) 顺序初始化(每个成员都必须初始化) : //声明一个student结构体类型的变量wuyichen
	var wuyichen student = student{"吴奕辰", 3, "海棠花班"}
	fmt.Println(wuyichen)
	// 2) 指定成员初始化(没有初始化的成员默认为0或nil) :
	wuzongze := student{name: "吴宗泽", age: 6} //自动推导,初始化两个
	fmt.Println(wuzongze)
	// * 操作成员(使用 . 运算符) :
	wuzongze.age = 7       //改写
	wuzongze.class = "不知道" //赋值
	fmt.Println(wuzongze)
	// * 指针操作(支持 new) :
	//在类型前加 &(取地址) 运算符
	// 1) 指针 :
	xxs := &student{"1", 2, "3"}
	xxs.age = 10
	// 3) new :
	zxs := new(student)
	zxs.age = 15
	fmt.Println(xxs, zxs)
	// *结构体的比较(可以使用==和!=,当==时) :
	fmt.Println(wuzongze == wuyichen) //false
	// *结构体之间赋值(当结构体类型相同时) :
	var xbb student
	xbb = wuyichen // 将 wuyichen的参数赋值给 xbb
	fmt.Println(xbb)
}
