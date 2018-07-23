package main

func main(){
// 打开/创建文件示例: 	os.OpenFile(“output.dat”,  os.O_WRONLY|os.O_CREATE, 0666)
// 解释  :  打开文件 output.dat, 执行只写和创建(如果不存在则新建)操作,所有权限设置为0666(可读可写)
// 第一个参数 : 文件名/路径
// 第二个参数 : 文件打开模式(多个用 或| 符号间隔) :
							/* 1. os.O_WRONLY：只写
							   2. os.O_CREATE：创建文件
							   3. os.O_RDONLY：只读
							   4.  os.O_RDWR：读写
							   5.  os.O_TRUNC ：清空
	// os.O_RDONLY // 只读
    // os.O_WRONLY // 只写
    // os.O_RDWR // 读写
    // os.O_APPEND // 往文件中追加（Append）
    // os.O_CREATE // 如果文件不存在则先创建
    // os.O_TRUNC // 文件打开时裁剪文件
    // os.O_EXCL // 和O_CREATE一起使用，文件不能存在
    // os.O_SYNC // 以同步I/O的方式打开
							*/
// 第三个参数 : 权限参数 : 		相加使用  6 == 4(r) + 2(w) == 可读可写
							/* r ——> 004  可读reader
							   w ——> 002  可写write
							   x ——> 001  可执行
							*/
							// 例 : 0 				6				 6				 6
							//	   特殊 	  文件所有者（owner）	   	用户组(group) 	其他人(others)

}