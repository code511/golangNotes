package main

/* 导包方式 :
 1) 常规 :
import (
	"fmt"	  //在()内导入多个包
	"./test"	//可以是路径
)*/
// 2)	. 操作(调用时省略包名) :
import . "fmt"

/* 3)	_ 操作(调用包里的其他函数,如: init函数)
import _ "fmt"
   4) 给包起一个别名 :
import io "fmt"
*/
func main() {
	a, b, c := 1, 2, 3
	Printf("a=%d,b=%d,c=%d\n", a, b, c) //省略 fmt.使用printf进行格式化打印
	// io.Printf("a=%d,b=%d,c=%d\n",a,b,c)   使用别名代替fmt调用printf进行打印(效果与上相同)
}
