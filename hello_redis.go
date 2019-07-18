package main

import (
// "fmt"
// "time"
//"github.com/garyburd/redigo/redis"
// "github.com/go-redis/redis"
)

func TestRedis() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "Mm!Ou@S2e1R", // no password set
	// 	DB:       0,  // use default DB
	// })

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
	// err = client.Set("hellosssss","nice_day",5*time.Second).Err()
	// fmt.Println(err)
	// val,_ :=client.Get("chat-server-1").Result()
	// fmt.Print(val)
	//c, err := redis.Dial("tcp", "127.0.0.1:6379",&redis.DialOption{
	//
	//})//如何传递一个密码呢？密码可以在dialOptional里面加，但是不知道如何构造那个参数
	////c.Do("auth","dialOptions ")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//defer c.Close()
	//
	//_, err = c.Do("SET", "mykey", "superWang")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}
	//
	//username, err := redis.String(c.Do("GET", "mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//} else {
	//	fmt.Printf("Get mykey: %v \n", username)
	//}
}
