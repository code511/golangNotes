package main

import (
	"encoding/json"
	"fmt"
)

/*
JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。它使得人们很容易的进行阅读和编写。
同时也方便了机器进行解析和生成。 JSON采用完全独立于程序语言的文本格式.
JSON基于两种结构：
JSON 有两种结构[2]
json简单说就是javascript中的对象和数组，所以这两种结构就是对象和数组两种结构，通过这两种结构可以表示
各种复杂的结构
1、对象：对象在js中表示为“{}”括起来的内容，数据结构为 {key：value,key：value,...}的键值对的结构，
在面向对象的语言中，key为对象的属性，value为对应的属性值，所以很容易理解，取值方法为 对象.key 获取属
性值，这个属性值的类型可以是 数字、字符串、数组、对象几种。
2、数组：数组在js中是中括号“[]”括起来的内容，数据结构为 ["java","javascript","vb",...]，取值方式
和所有语言中一样，使用索引获取，字段值的类型可以是 数字、字符串、数组、对象几种。
 */
 // 1) json编码(生成) :
 //   通过 结构体或 map(结构体成员变量名首字母必须大写) :
type IT struct { 		//定义结构体类型 IT		二次编码(可选) :
	Company string	  `json:"_"`  		 // 用 _ 过滤
	Subjects []string `json:"subjects"`	 //  改写
	Isok bool		  `json:",string"`	 //  转格式为string输出
	Price float64	  `json:",string"`	 //  转格式
}
func main(){
	//声明初始化map键值对 ss
	ss := (make(map[string]interface{},5))
	ss["company"] = "itcast"
	ss["subjects"] = []string{"go","c++","python","java"}
	ss["isok"] = "bool"
	ss["price"] = "666.666"
	jsons,er := json.MarshalIndent(ss,"","	")	// 将 ss 格式化编码(生成)为 json文本
	if er != nil{
		fmt.Println("err =",er)		//有错误则返回错误
		return
	}
	fmt.Println(string(jsons)) 		//打印输出结果 (map为无序打印!)
// 声明并初始化结构体变量 s
s := IT{"itcast",[]string{"go","c++","python","java"},true,666.666}
jsonstr,err := json.MarshalIndent(s,"","	")	// 将 s 格式化编码(生成)为 json文本
if err != nil{
	fmt.Println("err =",err)		//有错误则返回错误
	return
}
fmt.Println(string(jsonstr))  //打印输出结果
// 2) 解码(解析) json 到 结构体或 map :
// 到 结构体 :
var vv IT //声明一个结构体变量 vv
err1 := json.Unmarshal([]byte(jsonstr),&vv) //解码 并返回到 vv指针
if err1 != nil {
	fmt.Println("err1 =",err1)
	return
}
fmt.Printf("vv=%+v\n",vv) //输出结果
// 到 map :
vv2 := make(map[string]interface{},5)
	err2 := json.Unmarshal([]byte(jsonstr),&vv2)//解码 并返回到 vv2指针
	if err1 != nil {
		fmt.Println("err2 =",err2)
		return
	}
	fmt.Printf("vv2=%+v\n",vv2) //输出结果
// 3) 提取内容 :
//map 需要使用 类型断言 来提取 内容
}