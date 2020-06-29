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
