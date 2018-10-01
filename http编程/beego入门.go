package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}
func main(){
	beego.Router("/",&HomeController{})
	beego.Run()
}
func (this *HomeController)Get(){
	this.Ctx.WriteString("hello word")
}