package routers

import (
	"hellogo/webgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/usercenter", &controllers.UserController{})
}
