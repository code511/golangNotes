package main

func main() {
	/*类型转换
	1) 只能转换兼容类型
	2) 格式:     */
	s := '字'              // 创建一个rune型变量
	i := int(s)           // 转换为int类型并赋值给i,'字' = 23383
	str := "string"       //创建一个string型变量
	int := int(str[3])    //将str第三个角标的值转换为int并赋值给int
	print(s, i, str, int) //打印结果
}
