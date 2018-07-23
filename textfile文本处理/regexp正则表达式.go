package main

import (
	"regexp"
	"fmt"
)

/* 正则表达式 :
按照规则解析提取字符串文本
提取步骤 :
1) 导入regexp包
2) 创建解析器(规则) 用`反引号`
3) 使用解析器
*/
func main(){
s := "abc afc a2c a5c aab b77c"  //字符串样本
s1 := "fds324fgdsa4354323fds34543"
zz := regexp.MustCompile(`a(\d+)`)//创建解析器 ()代表分组,以方便过滤
if zz == nil{
	fmt.Println("zz regexp err ")//zz为空代表解析失败,返回错误提示
	return
}
jg := zz.FindAllStringSubmatch(s,-1)		//使用解析器 zz分段 提取字符串 s
fmt.Println("jg=",jg)
jg1 := zz.FindAllString(s1,-1)	//使用解析器 zz提取字符串 s1
fmt.Println("jg1=",jg1)
for _,v := range jg{	//过滤  a
	fmt.Println(v[1])
}
}