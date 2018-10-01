package main

import (
	"html/template"
	"fmt"
	"net/http"
)

/*
模版为template,之所以叫做模板是因为其由静态内容和动态替换内容组成，根据动态内容的变化而最终生成
客户端确认到的不同的内容信息，这就是模板的重要作用
 */

func main() {
	type persen struct {
		//创建结构体类型 人
		ID      int    //ID
		Name    string //姓名
		Country string //国家
	}
	//注册处理函数 :
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		xiaoming :=   //声明并赋值persen类型变量 xiaoming
			persen{001, "xiaoming", "china"}
//格式1: 模版接口							模版名		调用
		taml, err := template.New("tmpl").Parse(`
hello,{{.Name}}		//创建模版 taml以及创建调用数据
your ID :{{.ID}}`)
//格式2: 调用创建的模版文件 main.tmpl 来操作(文件名格式任意,内容可以使用模版引擎支持的html,go模版语法)
//teml,err :=template.ParseFiles("./main.tmpl")
		if err != nil {
			fmt.Println("", err)
			return
		}
		// err = tmpl.ExecuteTemplate(os.Stdout, "tmpl", sweaters)  //指定模板名 进行渲染
		err1 := taml.Execute(w, &xiaoming) //(渲染模版)将数据对象指针和输出方式用于模版
		if err1 != nil {
			fmt.Println("", err1)
			return
		}
	})
	http.ListenAndServe(":8000", nil) //绑定监听
}
