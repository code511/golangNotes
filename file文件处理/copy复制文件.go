package main
/*
Go语言中拷贝文件的几种常用的方式

简介
本篇文章将介绍Go语言中，最最最常用的3种拷贝文件的方法，这三种方法各有利弊，我们只需要在应用中选择最合适的即可，
不必盲目追求性能。

方法一
第一个方法将使用标准Go库的io.Copy()函数。以下是使用io.Copy()实现的拷贝文件代码片段：


func copy(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}

方法二
第二中方法是使用ioutil包中的 ioutil.WriteFile()和 ioutil.ReadFile()，但由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出。


 input, err := ioutil.ReadFile(sourceFile)
        if err != nil {
                fmt.Println(err)
                return
        }

        err = ioutil.WriteFile(destinationFile, input, 0644)
        if err != nil {
                fmt.Println("Error creating", destinationFile)
                fmt.Println(err)
                return
        }

方法三
最后是使用os包中的os.Read() 和 os.Write()，此方法是按块读取文件，块的大小也会影响到程序的性能。


  buf := make([]byte, BUFFERSIZE)
        for {
                n, err := source.Read(buf)
                if err != nil && err != io.EOF {
                        return err
                }
                if n == 0 {
                        break
                }

                if _, err := destination.Write(buf[:n]); err != nil {
                        return err
                }
        }

性能 :
这三种方式都能很方便的实现拷贝文件功能，那他们的性能如何呢，下面我们来尝试对比一下。三种方式都来拷贝同一个500M的文件
从数据来看3种方式的性能非常接近，但依然可以看出Go的io标准包的性能更优。
 */
func main(){
/*
怎样使用go语言将文件复制成另一文件？go语言标准包io里提供一个简单的函数copy实现了这个功能,下面是一个例子。

package main

import (
    "fmt"
    "io"
    "os"
)
func main() {
    CopyFile(os.Args[1], os.Args[2]) // os.Args[1]为目标文件，os.Args[2]为源文件
    fmt.Println("复制完成"，)
}
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()
    dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return
    }
    defer dst.Close()
    return io.Copy(dst, src)
}

注意这里defer的使用，在源文件和目的文件打开后都跟了一个defer延迟关闭文件。如果目标文件后面没有使用defer dst.Close(),那么一旦创建目标文件调试失败，它将直接返回错误，那么，会导致源文件一直保持打开状态，这样资源就得不到释放。因此，在go语言里，记得open一个文件做了错误判断后要紧跟一个defer close延迟调用。

io.Copy的函数原型如下
func Copy(dst Writer, src Reader) (written int64, err error)
Copy函数从源到目的的复制，直到读到源的EOF或者出现其它的错误，它返回所复制的字节数及复制过程出现的第一个错误。如果成功复制，Cop返回的是 err ==nil,因为Copy被定义为从源复制直到遇到EOF,它不把EOF当作一个错误。
io包里还有一个io.CopyN的函数，原型如下
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
CopyN将从源文件里复制n字节（或者遇到错误中断）到目标文件，并返到实际复制的节数
package main

import (
    "fmt"
    "io"
    "os"
)
func main() {
    counter,  := CopyFile(os.Args[1], os.Args[2], os.Args[3]) // os.Args[1]为目标文件，os.Args[2]为源文件
     fmt.Println("复制", counter, "字节!")
}
func CopyFile(dstName, srcName string, n int64) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()
    dst, err := os.OpenFile(dstName, os.OWRONLY|os.OCREATE, 0644)
     if err != nil {
        return
    }
    defer dst.Close()
    return io.CopyN(dst, src, n) //
}
复制 526 字节!
 */
}