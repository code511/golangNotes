package main

import "fmt"

/* 接口的继承 :
通过匿名字段的方式继承另一个接口		*/
type humaner interface {  // 子集
	sayhi()				//接口中包含一个 sayhi 方法
}
type personser interface {// 超集
	humaner				//以匿名字段的方式继承 humaner 接口
	sing(lrc string)	//自身方法
}
type students struct { //结构体 绑定了两个接口中的方法
	name string
	id	int
}
func (tmp *students)sayhi(){ //结构体方法1
fmt.Printf("students[%v,%v]sayhi\n",tmp.name,tmp.id)
}
func (tmp *students)sing(lrc string){ //结构体方法2
fmt.Println("students在唱着:",lrc)
}
var p personser //声明接口类型变量 p (超集)
func main(){
s := &students{"小明",8} // 初始化结构体类型 s
p = s	// s 赋值给 p
p.sayhi()  // 调用 继承的 方法
p.sing("进行曲") //调用 自身 方法
// 接口转换 :
//超集的变量可以传给子集 ,(不可反向)
var h humaner  //子集
p := &students{"小黑",11} //超集的变量
h = p	//转换给子集
h.sayhi() //使用
}