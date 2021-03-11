package main

import (
	"fmt"
	"hellogo/basic_grammar"
	_ "net/http/pprof"
)

//str := ""//包内必须以var const import func type  const, func,开头,短变量用来声明和初始化函数内部的局部变量
//var str string
//var str  = ""

//type People interface {
//	Show()
//}
//
//type Student struct{
//	name string
//}
//
//func (stu *Student) Show() {
//	//fmt.Println("nice show")
//	fmt.Println("nice show,",stu.name)
//}
//
//func live() People {
//	//var stu *Student//这个会引起panic,new不会，声明变量就可以调用对应的成员函数
//	var stu = new(Student)
//	//var stu =&Student{
//	//	name:"steve",
//	//}
//	return stu
//}

func main() {
	defer func() {
		if err := recover();err!= nil{
			fmt.Println("panic info: ", err)
		}
	}()

	//用于配合gops进行性能查看统计
	//if err2 := agent.Listen(agent.Options{
	//	Addr:"0.0.0.0:9981",
	//	//ConfigDir:"",
	//	ShutdownCleanup:true,
	//});
	//err2 != nil {
	//	fmt.Println(err2)
	//}
	//go func() {
	//	http.ListenAndServe("0.0.0.0:7777", nil)
	//}()
	//time.Sleep(time.Second*60)

	//--------------
	//basic_grammar.PrintNumAndLetter()
	//basic_grammar.UnsafeDemo()
	//algorithm.Visit()
	basic_grammar.ChanDemo01()
		//pprof性能监控

	//-----------
	//xxx := live()
	//if xxx == nil {
	//	fmt.Println("AAAAAAA")
	//} else {
	//	fmt.Println("BBBBBBB")
	//	xxx.Show()
	//}


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
	// basic_grammar.GroupConsume02()
	//basic_grammar.MultiRoutine()
	//basic_grammar.GlogDemo()
	//basic_grammar.ConstDemo()
	//basic_grammar.SliceDemo07()
	//basic_grammar.Update()
	//basic_grammar.Xxxxxxxx()
	//basic_grammar.Container01()
	//basic_grammar.RegDemo01()
	//basic_grammar.WriteFileDemo()
	//basic_grammar.TypeConvert()
	//middleware.StartGrpcClient()

	//go func() {
	//	//middleware.LinkStart()
	//	middleware.LinkClientStart()
	//
	//}()
	//middleware.LinkStart()
	//middleware.ContextStart()
	//fmt.Println(3|1)
	//mapAgainstId:= make(map[int32][]int32, 0)
	//mapAgainstId :=map[int32][]int32{}
	//
	//if againstId, ok := mapAgainstId[0]; ok {
	//	fmt.Println(againstId)
	////}
	//a := make([]int,5,10)
	//a=nil
	//fmt.Println(len(a))
	//middleware.GuiDemo01()
	//fmt.Println(3<<10|711)
	//middleware.CassandraDemo()

	//fmt.Println(3791-3072)

	//basic_grammar.GetLocalIp()
	//middleware.RedisMapDemo()
	//runes := []rune("rune22 test 001")
	//fmt.Println(runes[1:5])
	//fmt.Println(strings.Count("goood",""))

	//middleware.RedisPubDemo()
	//fmt.Println(strconv.Itoa(1))
	//fmt.Println(strconv.FormatInt(int64(200),10))
	//middleware.ServiceDiscover()
	//uuid.New()
	//middleware.PythonDemo2()

	//middleware.OrmDemo3()
	//a := 64
	//b := 1<<6
	//fmt.Println(a&b)
	//for i := int32(0); i < 5; i++ {
	//	fmt.Println(i)
	//}
	//tmpList := []int{1,2,3,4,5,6}
	//fmt.Println(append(tmpList[:2],tmpList[3:]...))
	//time.Sleep(time.Minute*10)
	//---------------
	//var(
	//	is_ok bool
	//	st string
	//)
	////basic_grammar.ByteDemo01()
	//_, err := fmt.Sscanln("f string", &is_ok, &st) //true --> t,false---> f
	//if err != nil {
	//	fmt.Println("错误:", err)
	//}
	//fmt.Println(is_ok, st)

	//middleware.LinkClientStart()
	//middleware.GenObj()
	//middleware.StartGrpcServer()
	//basic_grammar.MapNil()
	//middleware.CalDigest()
	//middleware.Log4goDemo01()
	//middleware.LuaDemo02()
	//basic_grammar.ReflectDemo02()

	//basic_grammar.SwitchDemo01()
	//basic_grammar.TimerDemo()
	//basic_grammar.TickerDemo()
}
