package basic_grammar

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
	beego.Run()
	//beego.Run("127.0.0.1:6666") //为什么访问不了？？？
}
