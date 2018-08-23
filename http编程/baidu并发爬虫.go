package main

/*
爬虫四步骤:
1)明确目标(爬哪个网站,以及规律)
百度贴吧-golang
https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=0
https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=50
https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=100
2)爬(将每一页的全部内容爬下来)
3)取(取出需要的数据)
4)处理数据(按照需求的方式储存和使用数据)

 */
import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

func main() {
	var start, end int //声明页数变量
	fmt.Println("请输入起始页") //提示
	fmt.Scan(&start)	//输入赋值给 start
	fmt.Println("请输入终止页") //提示
	fmt.Scan(&end)		//输入赋值给 end
	doWork(start, end)  //调用函数,导入首尾页数
}
func doWork(start, end int) {
	page := make(chan int) //声明通道
	fmt.Printf("爬取第%d页到第%d页\n", start, end) //提示
	for i := start; i <= end; i++ { //循环运行函数
		go spiderPape(i, page) //并发运行函数,导入i 和通道 page
	}
	for i := start; i <= end; i++ { //跳出后依次
		fmt.Printf("第%d页爬取完成\n", i)//提示
		<-page		// 丢掉通道数据
	}
}
func spiderPape(i int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)//得到地址
	fmt.Printf("正在爬取第%d页:%v\n", i, url) //提示
	result, err := httpGet(url) //运行函数,导入地址,得到网页内容 result
	if err != nil {  //检查错误
		fmt.Println("httpGet err :", err)
		return
	}
	fileName := "golang贴吧第" + strconv.Itoa(i) + "页.html" //创建文件名变量 fileName
	f, err2 := os.Create(fileName)	//根据 fileName创建文件
	if err2 != nil {	//检查错误
		fmt.Println("os.Rreate err :", err2)
		return
	}
	f.WriteString(result) //将网页内容 result写入文件
	defer f.Close()	//延迟关闭文件
	page <- i	//向通道发送数据
}
func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //使用get方法请求地址
	if err1 != nil {	//检查错误
		err = err1		//将错误返回
		return
	}
	defer resp.Body.Close() //延迟关闭网页访问
	buf := make([]byte, 1024*4) //声明数组
	for {
		n, _ := resp.Body.Read(buf) //将网页内容读取到 buf 数组,忽略错误
		if n == 0 {	//如果读取完毕
			//fmt.Println("resp.Body.Read err :", err)
			break	//跳出
		}
		result += string(buf[:n]) //将读取到的内容赋值给 result 并返回
	}
	return
}
