package main

import "fmt"

/* 切片	:
可以认为是 动态数组 ,不限定长度.作为参数为(指针传递)
一个切片包含 长度(等同数组长度) 和 容量(切片的最大长度)两个属性			*/
func main() {
	/* 创建:
		1) 标准语法(利用make函数):
	var s = make([]T ,len ,cap)   	:  []T(类型), len(长度) ,cap(容量)	*/
	var s []int
	s = make([]int, 5, 10) //	声明一个长度 5,容量 10 的 int类型切片
	s = []int{1, 2, 3, 4, 5}
	fmt.Printf("s的类型为%T,s的长度为%v,s的容量为%v\n", s, len(s), cap(s)) //打印 S的属性
	// 2) 简易语法 :
	s1 := make([]string, 3, 5)
	s1 = []string{"下标0", "下标1", "下标2"}
	// 3) 忽略容量 :
	s2 := []int{2, 5, 10}
	// 截取(从数组或切片生成新的切片)  注: 包头不包尾 :
	s3 := s1[1:] //从 s1 的 1下标 开始截取
	fmt.Println(s3)
	s4 := s2[:1] //从 s2 的 0下标 截取到 1下标
	fmt.Println(s4)
	s5 := s1[1:2] //从 s1 的 1下标 截取到 2下标 (包头不包尾)
	fmt.Println(s5)
	s6 := s1[:] //从 s1 的 开始 截取到 结束
	fmt.Println(s6)
	s7 := s2[1] //获得指定下标的值
	fmt.Println(s7)
	// 赋值 : 延伸出的切片的 新赋值操作会 改变原始切片的 元素值
	s8 := []int{888, 666, 888}
	s9 := s8[1:]
	s9[0] = 888     // s9[0] == s8[1]
	fmt.Println(s8) //改变了原始切片的对应值
	// 切片的生长（append 和 copy 函数）:
	// append :
	s9 = append(s9, 888, 888, 888) //往切片 s9的末尾添加元素,自动扩充容量( 2倍递增)
	fmt.Println("s9的长度为%v,容量为%v\n", len(s9), cap(s9))
	// copy :
	s11 := []int{1, 2, 3}
	s12 := []int{3, 3}
	copy(s11, s12)   //用 s11 的元素按下标 替换掉 s12的元素
	fmt.Println(s11) //打印结果 3,3,3
}
