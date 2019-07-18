package basic_grammar



import (
	"encoding/json"
	"fmt"
)

type proj struct {
	name string
	user string
	description string

}
type iphone struct{
	name string
}

type phone interface {
	call()
	display()
	//music() string
}

//这个是什么写法？？？
func (phone_demo iphone) call(){
	fmt.Println("this is interface call method")
}
func (phone_demo iphone) display() {
	fmt.Println("this is interface display method")

}
func (phone_demo iphone) music() int {
	fmt.Println("this is interface display method")
	return 20
}


func Basic(){
	var onephone phone
	onephone = new(iphone)//开启interface中music的定义就会报错，没实现这个方法，什么问题呢，因为interface定义里面没有返回类型，
	onephone.call()
	//map测试
	var proj_info = map[string]string{
		"proj_name":"demooooooo",
		"proj_id":"0001",
		"proj_user":"giggs",
	}
	fmt.Println(len(proj_info))
	fmt.Println(proj_info) //怎么突然这个也报错了。。。
	json_str,err :=json.Marshal(proj_info)
	if err!=nil{
		panic(err)
		return
	}
	fmt.Println(string(json_str))
	for key,val:=range proj_info{
		fmt.Println(key,val)
	}
	//fmt.Println(len("string lenth is"))
	//xxx := *proj{name:"xxxx",user:"sssss",description:"sssss"}
	var xxx proj
	xxx.user = "xxxxxxxx"
	xxx.name = "nnnnnnn"
	xxx.description = "ddddddddd"
	fmt.Println(xxx)
	var projx [5] proj
	projx[0] = xxx //panic: runtime error: index out of range???
	fmt.Println(projx)

	var proj_list [5] int //加了数组大小才不会报错，不然一直报错panic: runtime error: index out of range,因为数组长度是不可变的，要使用Python中的“高级数组”应该使用切片slice
	//append(proj_list,20)
	proj_list[0] = 10
	proj_list[1] = 20
	//proj_list = append(proj_list,20,30,40)//只有切片才能这样做
	var slice_demo [] int
	slice_demo = make([]int,5,20)
	fmt.Println(len(slice_demo))//5
	fmt.Println(cap(slice_demo))//20
	slice_demo[0] = 2
	//slice_demo[10] = 10 //panic: runtime error: index out of range,长度5，指定了更大的容量也还是报错
	slice_demo = append(slice_demo,20,30)//这个操作并没有扩大切片的容量，只是改变了长度
	fmt.Println(slice_demo) //[2 0 0 0 0 20 30]
	fmt.Println(len(slice_demo)) //7
	fmt.Println(cap(slice_demo))//20
	for index,val:=range slice_demo{
		fmt.Println(index,val)
	}

	//-------------数组测试---------------
	//var array =[5]int{1,2,3,4,5}
	var array =[]int{1,2,3,4,5}
	fmt.Println(array)
	array[0] = 10
	fmt.Println(array)



}








