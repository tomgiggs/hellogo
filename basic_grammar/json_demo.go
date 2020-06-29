package basic_grammar

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)
import "github.com/bitly/go-simplejson"

var jsonStr = `
       {
           "person": [{
              "name": "piao",
              "age": 30,
              "email": "piaoyunsoft@163.com",
              "phoneNum": [
                  "13974999999",
                  "13984999999"
              ]
           }, {
              "name": "aaaaa",
              "age": 20,
              "email": "aaaaaa@163.com",
              "phoneNum": [
                  "13974998888",
                  "13984998888"
              ]
           }, {
              "name": "bbbbbb",
              "age": 10,
              "email": "bbbbbb@163.com",
              "phoneNum": [
                  "13974997777",
                  "13984997777"
              ]
           }]
       }
       `

func SimpleJsonDemo() {
	js, err := simplejson.NewJson([]byte(jsonStr))

	fmt.Println(path.Base(filepath.ToSlash(`./log\Activity-2020-06-03.log`)))
	return
	if err != nil {
		panic(err.Error())
	}
	_, ok := js.CheckGet("person")
	if !ok {
		fmt.Println("key not exist")
	}
	for k, v := range js.MustArray() {
		fmt.Println(k, v)
	}

	personArr, err := js.Get("person").Array()
	fmt.Println(len(personArr))

	// 遍历
	for i, _ := range personArr {
		//fmt.Println(i, v)
		person := js.Get("person").GetIndex(i)
		name := person.Get("name").MustString()
		age := person.Get("age").MustInt()
		email := person.Get("email").MustString()

		fmt.Printf("name=%s, age=%d, email=%s\n", name, age, email)

		// 读取手机号
		phoneNumArr, _ := person.Get("phoneNum").Array()
		for ii, vv := range phoneNumArr {
			fmt.Println(ii, vv)

		}
	}

}

func Camel2underscore(name string) string {
	newStr := ""
	for i, c := range name {
		if int(c) >= 65 && int(c) <= 90 {
			if i != 0 {
				newStr += "_"
			}
		}
		newStr += strings.ToLower((string(c)))
	}
	fmt.Println(newStr)
	return newStr
}
