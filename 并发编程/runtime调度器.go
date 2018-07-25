package main
/*
并行与并发 :
并行 : 同一时刻,多个CPU同时执行多个指令
并发 : 同一时刻,执行一条指令,cpu按时间交替执行多个指令
 */
import (
	"runtime"
	"fmt"
)

/*
1) 延迟执行该协程  :  避免 主协程提前完毕, 子协程执行不到
runtime.Gosched()
2) 立即终止该协程  :  终止后该协程下 未运行的程序不再运行
runtime.Goexit()
3) 设置程序运行使用的cpu核数(线程)最大值(不使用此函数默认开启最大值) :
int := runtime.GOMAXPROCS(1)  //代表最大使用单核 ,返回当前进程可用的逻辑CPU数目
 */
func main(){
go func() {
	i := runtime.GOMAXPROCS(1)  //设置使用单核运行
	fmt.Println(i)		//打印当前cup核数
	runtime.Goexit()	//退出协程
}()
	for{					//无限循环打印cup运行双协程结果
		go fmt.Print(1)  // 子协程 1
		fmt.Print(0)		// 主协程 0
	}
	runtime.Gosched() 	//延迟 main协程
/* 空的死循环(可以让该协程无限运行) : for{}  */
}