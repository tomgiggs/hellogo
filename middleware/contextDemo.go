package middleware

import (
"context"
"fmt"
"time"
)

func ContextStart(){

	ctx,cancel := context.WithCancel(context.Background())//使用一个context很简单，在被调函数里面添加一个select 块监听退出信号即可
	go func1(ctx)
	go func2(ctx)

	time.Sleep(7*time.Second)
	cancel()
	fmt.Println("main over")
}

func func1(ctx context.Context){
	i:=0
	for {
		time.Sleep(1*time.Second)//这里休眠了会导致在休眠期间收到退出信号时后面的逻辑无法执行
		i++
		select {
		case <- ctx.Done():
			fmt.Println("d1 over")
			return

		default:
			fmt.Println("d1 ",i)
		}
	}
}

func func2(ctx context.Context){

	fmt.Println("func2 start")
	<- ctx.Done()
	fmt.Println("func2 over")
}
