package main
/* error接口(常规异常处理) :
如果程序某个部分可能发生异常,则需要设置异常处理说明			*/
import (
	"errors"  //需要 导入 errors包
	"fmt"
)
 func divide(a,b int)(result int,err error){ //增加 err 类型返回值(默认err==nil)
 	if b == 0 {			//错误条件
 		err = errors.New("分母不能为0")// 错误条件达成,则给err初始化 异常说明
	}else {
		result = a/b	//如果无错误,则正常返回
	}
	return
 }
func main(){
result,err :=divide(10,0) //自动推导接收返回值
if err != nil {	//如果err不为 nil,则说明有异常,打印异常说明
	fmt.Println("err=",err)
}else{			//否则正常打印结果
	fmt.Println("result=",result)
}
}