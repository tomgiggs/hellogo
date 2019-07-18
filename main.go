package main

import "hellogo/basic_grammar"

func main(){
	//运行报错： use of internal package github.com/go-redis/redis/internal/hashtag not allowed，这个是因为src文件夹底下有多个重名文件夹
	//basic_grammar.GetSysInfo()
//basic_grammar.ArrayDemo()
basic_grammar.Logic()

	//basic_grammar.MapDemo()
	//fmt.Println("hello")
}

