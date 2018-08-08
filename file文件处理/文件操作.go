package main
import (
	"io/ioutil"
	"fmt"
	"os"
)
func check(e error){  //创建一个检查器函数,以检查错误时使用
	if e != nil{
		panic(e)
	}
}
func main(){
// 1) 新建文件 :
// 已有该文件则返回错误,相当于OpenFile的快捷操作，打开后等同于OpenFile(name string, O_CREATE,0666)
fff,err := os.Create("newFile.txt") //当前目录新建文件 newFile.txt,返回文件描述符(句柄)和错误(可忽略)
check(err)		//检查错误
s := "写入成功"  //初始化字符串
num,err := fff.WriteString(s) //写入字符串到fff文件变量,返回写入字节数和错误(可忽略)
fmt.Printf("写入%d个字节,内容为: %s\n",num,s)  //打印结果
// 2) 读取文件 :
//  把整个文件的内容读取到内存 :
dat,err := ioutil.ReadFile("./readerFile.txt") //读取文件
check(err)	//检查错误
fmt.Println(string(dat)) //输出文件内容到终端
// 3) 打开文件 :
// 打开一个文件，返回文件描述符，该文件描述符只有只读权限．他相当于OpenFile(name string,O_RDWR,0)
//读取指定长度:
f,err := os.Open("./readerFile.txt") //返回文件描述符(句柄) dat1
check(err)		//检查错误
b11 := make([]byte,11)  //声明 长度11 byte数组
d1,err := f.Read(b11) //把内容读取到字节数组中,最多11字节,返回读取数量和错误
check(err)			//检查错误
	fmt.Printf("%v 个字节为: %v\n", d1, string(b11)) //打印结果
// 读取文件指定位置:
	//从头(0文件头,1当前位置,2文件末尾)开始 ,文件指针偏移 5 byte
number,err := f.Seek(5,0)   // 得到文件位置变量 date1
check(err)		//检查错误
b5 := make([]byte,5)	//声明 byte数组 b5
d2,err := f.Read(b5)	//读取文件内容到 b5数组,并得到读取的字节数 d2
check(err)		//检查错误
fmt.Printf("@%d位置读取了%d字节,内容为 :%v\n",number,d2,string(b5)) //打印读取内容
// 自定义模式打开文件,并返回文件描述符(句柄) :
	ff,err := os.OpenFile("./writeFile.txt",os.O_APPEND,0666)
	check(err)
// * 写文件 :
// 1) 无文件则新建,有文件则清空后写入 :
bs := []byte{'h','e','l','o','o'}
errr := ioutil.WriteFile("./writeFile.txt",bs,644)
check(errr)
// 2) 使用格式化打开文件,追加写入内容方案 :
nu,err := ff.Write(bs)
if err ==nil && nu < len(bs){
	return
}
	defer f.Close()		// 延迟关闭文件 ...
		 ff.Close()
		fff.Close()
/*创建一个新目录1):该目录具有FileMode权限，当创建一个已经存在的目录时会报错
func Mkdir(name string, perm FileMode) error　
创建一个新目录2):该目录是利用路径（包括绝对路径和相对路径）进行创建的，
如果需要创建对应的父目录，也一起进行创建，如果已经有了该目录，则不进行新
的创建，当创建一个已经存在的目录时，不会报错
func MkdirAll(path string, perm FileMode) error
创建临时目录用来存放临时文件，这个临时目录一般为/tmp
func TempDir() string
*/
/* 删除文件或者目录1) :
func Remove(name string) error
   删除目录以及其子目录和文件 :
func RemoveAll(path string) error
 */
 /* 重命名文件 :
 func Rename(oldpath, newpath string) error
 复制文件 :
io.Copy()
io.CopyN()
io.CopyBuffer()
 同步操作，将当前存在内存中的文件内容写入硬盘．
func (f *File) Sync() (err error)
 关闭文件，使其不能够再进行i/o操作，其经常和defer一起使用
 func (f *File) Close() error
  */
}
