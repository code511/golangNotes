package main

import (
	"math/rand"
	"fmt"
)

/* 冒泡排序 :
对数组内的元素进行两两比较交换,最终进行某种排序
 */
func main(){
	var a [5]int
	i := len(a)
	for i=0;i<5 ;i++  {
		a[i] =rand.Intn(9)}// 将小于10 的随机数存入a的 10个下标
	fmt.Println(a)//打印产生的随机数组
	for j:=0;j<len(a)-1 ; j++{ 			//元素需要比较 len(a)-1 次
		for i = 0; i < len(a)-1; i++ { // 元素需要交换 len(a)-1 次
			if a[i] > a[i+1] {			 //通过> ,<比较相邻元素来决定 正反排序
				a[i+1], a[i]=a[i], a[i+1]//交换大小元素位置
			}
		}
	}
	fmt.Println(a)//打印结果 (从小到大排序)

	}