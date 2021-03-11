package basic_grammar

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

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
type TblUserMonitorTimes struct {
	Id                   int64
	UserId 				 int64
	IndexId				 int32
	Data1				 int32
	Data2				 int32
	Data3				 int32
	Data4				 int32
	Data5				 int32
	Data6				 int32
	Data7				 int32
	Data8				 int32
	Data9				 int32
	Data10				 int32
	Data11				 int32
	Data12				 int32
	Data13				 int32
	Data14				 int32
	Data15				 int32
	Data16				 int32
	Data17				 int32
	Data18				 int32
	Data19				 int32
	Data20				 int32
}

func ReflectDemo02(){
	dbData := TblUserMonitorTimes{}
	//dbData := new(TblUserMonitorTimes)
	//aaaa := 200
	//b := &aaaa
	DataAssemble(10,dbData)
	//Idemo(10,b)
}

//func DataAssemble(index int,dbData TblUserMonitorTimes){
func DataAssemble(index int,dbData interface{}){//改成这样就报错
	//dbData := TblUserMonitorTimes{}
	dbDataInfo := reflect.TypeOf(dbData)
	valueInfo := reflect.ValueOf(&dbData)
	//valueInfo := reflect.ValueOf(&dbData)
	elem := valueInfo.Elem()
	for i:=0;i<dbDataInfo.NumField();i++{
		fieldInfo:= dbDataInfo.Field(i)
		fmt.Println("filed info: ",fieldInfo.Name,fieldInfo.Type,fieldInfo.Index,fieldInfo.Tag)//这个index是严格按照定义顺序来的吗，打印出来的好像是这样的
		//fmt.Println("type of filedname is: ",reflect.TypeOf(fieldInfo.Name))
		if strings.HasSuffix(fieldInfo.Name, strconv.Itoa(index)){
			fmt.Println(elem.Field(0))
			fieldValue := elem.FieldByName("Data10")//如果参数定义为interface{}的话就会报错： panic: reflect: call of reflect.Value.FieldByName on interface Value
			if fieldValue.Kind()==reflect.Int32{
				*(*int32)(unsafe.Pointer(fieldValue.Addr().Pointer())) = 200
			}
		}
	}
	fmt.Printf("result is: %v",dbData)
}
func Idemo(aa int,bb *interface{}){//interface前面加*号，怎么调用都是失败的Cannot use 'b' (type *int) as type *interface{}
	fmt.Println("goooooood")
}

type Filter interface {
	ConstructPBFilter() (int, error)
}

type SingleColumnValueFilter struct {

}
func (f *SingleColumnValueFilter) ConstructPBFilter()(int, error){
	return 0,nil
}
//判断接口是否实现
func IsInterfaceValid(){
	var _ Filter = (*SingleColumnValueFilter)(nil)//new是编译的时候检查，这样写是运行的时候检查
	var _ Filter = new(SingleColumnValueFilter)
	var _ Filter = &SingleColumnValueFilter{}
}