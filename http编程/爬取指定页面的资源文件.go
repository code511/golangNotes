package main
/*本程序由 itruirui@outlook.com 瑞哥 开发制作*/
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"regexp"
	"errors"
	"sort"
	"time"
)


var number int64	//已下载序号
var i int	//页面编号
var jpgnamemmin int //图片文件名最小字节(过滤图片使用)

func main() {
	jpgnamemmin = 10	//设置最小为10
	/*
	var min int //搜索开始地址
	var max int //搜索结束地址
	min = 2294897	//最小地址
	max = 2295897	//最大地址
	for i = max; i >= min; i--{
		url := fmt.Sprintf("http://war.163.com/photoview/4T8E0001/%d.html", i)
		fmt.Printf("获取 %d.html中\n", i)
		DownloadUrlImg(url, ".jpg")	//后面的.jpg是要下载的文件类型
		time.Sleep(time.Second)	//延时 1 秒
	}
	*/
	DownloadUrlImg("http://news.qq.com/photon/photoex.htm",".jpg")
	time.Sleep(time.Second)	//延时 1 秒
}


/*
* 作用：下载某个URL里面的资源
* 参数url：URL地址
* 参数filetype：文件类型，.jpg或者.png
* 返回值sum：一共下载多少个资源
*/
func DownloadUrlImg(url string, filetype string)(sum int){
	body, err := GetHtml(url)	//获取html代码
	if err != nil{
		fmt.Printf("获取 %d.html 代码失败！\n", i)
		return 0
	}
	imgsrc := GetImg(body, filetype)
	var sz []string
	for i := 0; i < len(imgsrc); i++{
		sz = append(sz, imgsrc[i][1])
	}
	sz = Overkill(sz)
	for _, data := range sz{
		_, err := DlImg(data)
		if err != nil{
			fmt.Println("下载失败，错误：", err)
			return 0
		}
	}
	return 0
}

/*
* 获取某个url的html代码
* 参数：url地址
* 返回值：string类型，error类型
*/
func GetHtml(url string)(value string, err error){
	resp, err := http.Get(url)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		err = errors.New(fmt.Sprintln("错误代码：",resp.StatusCode))
		return value, err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}
	return string(all), nil
}

/*
* 解析某个string里面的ed2k
* 参数：要解析的string
* 返回值：ed2k内容[][]string
*/
func GetImg(body string, filetype string)(result [][]string){
	str := fmt.Sprintf("src=\"([^\"]*%s)\"", filetype)
	reg := regexp.MustCompile(str)	//正则表达式规则
	if reg == nil{
		fmt.Println("正则表达式，编译错误：", reg)
		return
	}else{
		result = reg.FindAllStringSubmatch(body, -1)
		return result
	}
}

/*去除重复的数据*/
func Overkill(a []string) (ret []string){
		sort.Strings(a)
    a_len := len(a)
    for i:=0; i < a_len; i++{
        if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
            continue;
        }
        ret = append(ret, a[i])
    }
    return ret
}

/*
* 作用：获取图片资源
* 参数：图片资源的URL地址
* 返回值：图片大小，错误
*/
func DlImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	if len(name) <= jpgnamemmin{	//如果图片名称小于10个字节，就会跳过不下载
		return
	}
	name= fmt.Sprintf("DownloadImg/%d_%s", number, name)	//DownloadImg当前目录下的文件夹，必须要有
	number++
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}