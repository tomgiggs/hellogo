package basic_grammar

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

var a = flag.String("a", "", "全部使用")
var b = flag.Bool("b", false, "bool类型参数")
var s = flag.String("s", "", "string类型参数")
var t = flag.String("t","ok","ssssss")



func GetSysInfo() {
	//解析命令行参数方式1：
	// for idx, args := range os.Args {
	// 	fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	// }
	//解析命令行参数方式2：
    flag.Parse()
	fmt.Println("-a", *a)
	fmt.Println("-b:", *b)
	fmt.Println("-s:", *s)
	fmt.Println("其他参数：", flag.Args())
	fmt.Println(runtime.GOOS)
	fmt.Println(time.Now())
    fmt.Println(runtime.NumCPU())
    pwd,_ := filepath.Abs("./")

    fmt.Printf("current dir is: %s",pwd)
	// var timestamp int64
	// timestamp, _ = time.Parse("2019-01-02 15:04:05", "2018-04-23 12:24:51")
	// fmt.Println(timestamp.Unix())
	//dt, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51") //这写法真奇怪
	//fmt.Println(dt.Unix())

}
