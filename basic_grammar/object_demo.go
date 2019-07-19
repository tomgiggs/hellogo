package basic_grammar

import (
	"fmt"
	"time"
)

//golang使用interface来实现接近于面向对象语言中的接口效果，使用struct来实现近似于类的效果，定义一个interface很简单跟其他语言很像，但是怎么实现它呢？
type Animal interface {
	Say() string
	Eat(foods []string)
	Sleep(sleepTime int)
	//LegNum int interface不能定义成员，只能定义函数
}

//再声明一个struct，这个struct将被用于当做匿名函数的参数传递给interface中定义的函数
type Dog struct {
	Name string
	Price float32
}

//下面的函数形式跟之前的普通函数有点不一样就是函数名前面还有一个括号传入一个结构体，这个跟Python中类的每个函数第一个参数是self有点类似，用于将函数绑定到struct上面，golang中认为一个struct绑定了接口的所有函数就默认是实现了接口
func (d *Dog)Say() string  {//这个绑定可以是一个指针，也可以是直接传一个struct，接收者为数据类型的方法称为值方法，接收者为指针类型的方法称之为指针方法。
	fmt.Printf("wang wang!! I am %s\n",d.Name)
	return "wangwang!"
}
func (d *Dog) Eat(foods []string){
	fmt.Println("I eat "+StringJoin(",",foods))
}

func (d *Dog)Sleep(s int){
	fmt.Printf("sleeping %d second... \n",s)
	wait:= time.Duration(s)*time.Second
	time.Sleep(wait)
	fmt.Println("awaken")
}
func (d *Dog) Sell() float32{
	return d.Price
}

func Create(s Dog){
	s.Say()
	foods :=[]string{"meat","rice","fruit"}
	s.Eat(foods)
	s.Sleep(2)
	fmt.Printf("my price is:%f ￥",s.Sell())
}

