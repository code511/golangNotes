package main

import "fmt"

/* 匿名字段 :
在一个结构体内添加另一个结构体,普通,别名类型的元素类型
语法 :			*/
type person struct { //声明结构体 person(人)
	name string //名字
	age  int    //年龄
	sex  string //性别
}
type student struct {
	person         //把 person的元素类型添加进 student
	learnId int    //学号
	class   string //班级
	address string //住址
	int            //匿名普通类型
}

func main() {
	// * 初始化 :
	// 1) 顺序初始化 :
	var xm student = student{person{"xm", 10, "男"},
		11, "三班", "地球", 1}
	// 2) 自动推导 :
	xy := student{person{"xy", 9, "男"},
		11, "三班", "地球", 1}
	// 3) 指定成员初始化(没有初始化的成员,默认为0或nil) :
	xh := student{person: person{name: "xh"}, learnId: 777}
	// 打印 :
	fmt.Printf("xm%+v\nxy%+v\nxh%+v\n", xm, xy, xh)
	// * 成员的操作 :
	xh.name = "xhtx"                  //改
	xh.age = 11                       //增
	xh.person = person{"xh", 12, "女"} //对匿名字段操作
	fmt.Printf("xh%+v", xh)
	// * 同名字段(如不指定,采用就近原则)
	// * 指针操作 (使用* ,& , new)

}
