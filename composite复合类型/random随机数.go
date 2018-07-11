package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1) 伪随机数 :
//就是使用math包中的rand函数
func main() {
	//rand.Seed(1)  //(死种子)以 1为种子 ,导致每次重复
	rand.Seed(time.Now().UnixNano()) //(活种子)以当前本地时间(最小单位 纳秒)作为种子 ,保证随机数的唯一性
	for i := 0; i < 10; i++ {
		//println(rand.Int())  //打印最大位数
		//println(rand.Intn(9999))  //打印9999以内的随机数
		fmt.Println(fmt.Sprintf("%04v", rand.Int31n(9999))) //打印固定四位数的随机数
	}
	println("--------------------------------")
	s := rand.Perm(4) //生成 n个(n 大小范围内)的 int切片
	fmt.Println(s)
	/* 2)真随机数 :
	crypto/rand包中的函数    */

}
