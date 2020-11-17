package middleware

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	goredis "github.com/go-redis/redis"
	"time"
)

func GoRedisDemo() {
	client := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		//Password: "password", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	err = client.Set("hello","nice_day",1000*time.Second).Err()
	if err!=nil{
		fmt.Println(err)
	}
	val,_ :=client.Get("hello").Result()
	fmt.Print(val)
}

func RedigoDemo(){
	//redis这个连接参数究竟要怎么构造？一个结构体，里面包含一个私有函数成员，难搞。。。
	//option := redigo.DialOption{
	//}
	//c, err := redigo.Dial("tcp", "127.0.0.1:6379",option)//如何传递一个密码呢？密码可以在dialOptional里面加，但是不知道如何构造那个参数
	c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	_,err = c.Do("AUTH","password")//使用命令进行授权，AUTH不能写成小写的auth，不然会报错。。。。
	if err!=nil{
		fmt.Println("password auth failed",err)
		return
	}

	defer c.Close()

	_, err = c.Do("set", "redigo_key", "i am not ok")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	info, err := redigo.String(c.Do("GET", "redigo_key"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %s \n", info)
	}
}

func RedisSubDemo(){
	client := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "root",
		DB:       0,
	})
	sub := client.Subscribe("gim-user-login-server")
	defer sub.Close()

	 for msg := range sub.Channel(){
		fmt.Println(msg.Payload)
	}

}
func RedisMapDemo(){
	client := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "root",
		DB:       0,
	})
	servermap := client.HGetAll("gim-user-login-server:111")
	fmt.Println(servermap.Result())

	//for v,msg := range servermap{
	//	fmt.Println(msg,v)
	//}

}