package main

import "fmt"

/* 类型断言 :
判断 ...interface{}中元素的类型
 */
type t struct {
	name string
	age  int
}
 func main(){
 itf := make([]interface{},3) 	 //声明 空接口类型 切片
 itf[0] = t{"名字",18}// 初始化... 结构体 t
 itf[1] = "字符串"				// 			string s
 itf[2] = 555					// 			i	int
// 1) if 类型断言 :
	 for i,d := range itf {		//在迭代中进行判断
		 if c, ok := d.(int);ok ==true {  //语法
			fmt.Printf("itf[%v]的类型为int,值为%v\n",i,c)
		 }else if c, ok := d.(string);ok ==true {
			 fmt.Printf("itf[%v]的类型为string,值为%v\n",i,c)
		 }else if c, ok := d.(t);ok ==true {
			 fmt.Printf("itf[%v]的类型为t结构体,值为name=%v,age=%v\n",i,c.name,c.age)
		 }
	 }
// 2) switch 类型断言 :
	 for i,d := range itf {		//在迭代中进行判断
		switch c := d.(type){	//语法 c 对比case
		case int :
			fmt.Printf("itf[%v]的类型为int,值为%v\n",i,c)
		case string :
			fmt.Printf("itf[%v]的类型为string,值为%v\n",i,c)
		case t :
			fmt.Printf("itf[%v]的类型为t结构体,值为name=%v,age=%v\n",i,c.name,c.age)
		}
	 }
}