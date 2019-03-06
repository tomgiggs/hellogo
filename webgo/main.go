package main

import (
	_ "hellogo/webgo/routers"
	"github.com/astaxie/beego"
)
//beego可以实现代码热更新，跟flask一样

func main() {
	beego.Run()
}

