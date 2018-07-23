package main

import (
	"strconv"
	"fmt"
)

/* strconv字符串转换 : 				*/
func main(){
// 1) 转换为字符串后追加到字节切片 :
sl := make([]byte ,0 ,1024)		//声明一个切片 sl
	//追加一个布尔(bool)类型
sl = strconv.AppendBool(sl,true)
	//追加 666 ,以 10进制的方式
sl = strconv.AppendInt(sl,666,10)
	//追加一个字符串
sl = strconv.AppendQuote(sl,"字符串")
fmt.Println(string(sl))			//打印
// 2) 其他类型转换字符串 :
// bool类型转换字符串
st := strconv.FormatBool(false)
fmt.Println(st)			//float值   标记格式   精度   float32/64
// float浮点型转换为字符串
st = strconv.FormatFloat(3.14,'f',-1,64)
fmt.Println(st)
// int整型转字符串
st = strconv.Itoa(888)
fmt.Println(st)
/* 3) 字符串转其他类型 :
   因为字符串内容需要严格对应要转换的类型,所以多了一个error函数返回值,谨防错误
字符串转bool类型 :				*/
b,e := strconv.ParseBool("false") //返回两个值,bool(b)和error函数(e)
if e ==nil{			//如果没有错误,打印
	fmt.Println(b)
}else {  			//否则(有错误),打印错误
	fmt.Println("e=",e)
}					// 下面示例,忽略error函数
//将字符串转换为int整型(1) :  字符串    进制	  int/8/16/32/64
i,_ := strconv.ParseInt("555",10,16) //忽略了error
fmt.Println(i)
//将字符串转换为int整型(2) :
ii,_ := strconv.Atoi("666") //返回 int ,error() 在此忽略
fmt.Println(ii)
//将字符串转换了float浮点型 :
f,_ := strconv.ParseFloat("3.14",16)// string , float大小
fmt.Println(f)
}