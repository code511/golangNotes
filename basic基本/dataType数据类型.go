package mian

/*数据类型
1) 布尔类型 bool (布尔零值(初始值)为false,其余皆为0,长度1)
	turn(真)
	false(假)
2) 字符型 rune     (对应ascii码表,同uint32,长度1)
	'A' + 32 = 'a'(通过ascii码表转换大小写)  字符用''表示
3) 字符串型 string (utf8字符串,长度1)
	str := "1234"
	len(str) = 4  len是得到字符串长度的内建函数
	str[0] = 1    字符串可以进行角标操作
4) 字节型 byte     (同uint8,长度4)
5) 整型
	int8,int16,int32,int64	(长度1,2,4,8,同下)
	uint8,uint16,uint32,uint64 (根据CPU计算)
	uint ptr (足以储存指针的uint32或uint64整数,长度4或8)
6) 浮点型
	float32 (长度4)
	float64	(长度8)
7) 复数类型
	complex64   (长度8)
	complex128  (长度16)
	格式:
	b3 := 3.14(实数) + 5.12i (加i代表虚数)
8) 函数类型 :
	函数也是一种数据类型,通过 type 给一个函数类型起名
	函数的参数返回值类型和定义的函数类型对应时才可以给函数类型赋值		*/

type lx func(int, int) int //定义了一个叫 lx 的函数类型
// 关键字 命名 关键字     形参	  返回值 (参数要一一对应)
func main() {
	var jjj lx          //声明一个函数类型的变量 jjj
	jjj = jia           //向 jjj 赋值 jia 函数  ,赋予了jjj和jia同等的功能
	j := jjj(1, 2)      //使用 jjj 函数类型
	println("1+2 =", j) //打印使用 jjj 的结果
	var xxx lx = jian   //声明一个函数类型 xxx ,并赋值
	x := xxx(1, 2)      //使用 xxx
	println("1-2 =", x) //打印结果
}
func jia(a, b int) int { //相加函数
	return a + b
}
func jian(a, b int) int { //相减函数
	return a - b
}
