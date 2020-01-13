package main

import "hellogo/basic_grammar"

func main(){
	//运行报错： use of internal package github.com/go-redis/redis/internal/hashtag not allowed，这个是因为src文件夹底下有多个重名文件夹
	//basic_grammar.GetSysInfo()
	//basic_grammar.ArrayDemo()
	//basic_grammar.Logic()
	//basic_grammar.MapDemo()
	//--------------------
	//这两种方式声明变量是有区别的，使用:= 可以直接赋值，使用var需要额外赋值
	//var dog1 basic_grammar.Dog
	//dog1.Name = "旺财"
	//dog2 :=basic_grammar.Dog{
	//	Name :"小黑",
	//	Price:2000,//这里需要多一个逗号，不然提示语法错误
	//}
	//dog1.Say()
	//basic_grammar.Create(dog2)
	//可以直接调用
	//dog.Price = 2000
	//dog.Say()
	//foods :=[]string{"meat","rice","fruit"}
	//dog.Eat(foods)
	//dog.Sleep(2)
	//fmt.Printf("my price is:%f ￥",dog.Sell())
	//-----------------------
	//basic_grammar.GetDBData()
	//basic_grammar.IoDemo()
	//basic_grammar.GoRedisDemo()
	//basic_grammar.RedigoDemo()
	//basic_grammar.OrmDemo2()
	//basic_grammar.BeegoDemo()
	//basic_grammar.StartHttpServer()
	//basic_grammar.WebSocketDemo()
	basic_grammar.GroupConsume02()
	//basic_grammar.MultiRoutine()


}


