package basic_grammar

import (
	"fmt"
	"time"
)

func TimerDemo() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(10 * time.Second)
	fmt.Printf("%v\n", time.Now().Add(time.Second*10))
	//timer.Stop()
	<-timer.C
	print("timeout,finshing\n")

}
func TimerDemo02() {
	time.AfterFunc(3*time.Second, func() { //如果在函数调用前程序已经退出，那time.after注册的函数不会被调用
		fmt.Println("time after func called")
	})
	time.Sleep(4 * time.Second)

	fmt.Println("waiting for time after func\n")
}
func TickerDemo() {
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for t := range ticker.C {
			//fmt.Println(t.Format("20060102150405"))
			fmt.Println(t.Format("2006-01-02-15-04-05"))

			fmt.Println("I Love You!")
		}
	}()

	time.Sleep(time.Second * 18)
	//停止ticker
	ticker.Stop()
}

func IsSameWeek(){
	weekSecond := 3600*24*7
	//curTime := time.Now()
	lastWeek,_ := time.ParseInLocation("2006-01-02 15:04:05","2020-11-29 16:59:59",time.Local)
	curTime,_ := time.ParseInLocation("2006-01-02 15:04:05","2020-12-05 16:00:01",time.Local)
	//fmt.Println(lastWeek,err1)
	fmt.Println(curTime.ISOWeek())
	fmt.Println(lastWeek.ISOWeek())
	fmt.Println(lastWeek.Unix()/int64(weekSecond))
	fmt.Println(curTime.Unix()/int64(weekSecond))
	//t := curTime.Unix() + 3 * 24 * 60 * 60
	t := curTime.Unix() + 3 * 24 * 60 * 60 + 8 * 60 * 60
	t = t/int64(weekSecond)
	//t2 := lastWeek.Unix() + 3 * 24 * 60 * 60
	t2 := lastWeek.Unix() + 3 * 24 * 60 * 60 + 8 * 60 * 60
	//time.Unix(curTime.Unix(),0)
	t2 = t2/int64(weekSecond)
	fmt.Println(t-t2)
}