package basic_grammar

import (
	"fmt"
	"time"
)

var channel01 chan string
func MultiRoutine(){
	//go ReceiveMsg()
	//go SendMsg()
	go SelectDemo()
	channel01 = make(chan  string,5)//创建具有固定缓冲数量的通道避免阻塞，因为goroutine的通道默认是阻塞的
	//channel01 = make(chan  string)//需要创建才能执行下面的发送语句，不然就阻塞了,goroutine的非缓冲通道里头一定要一进一出，成对出现才行，否则会出现死锁
	for i:=0;i<5;i++{
		channel01 <- "hello server"//往通道写消息,如果是没有缓冲区的通道的话是会发生阻塞的，
	}

	//fmt.Println(<-channel01)
	time.Sleep(2*time.Second)
	fmt.Println("exiting......")
}

func ReceiveMsg(){
	msg:= <- channel01//从通道取出消息
	fmt.Println(msg)
	fmt.Println("hello,i am receiver")
	channel01<-"job finshed....."//测试看看是不是可以双向收发
}
func SendMsg(){
	fmt.Println("hello,i am sender")
}

func SelectDemo(){
	//如何才能多次取呢？
	//select{
	//case <-channel01:
	//	fmt.Println("case runing...")
	//case <-channel01:
	//	fmt.Println("case 2 running...")
	//}
	fmt.Printf("channel msg num is:%d \n",len(channel01))
	for len(channel01)>0{
		select{
		case <-channel01:
			fmt.Println("case 1 runing...")
		case <-channel01:
			fmt.Println("case 2 running...")
		default:
			fmt.Printf("channel msg is not handled")
		}
	}
}




