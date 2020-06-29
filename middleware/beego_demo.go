package middleware

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func BeegoDemo() {
	log.Info("this is a logrus info msg")
	beego.Router("/", &MainController{})
	//beego.Run()
	beego.Run("127.0.0.1:8080") //为什么访问不了？？？可能是因为端口号太低不允许还是什么原因，端口号改成大于8080的可以
}
