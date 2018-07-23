package main

import "fmt"

func main() {
	var (
		s2,s1,s			   string					  //接收者
		i                      int					  //...
		f                      float32				  //...
		input                  = "56.12 / 5212 / Go"  // 待读取
		format                 = "%f / %d / %s"		  // 格式化
	)

	fmt.Println("请输入您的全名: ")				// 打印字符串输出到终端
	fmt.Scanln(&s1, &s2)						//从终端读取到 s1 ,s2 (只读字符串,多个空格分隔)
	// fmt.Scanf("%s %s", &s, &ss)  			//从终端格式化读取到s1,s2(可读多种类型)
	fmt.Printf("Hi %s %s!\n", s1, s2)	   // 格式化打印输出到终端
	fmt.Sscanf(input, format, &f, &i, &s)		//读取字符串并按格式提取输出到终端变量
	fmt.Println("让我们看看接收到的变量: ", f, i, s)

}
