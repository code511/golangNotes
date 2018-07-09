package main
import (
	"math/rand"
	"time"
	"fmt"
)
// 猜数字小游戏 : 随机给出四位数字,回答时 给与每一位提示(大或小),直到全部答对
func main(){
	rand.Seed(time.Now().UnixNano())//创建时间活随机种子
	var a []int	//声明一个长度 5的 int数组
	for i:=0;i<4 ;i++  {//通过循环以此把 随机数放入 下标
	r := rand.Intn(10)// r == 10以内的随机数
	a = append(a,r)		}		//往 a下标赋值
	for{
		var s []int
		var z int
		for {
			println("请输入四位数字:")
			fmt.Scan(&z)	//键盘取值 赋值给 z
			if z>999 && z<10000 {
				break
			}
			println("输入错误")
		}
		s=getnum(z)		//完成 键盘输入四位数 变成切片赋值给 s
		fmt.Println(s)	//打印 手动输入结果
		nb := 0		//定义一个计数器 : nb == 答对个数
		for i := 0; i < 4; i++ {	//测试答案 (每次 i+1 只打印一个结果)
			if s[i]>a[i] {
				fmt.Printf("第%v位大了点,",i+1) //如果
			}else if s[i]<a[i] {
				fmt.Printf("第%v位小了点,",i+1) //否则,如果
			}else{
				fmt.Printf("第%v位猜对了,",i+1) //否则
				nb++
			}
		}
		if nb == 4 {	//如果全部答对,跳出循环,游戏结束
			println("恭喜过关")
			break
			}
	}
}
func getnum(z int)(s []int){  // int 转 切片 类型
	s= append(s,z/1000)		//取商
	s= append(s,z%1000/100)//取余再取商
	s= append(s,z%100/10)	//取余再取商
	s= append(s,z%10)	//取余
	return
}
func smelled(){

}
