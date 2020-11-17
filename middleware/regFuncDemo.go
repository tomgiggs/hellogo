package middleware

import "fmt"

func Begin() {
	proxyFunc(2, func(i int) {
		fmt.Println("i am callBack, input is: ",i)
	})
}
//
////需要传递函数
//func callback(i int) {
//	fmt.Println("i am callBack, input is: ",i)
//}

func proxyFunc(i int, f func(int)) {
	//exec(i, Stub(f))//把传入的函数 f 强转成 Call 类型，这样 f 就实现了 Call 接口,将目标函数转换成接口函数
	f(i)//这样就行了，为什么还要用上面那种复杂的写法来调用呢？
}

func exec(i int, c Call) {
	c.call(i)
}

//定义的type函数
type Stub func(int)
//Stub实现的Call接口的call()函数
func (s Stub) call(i int) {
	s(i)
}

//接口
type Call interface {
	call(int)
}

