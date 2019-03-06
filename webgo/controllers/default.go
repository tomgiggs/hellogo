package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	//_ "github.com/astaxie/beego/cache/redis"
	//"github.com/astaxie/beego/cache"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Get(){

	//redis, err := cache.NewCache("redis", `{"key":"usercount_conn","conn":":6379","dbNum":"0"}`)

	username := c.GetString("username")
	islogin,err:=c.GetBool("islogin")
	if err!=nil{
		log.Info("miss islogin param ,please check your request")
	}
	fmt.Println(username)
	if islogin{
		log.Info("user has login:")
	}

	log.Info("the username is :"+username)
	//这种方式可以使用，但是每次连接度新建一个Redis客户端，这样消耗是很大的，怎么把连接放到全局变量中，然后做个连接池呢？
	redis_client, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer redis_client.Close()
	redis_client.Do("incrby","usercount","1")


	usernum,_:=redis.String(redis_client.Do("get","usercount"))

	c.Ctx.WriteString("hello, this is  user center，you are the %s: user"+usernum) //好像没法想Python那样使用%来生成字符串
}


