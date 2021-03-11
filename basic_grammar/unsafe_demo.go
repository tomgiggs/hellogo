package basic_grammar

import (
	"fmt"
	"time"
	"unsafe"
)
type Cat struct {
	height int
	weight int32
	age int
}

func UnsafeDemo(){
	one := Cat{
		height:200,
		weight:int32(2000000),
		age:30,
	}
	fmt.Println(unsafe.Sizeof(int(2000)))
	fmt.Println("cat size is:",unsafe.Sizeof(one))//24，因为内存对齐了，字段顺序影响了内存大小
	fmt.Println(unsafe.Alignof(one))//等价于reflect.TypeOf(x).Align()，对齐值=min(默认对齐值，字段最大类型长度)
	fmt.Println(unsafe.Offsetof(one.age))
}

type Job struct {
	CC <- chan int//只读通道
	CC2 chan <- int//只写通道
	Name string
}

func ChanDemo01(){
	//read_only := make (<-chan int)//定义只读和只写的channel意义不大，一般用于在参数传递中
	//write_only := make (chan<- int)
	read_write := make (chan int,20)
	go func() {
		for i:=0;i<20;i++{
			read_write <- i
			WriteOnlyChan(read_write)
		}
	}()
	go func() {
		ReadOnlyChan(read_write)
	}()
	time.Sleep(time.Second*10)

}
func ReadOnlyChan(notify <- chan int){
	for n := range notify{
		fmt.Println("read from chan,",n)
	}
}
func WriteOnlyChan(notify chan<- int){
	notify <- 290
}
