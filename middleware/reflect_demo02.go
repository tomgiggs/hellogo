package middleware

import (
	"fmt"
	"reflect"
)

type Struct01 struct {
	A int
}
type Struct02 struct {
	B int
}

type ConTainer struct {
	types map[string]reflect.Type
	names map[reflect.Type]string
}
func (c *ConTainer)Register(t interface{}) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	name := rt.PkgPath() + "/" + rt.Name()
	c.types[name] = rt
	c.names[rt] = name
}

func GenObj(){
	var container ConTainer
	container = ConTainer{
		types:make(map[string]reflect.Type),
		names:make(map[reflect.Type]string),
	}
	container.Register(Struct01{})
	container.Register(Struct02{})
	//var body interface{}
	if t, exists := container.types["hellogo/middleware/Struct01"]; exists {
		body := reflect.New(t)
		fieldValue := body.Elem().FieldByName("A")
		fieldValue.Set(reflect.ValueOf(200))
		fmt.Println(body.Elem())

		//无效代码
		//valueInfo := reflect.ValueOf(&body)
		//elem := valueInfo.Elem()
		//fieldValue := elem.FieldByName("A")
		//fmt.Println(fieldValue.Kind())
		//if fieldValue.Kind()==reflect.Int{
		//	*(*int)(unsafe.Pointer(fieldValue.Addr().Pointer())) = 200
		//}
		fmt.Println(body)
	}


}
