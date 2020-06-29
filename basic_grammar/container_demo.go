package basic_grammar

import (
	"container/list"
	"fmt"
	"sort"
)

func Container01() {
	intList := []int{1, 4, 2, 3, 9, 0, 20}
	sort.Ints(intList)
	fmt.Println("sorted list:", intList)
	strList := []string{"aca", "bde", "name", "aac", "bbd"}
	sort.Strings(strList)
	fmt.Println("sorted string:", strList)
	//链表，可以实现很灵活的需求
	listDemo := list.New()
	listDemo.PushBack("20")
	listDemo.PushBack(202)
	fmt.Println("pop from listDemo:", listDemo.Front().Value) //取第一个元素。但不会移除，没有pop出栈功能
	//实现pop功能
	var next *list.Element
	for i := listDemo.Front(); i != nil; i = next { //用 for 语句进行遍历，其中 i:=listDemo.Front() 表示初始赋值，只会在一开始执行一次；每次循环会进行一次 i!=nil 语句判断，如果返回 false，表示退出循环，反之则会执行 i=next
		next = i.Next() //如果没有next保存被删除元素的指针信息，后面循环就无法继续
		fmt.Println(i.Value)
		listDemo.Remove(i)
	}
	fmt.Println("listDemo length is:", listDemo.Len())
}
