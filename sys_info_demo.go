package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(time.Now())
	// var timestamp int64
	// timestamp, _ = time.Parse("2019-01-02 15:04:05", "2018-04-23 12:24:51")
	// fmt.Println(timestamp.Unix())
	dt, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51") //这写法真奇怪
	fmt.Println(dt.Unix())

}
