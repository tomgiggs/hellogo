package basic_grammar

import "fmt"

func ArrayDemo(){
	var userNames = [7]string{"tom","giggs","jerry","james","kotlin"}//创建一个数组，数组是长度固定的，后面是无法改变的，初始化时也无法超过指定长度
	unknowUsers := [...]int{9:1}//创建长度为9，默认值为0的数组，最后一个值为1的数组，string类型同理，最后一个值是指定值，默认值全是空串
	fmt.Println(unknowUsers)
	//userNames = append(userNames, "case")//这个会报错
	fmt.Println(userNames)
	subProductNmae := userNames[2:4]
	fmt.Println(subProductNmae,len(subProductNmae),cap(subProductNmae)) //切片长度和容量是一样的。。。，说好的切片长度是数组长度的呢？原来
	// 切片的长度是指切片中元素的个数。切片的容量是指从切片的起始元素开始到其底层数组中的最后一个元素的个数，当我第一次设置userNames 长度为5时两者就是一样的
	 var productNames []string
	 productNames = append(productNames,"phone","computer","keyboard","mouse","mouse","pad")//这个跟Python有点不一样，Python是直接product.append进行添加成员
	fmt.Println(productNames,len(productNames),cap(productNames))
	 //productNames = append(productNames,userNames)//将一个数组追加到一个切片后面也是不行的，会报错，那么怎么将这两个合并呢？
	someProduct := make([]string,5,20)
	someProduct = append(someProduct,"test001","test002","test003","test004")
	fmt.Println(someProduct)//从输出结果可以看出使用make函数会创造一个具有默认值的切片
	 foods := []string{"rice","chicken","noodle","meat"}//初始化同时赋值
	fmt.Println(foods[3:])
	fmt.Println("hello this is array demo!")

	 sum :=float64(0)
	 nums := [...]float64{12,3,5,7,34453,43}
	 for index,value:= range nums{
	 	sum+=value
	 	println(index)
	 }
	fmt.Printf("the sum is:%f \n", sum)
	//如何合并两个切片呢？
	merged_slice := append(productNames,foods ...)//append第二个参数不接受一个切片，但是这样是可以的，很神奇。。。这样操作的前提是两个切片类型要一样不然会报错
	fmt.Println("before passed:",merged_slice)
	//fmt.Printf("men addr is:%p \n",merged_slice)
	sliceDemo(merged_slice)
	fmt.Println("after pass to sliceDemo:",merged_slice)
	//go传参是使用切片来实现的,切片的复制引用是一个问题
	sliceDemo2(&merged_slice) //从结果来看这是一个引用传递，因为值被修改了
	fmt.Println("after pass to sliceDemo2:",merged_slice)
	copy(productNames,foods)//使用copy函数对两个切片进行操作,这个会将第一个参数和第二个参数位置相同的元素全部替换为第二个参数的元素
	fmt.Println(userNames)//[tom giggs jerry james kotlin  ]
	arrayFunc(userNames)
	fmt.Println(userNames)//[tom giggs jerry james kotlin  ] 使用数组进行传参时要求函数的参数跟要传进去的参数一样，这里的一样是指类型长度要一样
	fmt.Println(arrayFunc(userNames))//[tom giggs goood james kotlin  ]//从结果可以看出传递数组是不会导致数组值被改变的
}
func sliceDemo(names []string){
	//fmt.Println(names)
	names[2] = "new_name"
	//fmt.Printf("men addr in func is:%p \n",names)
}

func sliceDemo2(names *[]string){
	//fmt.Println(names)
	for index,_ :=range *names{
		(*names)[index] = "2222"
	}
}

func arrayFunc(names [7]string) [7]string{
	names[2] = "goood"
	return names
}



