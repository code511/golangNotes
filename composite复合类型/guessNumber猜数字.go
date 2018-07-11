package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 猜数字小游戏 : 手动输入匹配 随机给出四位数字,每次匹配时 给与每一位提示(大或小),直到全部正确
func main() {
	a := dice() //得到骰子
	for {
		s := input()
		nb := 0                  //定义一个计数器 : nb == 答对个数
		for i := 0; i < 4; i++ { //测试答案 (每次 i+1 只打印一个结果)
			if s[i] > a[i] {
				fmt.Printf("第%v位大了点,", i+1) //如果
			} else if s[i] < a[i] {
				fmt.Printf("第%v位小了点,", i+1) //否则,如果
			} else {
				fmt.Printf("第%v位猜对了,", i+1) //否则
				nb++
			}
		}
		if nb == 4 { //四位数全部答对,跳出循环,游戏结束
			println("恭喜过关")
			break
		}
	}
}
func getnum(z int) (s []int) { // int 转 切片 函数
	s = append(s, z/1000, z%1000/100, z%100/10, z%10) //取商   ...append 自动添加到切片
	return
}
func input() (s []int) {
	for {
		var a int
		for {
			println("请输入四位数字:")
			fmt.Scan(&a)              //键盘取值 赋值给 z
			if a > 999 && a < 10000 { //筛选四位数字
				break
			}
			println("输入错误")
		}
		s = getnum(a) //完成 键盘输入四位数 变成切片赋值给 s
		return
	}
}
func dice() (a []int) {
	rand.Seed(time.Now().UnixNano()) //创建时间活随机种子
	i := rand.Perm(10)
	a = append(a, i[0], i[1], i[2], i[3])
	return
}
