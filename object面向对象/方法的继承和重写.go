package main

import (
	"fmt"
)

/* 方法的继承 :
一个方法继承它的 接收类型中 包含的方法
*/
type ps struct {
	name string
}

func (p ps) show() {
	fmt.Printf("p.name=%v\n", p.name)
}

type stdt struct {
	ps  // stdt继承了 ps的show方法
	age int
}

func (s stdt) show() { //stdt 重写了一个show方法
	fmt.Printf("s.age=%v\n", s.age)
}
func main() {
	/* 调用:
	   默认自动(使用就近原则)调用 :   */
	p := ps{"小明"}
	p.show() //优先使用自身方法
	s := stdt{ps{"小黑"}, 15}
	s.show() //优先使用自身方法
	// 指定调用 :
	s.ps.show() // s调用 ps中的 show方法

}
