package basic_grammar

import (
	"fmt"
	"strings"
)

func BasicString(){
	str :='树'//单引号只能用于声明一个字符(rune)，会被转成byte值，双引号可以很长的字符同时做转义处理，反引号可以原样输出不转义
	fmt.Println(str)
	//在 strings 包中 Index 函数可以返回指定字符或字符串的第一个字符的索引值，如果不存在则返回 -1，如果处理包含多个字节组成的字符的字符串，需要使用 IndexRune 函数来对字符进行定位
	fmt.Println(strings.IndexRune("好树结好果", '树'))//3
	fmt.Println(strings.Index("好树结好果", "树"))//3，为什么是3呢
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
}