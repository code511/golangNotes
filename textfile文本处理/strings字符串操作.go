package main

import (
	"fmt"
	"strings"
	"sort"
)

// 字符串操作 :
func main(){
// 1) contains 检查是否包含(包含true,不包含false) :
fmt.Println(strings.Contains("hellogo","go"))
// 2) Joins 提取组合 :
m := []string{"字","符","串","切","片"}
s := strings.Join(m,"") //可以为用任何字符串连接
fmt.Println(s)
// 3) Index 查找包含字符串的下标位置(不包含返回 -1) :
fmt.Println(strings.Index("hellogo","go")) //  5 下标开始
fmt.Println(strings.Index("hellogo","hi")) // -1 (不包含)
// 4) Repeat 重复拼接指定数量字符串 :
ss := strings.Repeat("go",3)
fmt.Println(ss)
// 5) Split 以指定的分隔符拆分(返回切片) :
sss := strings.Split(ss,"o") //以 "o" 拆分
fmt.Printf("sss=%v\n",sss)
// 6) Trim 去掉两头指定字符 :
a := "...good morning." //两头随机数 .
a1 := strings.Trim(a,".") //去掉两头 .
fmt.Println(a1)
// 7) Fields 去掉所有空格,并以空格为分隔符把元素分别放入切片中 :
b := "Whatever is worth doing is worth doing well"
b1 := strings.Fields(b) // 用切片 b1 接收
for i,d := range b1{ //迭代打印
	println(i,"=",d)
}
// 8) Replace 替换指定内容
	str := "welcome to bai\ndu\n.com"
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(str)
// 9)  sort 排序 (按照ascii码表)
	stri := []string {"1","3","2"}
	sort.Strings(stri)
	fmt.Println(stri) // 1 2 3
}