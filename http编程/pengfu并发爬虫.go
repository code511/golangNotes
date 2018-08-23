package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"os"
	"strings"
)

/*
爬虫四步骤:
1)明确目标(爬哪个网站,以及规律)
捧腹网段子页面链接(含十个段子) :
https://www.pengfu.com/xiaohua_1.html  下一页 +1
https://www.pengfu.com/xiaohua_2.html
单独段子链接 :
<h1 class="dp-b"><a href=" 链接 " target="_blank">
2)爬(将每一页的全部内容爬下来)
从大页面中爬出小页面地址
3)取(取出需要的数据)
标题 :
<h1>标题</h1>
内容 :
<div class="content-txt pt10"内容<a id="prev"
4)处理数据(按照需求的方式储存和使用数据)
每个大页面按一标题一内容的形式储存在一个文本文件中
 */
func main() {
	var start, end int
	fmt.Println("请输入起始页")
	fmt.Scan(&start)
	fmt.Println("请输入结束页")
	fmt.Scan(&end)
	gut(start, end)  //取大页面body内容
}
func gut(start, end int) {
	fmt.Printf("准备爬取第%d页到第%d页\n", start, end)
	page := make(chan int) //创建调控通道
	for i := start; i <= end; i++ { //每页
		go  oneUrl(i,page)	//循环新开协程 ,取小页面body内容
	}
	for i := start;i <= end; i++{
		n := <-page	//把通道数据 : 当前完成的页数,赋值给 n
		fmt.Printf("第%d页提取完毕\n",n)
	}
}
func oneUrl(i int,page chan<- int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html" //确定大页面地址
	resp, err := http.Get(url)	//get请求大页面,得到接口 resp
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close() //延迟关闭 请求接口
	buf := make([]byte, 2048*4) //声明 字节切片 buf
	var data string	//声明 字符串 data
	for {
		n, _ := resp.Body.Read(buf) //把大页面body内容读取到 buf 中
		data += string(buf[:n])		//再赋值给 字符串 data
		if n == 0 {
			//fmt.Println("大页面读取完毕")
			break
		}
	}
	f,err := os.Create("段子第"+strconv.Itoa(i)+"页.txt") //创建文件
	if err !=nil {
		fmt.Println("os.Create err",err)
		return
	}
	defer f.Close() //延迟关闭文件
	er := regexp.MustCompile(`<h1 class="dp-b"><a href="(.*)" target="_blank">`) //创建正则解析器
	if er == nil {
		fmt.Println("regexp.MustCompile : oneUrl err", er)
		return
	}
	oneUrl := er.FindAllStringSubmatch(data, -1) //解析大页面body内容,并保存在 oneUrl中
	for i, _ := range oneUrl {	//遍历 oneUrl
	title,article := oneGut(i, oneUrl)  //取小页面内body内容函数 oneGut,返回标题和正文
	f.WriteString(title+"\n")	//将标题写入文件
	f.WriteString(article+"\n") //将正文写入文件
	f.WriteString("----------------------------------------------------\n")//写入分割线
	}
	page <- i //向通道写入数据 :当前页数 i
}
func oneGut(i int, oneUrl [][]string)(title,article string) { //提取小页面body内容
	resp, err := http.Get(oneUrl[i][1]) //get请求当前小页面
	if err != nil {
		fmt.Println("http.Get :oneUrl err", err)
		return
	}
	defer resp.Body.Close() //延迟关闭请求接口
	buf := make([]byte, 1024*4) //声明 字节切片 buf
	var data string		//声明 字符串 data
	for {
		n,_ := resp.Body.Read(buf) //将小页面body内容读取到 buf中
		data += string(buf[:n])		//再将buf内容赋值给 data
		if n == 0 {
			//fmt.Println("小页面读取完毕")
			break
		}
	}
	s1,s2 := bodyGut(data) //处理小页面body内容
	title = s1  //返回值
	article = s2//...
	return
}
func bodyGut(data string) (title,article string) {
	tit := regexp.MustCompile(`<h1>(.*)\s*</h1>`) //创建正则解析器
	if tit == nil{
		fmt.Println("regexp.MustCompile : title err",tit)
		return
	}
	titler := tit.FindAllStringSubmatch(data,-1)   //解析小页面body内容,赋值给 title
	arti := regexp.MustCompile(`<div class="content-txt pt10">\s*(.*|.*\n*.*)\s*<a id="prev"`)//解析器
	if arti == nil{
		fmt.Println("regexp.MustCompile : article err",arti)
		return
	}
	articler := arti.FindAllStringSubmatch(data,-1)//解析小页面body内容,赋值给 articler
	if len(titler)>0{
	title = strings.Replace(titler[0][1],"  ","",-1) //得到标题 ,返回
	}
	if len(articler)>0{
	s := strings.Replace(articler[0][1]," ","",-1) //替换过滤
	s1 := strings.Replace(s,"&nbsp;","",-1)			//...
	s2 := strings.Replace(s1,"<p>","",-1)
	s3 := strings.Replace(s2,"</p>","",-1)
	s4 := strings.Replace(s3,"<br>","",-1)
	s5 := strings.Replace(s4,"<br/>","",-1)
	article = s5		//得到正文 ,返回
	}
	return
}
