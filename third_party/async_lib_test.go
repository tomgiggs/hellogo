package third_party

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sync"
	"testing"
)

func TestAsync_Parallel(t *testing.T) {

	async := Async{
		mu:        sync.Mutex{},
		wait:      sync.WaitGroup{},
		tasks:     make([]RunFunc,0),
		err:       make(chan error),
		paramMap:  make(map[string]interface{}),
		resultMap: make(map[string]interface{}),
		results:   make(chan interface{}),
		callback:  nil,
	}
	async.AddFunc(Input)
	async.AddFunc(Input2)
	async.AddCalalbackFunc(Output)
	async.Parallel()
}
func Input(ii interface{},output interface{}) error  {
	log.Info("input1 is:",ii)
	output = 20
	return nil
}

func Input2(ii interface{},output interface{}) error  {
	log.Info("input2 is:",ii)
	output = 30
	return nil
}

func Output(res map[string]interface{}) error{

	log.Info("res is: ",res)
	return nil
}



func parseResp(resp interface{}) {
	if resp == nil {
		return
	}
	fmt.Println("resp===========>", reflect.TypeOf(resp).Kind().String())
	switch reflect.TypeOf(resp).String() {
	case "[]interface {}", "slice":
		for _, val := range resp.([]interface{}) {
			if val == nil {
				continue
			}
			fmt.Printf("array val:%s, %+v\n", reflect.TypeOf(val).String(), val)
			parseResp(val)
		}
	case "map", "map[interface {}]interface {}":
		for k2, val2 := range resp.(map[interface{}]interface{}) {
			if val2 == nil || k2 == nil {
				continue
			}
			fmt.Printf("map key type: %s,val type:%s,val:%+v,%+v\n", reflect.TypeOf(k2).String(), reflect.TypeOf(val2).String(), k2, val2)
			parseResp(val2)
		}
	case "string":
		fmt.Println("string val is:", resp)
	default:
		fmt.Printf("unknow type:%s,%+v\n", reflect.TypeOf(resp), resp)
	}

}