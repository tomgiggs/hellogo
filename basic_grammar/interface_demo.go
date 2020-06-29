package basic_grammar

import "fmt"

type Talker interface {
	Talk() string
}

type Person struct {
}

func (p *Person) Talk() string {

	return "Hello"

}

func CheckNil(talker Talker) {

	if talker == nil {

		fmt.Println("talker is nil")

	} else {

		fmt.Println("talker is not nil")

	}

}
func Xxxxxxxx() {
	//CheckNil(nil)
	//var pPerson *Person = nil
	//CheckNil(pPerson)
	IfaceInput(200)

}

func IfaceInput(unknow interface{}) {
	age, ok := unknow.(int) //类型转换
	if !ok {
		fmt.Println("param unknow is not int")
		//return
	}
	fmt.Println(age)
}
//func forDemo001(){
//	var elm int
//	slice002 :=[]int{2,3,4,5,6,7}
//	for _,val :=range slice002; val!=nil;elm=val{
//		fmt.Println(val)
//	}
//}
