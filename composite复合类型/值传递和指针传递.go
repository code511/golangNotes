package main

/*	值传递和指针传递	:
	1) 值传递 :
	只使用该指针的值,不改变该指针的值*/
func zcd(a int) {
	a = 888
	println(a)
}

/*	2) 指针传递 :
	把指针传递进形参,操作会改变该指针的值*/
func zzcd(a *int) {
	*a = 777
}
func main() {
	a := 666
	//值传递 :
	zcd(a)
	println(a) //打印外部 a
	//指针传递 :
	zzcd(&a)   //传递 a的指针
	println(a) //指针传递后,外部 a的值已发生改变
}
