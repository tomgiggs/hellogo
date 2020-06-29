package basic_grammar

import "fmt"

//var str string
//str :=""""
//var str =""""

func ArrayDemo() {
	var userNames = [7]string{"tom", "giggs", "jerry", "james", "kotlin"} //创建一个数组，数组是长度固定的，后面是无法改变的，初始化时也无法超过指定长度
	unknowUsers := [...]int{9: 1}                                         //创建长度为9，默认值为0的数组，最后一个值为1的数组，string类型同理，最后一个值是指定值，默认值全是空串
	fmt.Println(unknowUsers)
	//userNames = append(userNames, "case")//这个会报错
	fmt.Println(userNames)
	subProductNmae := userNames[2:4]
	fmt.Println(subProductNmae, len(subProductNmae), cap(subProductNmae)) //切片长度和容量是一样的。。。，说好的切片长度是数组长度的呢？原来
	// 切片的长度是指切片中元素的个数。切片的容量是指从切片的起始元素开始到其底层数组中的最后一个元素的个数，当我第一次设置userNames 长度为5时两者就是一样的
	var productNames []string
	productNames = append(productNames, "phone", "computer", "keyboard", "mouse", "mouse", "pad") //这个跟Python有点不一样，Python是直接product.append进行添加成员
	fmt.Println(productNames, len(productNames), cap(productNames))
	//productNames = append(productNames,userNames)//将一个数组追加到一个切片后面也是不行的，会报错，那么怎么将这两个合并呢？
	someProduct := make([]string, 5, 20)
	someProduct = append(someProduct, "test001", "test002", "test003", "test004")
	fmt.Println(someProduct)                               //从输出结果可以看出使用make函数会创造一个具有默认值的切片
	foods := []string{"rice", "chicken", "noodle", "meat"} //初始化同时赋值
	fmt.Println(foods[3:])
	fmt.Println("hello this is array demo!")

	sum := float64(0)
	nums := [...]float64{12, 3, 5, 7, 34453, 43}
	for index, value := range nums {
		sum += value
		println(index)
	}
	fmt.Printf("the sum is:%f \n", sum)
	//如何合并两个切片呢？
	merged_slice := append(productNames, foods...) //append第二个参数不接受一个切片，但是这样是可以的，很神奇。。。这样操作的前提是两个切片类型要一样不然会报错
	fmt.Println("before passed:", merged_slice)
	//fmt.Printf("men addr is:%p \n",merged_slice)
	SliceDemo(merged_slice)
	fmt.Println("after pass to sliceDemo:", merged_slice)
	//go传参是使用切片来实现的,切片的复制引用是一个问题
	sliceDemo2(&merged_slice) //从结果来看这是一个引用传递，因为值被修改了
	fmt.Println("after pass to sliceDemo2:", merged_slice)
	copy(productNames, foods) //使用copy函数对两个切片进行操作,这个会将第一个参数和第二个参数位置相同的元素全部替换为第二个参数的元素
	fmt.Println(userNames)    //[tom giggs jerry james kotlin  ]
	arrayFunc(userNames)
	fmt.Println(userNames)            //[tom giggs jerry james kotlin  ] 使用数组进行传参时要求函数的参数跟要传进去的参数一样，这里的一样是指类型长度要一样
	fmt.Println(arrayFunc(userNames)) //[tom giggs goood james kotlin  ]//从结果可以看出传递数组是不会导致数组值被改变的
}
func SliceDemo(names []string) {
	//fmt.Println(names)
	names[2] = "new_name"
	//fmt.Printf("men addr in func is:%p \n",names)
}

func sliceDemo2(names *[]string) {
	//fmt.Println(names)
	for index, _ := range *names {
		(*names)[index] = "2222"
	}
}

func arrayFunc(names [7]string) [7]string {
	names[2] = "goood"
	return names
}

func SliceDemo3() {
	score := make([]int, 4, 10) //第一个参数是类型，第二个是长度，第三个是容量
	fmt.Printf("append before: %v,length is: %d ,cap is:%d \n", score, len(score), cap(score))
	score = append(score, 200)
	score[3] = 200
	fmt.Printf("score after: %v,length is: %d,cap is:%d\n", score, len(score), cap(score))
	b := []int{1, 2, 3}
	fmt.Printf("array cap=%d,len=%d\n", cap(b), len(b))
	ch := make(chan int, 2)
	fmt.Printf("channel cap=%d,len=%d\n", cap(ch), len(ch))
	m := make(map[int]string, 10)
	m[1] = "string demo"
	fmt.Printf("map no cap has len=%d\n", len(m))

}

func SliceDemo4() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s2 := s[2:]
	s3 := s[1:]
	//s3 = append(s3, 20, 40, 50)
	s3[2] = 20 //使用同一个切片切出来的切片，其中一个修改了值会影响其他切片的值，除非修改值的切片独立出去了
	printSlice(s)
	printSlice(s2)
	printSlice(s3)
	// 拓展其长度
	fmt.Print("--------------") //内置函数print()会比fmt.Print早运行,这个很奇怪
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

	s = append(s, 20)
	printSlice(s)
	s = append(s, 300, 1, 2)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v ptr=%p \n", len(s), cap(s), s, s)
}
func SliceDemo5() {
	arr := [7]int{1, 2, 3, 4, 5}
	slice01 := arr[0:2]
	fmt.Println(slice01)
	fmt.Println(arr[:2])
	slice01 = append(slice01, 6, 7, 8)
	slice01 = append(slice01, 22, 33)
	fmt.Println(arr)
	//b:=(1==2)
	b := true
	fmt.Println(b)
	var special = struct {
		int //匿名字段可以实现继承
		m   map[string]string
	}{m: make(map[string]string)}
	special.m["name"] = "nice"
	special.int = 10
	fmt.Println(special.int)

}

func SliceDemo06() {
	a := []int{22, 33, 44, 55}
	a = append(a[:1], a[2:]...) //移除指定下标的元素
	fmt.Println(a)
}

type Student struct {
	Name string
	Age  int
}

var stus = []Student{
	{Name: "chen", Age: 20},
	{Name: "yi", Age: 21},
	{Name: "xun", Age: 22},
}

func SliceDemo07() {
	fmt.Println(PaseStudent())
}

func PaseStudent() map[string]*Student {
	m := make(map[string]*Student)
	for i, _ := range stus {
		stu := stus[i]
		m[stu.Name] = &stu
	}
	//for _, stu := range stus {
	//	stu1 := stu
	//	m[stu.Name] = &stu1
	//}
	b := funcDemo([]int{2, 3, 4}...)
	fmt.Println(b)

	return m

}

func funcDemo(a ...int) int {
	sum := 0
	for _, a := range a {
		sum += a
	}
	return sum
}
