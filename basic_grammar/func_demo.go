package basic_grammar

import (
	"fmt"
	"reflect"
)

func Logic() {
	//s:= "raw input"
	//simple(s)
	//fmt.Println(s)//raw input
	//simple2(&s)
	//fmt.Println(s)//changed by func 从这里可以看到如果是简单的传递一个值是不会改变原来的值的，如果是传递的一个指针，那就可以修改原来的值
	//funcWithUnknowArgs("arg1", "arg2", "arg3", "arg4")
	//fmt.Println(funcWithOneReturn())
	//fmt.Println(funcWithMutliReturn())
	//fmt.Println(lazyFunc2(2, 3, 4, "hello", "i", "am", "tom", "how", "are", "you"))
	//fmt.Println(lazyFunc(3, 6, 6))
	//fmt.Println(veryLazyFunc(5, 77, 2))
	//funcWithUnknowType("2000")
	//f := funcWithOneReturn //函数也可以是一个类型，可以赋值给一个变量
	//fmt.Println(f())
	//m := closure(20)
	//fmt.Println(m(5))
	//-------------
	//deferDemo(40)
	input := 40
	deferDemo2(&input)
	//tryCatch()
}
func simple(a string) {
	a = "changed by func"
}
func simple2(a *string) {
	*a = "changed by func"
}

func funcWithOneReturn() string {
	return " return from func!!"
}
func funcWithMutliReturn() (string, int) {
	return "goood", 20
}

//带有不知道多少个参数的函数
func funcWithUnknowArgs(kw ...string) {
	fmt.Println(kw) //kw是一个切片
	args := kw[:]
	fmt.Print("args are:", args)

}
func lazyFunc(a, b, c int) int { //参数类型是相同的时候就可以省去每个参数都声明类型
	return a + b + c
}
func lazyFunc2(a, b, c int, d, e, f string, unknow ...string) (int, string) { //参数类型是相同的时候就可以省去每个参数都声明类型

	return a + b + c, d + " " + e + " " + f + " " + StringJoin(" ", unknow)
}
func veryLazyFunc(a, b, c int) (sum, min, max int) { //可以这里声明返回值的变量名称，这样就可以更加清楚的知道返回值是什么意思了，如果是直接使用int,int,int会不容易知道返回值是什么
	return a + b + c, Max(a, b, c), Min(a, b, c)
}
func Max(args ...int) int {
	var max = args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}
func Min(args ...int) int {
	var min = args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}
	return min
}
func StringJoin(sep string, words []string) string {
	var str = ""
	for _, word := range words {
		str += word + sep
	}
	return str[:len(str)-1]
}

//泛型函数
func funcWithUnknowType(kw interface{}) interface{} {
	fmt.Println("input is:", kw)
	fmt.Printf("the type of input is: %s \n", reflect.TypeOf(kw))
	return 2000
}

//闭包函数，类似于Python的装饰器
func closure(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

//try-catch
func tryCatch() {
	//使用defer来捕获异常使程序继续运行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from panic")
		}
	}()
	panic("error happend")

}

//defer demo
func deferDemo(loop int) {
	fmt.Printf("input before loop is:%d", loop)
	defer fmt.Printf("input by defer before exit is: %d", loop)
	//for i:=20;i<loop;i++{//这是一个死循环，loop一直在变大//input after loop is:%d -9223372034707292310怎么会变得这么大呢？
	//	loop+=i
	//}
	i := 20
	for i < 30 {
		i++
		loop += i
	}

	fmt.Println("input after loop is:%d", loop)
}

func deferDemo2(loop *int) {
	fmt.Printf("input before loop is:%d\n", *loop)
	defer fmt.Printf("input before exit is:%d\n", *loop)
	for i := 20; i < 30; i++ {
		defer fmt.Println("print from simple defer", i)
		//defer func() {
		//	fmt.Println("print from closure defer", i) //可以看到这个每次都是打印30，而上面的defer每次打印都不一样，
		//	// 这是因为第二个函数使用了闭包函数，而闭包函数使用了引用而不是传值，而第一个defer是直接传值进去的，压到栈里面的只是代码和内存地址
		//}()
		defer func(i int) {
			fmt.Println("print from closure defer", i) //这里打出来的i就是不同的了，因为直接将值传递到函数中同时压栈了
		}(i)
		*loop += i
	}
	fmt.Printf("input after loop is:%d\n", *loop)
}

func Test002(a int) (int, error) {
	//var bb error
	//bb.Error()
	return a, nil
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func DeferDemo03(){
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	fmt.Println("=================")
}
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
func OverWriteDemo01(){
	t :=Teacher{}
	t.ShowA()
	t.ShowB()
	//----
	//showA
	//showB
	//teacher showB
}
type student struct {
	Name string
	Age  int
}
func StudentLoop01() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		m[stu.Name] = &stus[i]
		//fmt.Printf("%p\n",stu)
	}
	fmt.Printf("%v",m)
}