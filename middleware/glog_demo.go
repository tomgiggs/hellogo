package middleware

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/jeanphorn/log4go"
	"strconv"
)

//import "github.com/bitly/go-simplejson"

func GlogDemo() {
	flag.Parse()
	//glog.MaxSize=100000

	defer glog.Flush()

	glog.Info("hello this is write by glog")
}

func ParseWithExactType(datas []string, typeInfo []string) []interface{} {
	var newData = make([]interface{}, 0)
	//typeInfo := []string{"string","string","string","string","int64","string"}
	for idx, data := range datas {
		var tmp interface{}
		switch typeInfo[idx] {
		case "int64":
			tmp, _ = strconv.ParseInt(data, 10, 64)
		case "float64":
			tmp, _ = strconv.ParseFloat(data, 10)
		case "bool":
			tmp, _ = strconv.ParseBool(data)
		case "uint64":
			tmp, _ = strconv.ParseUint(data, 10, 64)
		default:
			tmp = data
		}
		newData = append(newData, tmp)
	}
	return newData

}

func TypeConvert() {
	data := []string{"name", "2000", "200", "20.1"}
	typeConfig := []string{"string", "int64", "int64", "float64"}
	target := ParseWithExactType(data, typeConfig)
	//simplejson.New()
	jsonStr, _ := json.Marshal(target)
	fmt.Printf("%s", string(jsonStr))
	//	if true {
	//		goto LABEL01
	//	}
	//中间不能隔着其他代码，不然容易报错
	//LABEL01:
	//	{
	//		fmt.Println("sssssssssss")
	//		fmt.Println("sdwsedw")
	//	}

}

func Log4goDemo01(){
	defer func() {
		log4go.Close()
	}()
	log4go.LoadConfiguration("./middleware/log4go.json")
	log4go.LOGGER("Test").Info("this is a output demo")
}