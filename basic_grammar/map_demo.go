package basic_grammar
/*
键值对数据结构相关操作
 */

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type UserInfo struct{
	Name string `json:"user_name"`
	Age int`json:"age"`
	Location string `json:"location"`
	Job string `json:"job"`

}

func MapDemo()  {
	userMap :=map[string]string{//这里必须指定value的类型，如果我想在让value有很多种类型怎么办？如果想要使用类似json的数据结构就需要自己定义一个数据类型，然后使用，而不是使用map来存
		"name":"cyl",
		"age":"27",
		"pop":"yes",
	}
	fmt.Println(userMap)
	jsonMap,_ := json.Marshal(&userMap)
	fmt.Println(string(jsonMap)) //可以取map的长度
	delete(userMap,"pop")//删除一个键值
	for k,v :=range userMap{
		fmt.Printf("key is: %s,value is: %s \n",k,v)
	}
	fmt.Printf("map key num is: %d \n",len(userMap))
	userMap["name"]="no one"
	fmt.Printf("user name is: %s \n",userMap["name"])

	//判断一个键值对是否存在
	noExist,v:=userMap["unexist"]
	fmt.Printf("key exist:%t \n",v)
	fmt.Println("key does not exist:",noExist)
	if noExist==""{
		fmt.Println("the key is not exist!!!")
	}
	newMap := new(map[string]string)
	json.Unmarshal(jsonMap,newMap)
	fmt.Println(newMap)
	//合并两个map，没有内置方法，只能手动写。。。
	userMapNew :=map[string]string{
		"school":"nouc",
	}
	for k,v := range userMap{
		userMapNew[k] = v
	}
	fmt.Println("new map is: ",userMapNew)

	//map排序
	var keys []string
	for k :=range userMap{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,k :=range keys{
		fmt.Println(userMap[k])
	}

	fmt.Println("-------------struct operation-----------------------")
	userStruct :=UserInfo{//有没有取地址好像都可以使用，有什么区别吗？有区别，如果取地址了就无法使用reflect获取成员数量了
		Name:"cyl",
		Job:"progammer",
		Age:26,
		Location:"fujian province,fuzhou city",
	}
	fmt.Println(userStruct)
	//使用反射获得成员
	t:=reflect.TypeOf(userStruct)
	fmt.Println(t.NumField())
	fmt.Println(t.Field(1).Name)
	fmt.Println(t.FieldByName("Age"))
	//fmt.Println(t.Method(0).Name)
	fmt.Print("struct contain method num: %d \n",t.NumMethod())
	for i:=0;i< t.NumField();i++{
		fmt.Printf("member:%s \n",t.Field(i).Name)
	}

	json_str,err :=json.Marshal(userStruct)
	if err!=nil{
		panic(err)
		return
	}

	fmt.Println(string(json_str)) //为什么输出为空呢？因为结构体里面定义键名称为小写无法访问所以输出为空，要指定输出的json格式需要额外指定json键名，如上面
	//json转struct
	revealed:= new(UserInfo)
	json.Unmarshal(json_str,revealed)
	fmt.Println(revealed.Age)
}