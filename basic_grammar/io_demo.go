package basic_grammar

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func IoDemo(){
	f,err :=os.Open("hello.go")
	if err !=nil{
		fmt.Print("error::")
		fmt.Print(err)
		return
	}

	//fileList,err := ioutil.ReadDir("./")
	//if err!=nil{
	//	return
	//}
	//for  i:=0; i<len(fileList);i++  {
	//	var fileInfo = fileList[i]
	//	if fileInfo.IsDir(){
	//		continue
	//	}
	//	fmt.Println(fileInfo.Name())
	//	fmt.Println(fileInfo.Size())
	//
	//}
	//os.File.Seek
	log.Println("this is a logger output")
	defer f.Close()
	//全部一次性读取
	//var body  []byte
	//body,_ = ioutil.ReadAll(f)
	//fmt.Println(string(body[:])) //byte转字符串

	//str = "this is a string demo"-------------------
	var str string
	str = "这是一个中文字符串"
	var data []byte = []byte(str)
	fmt.Println(data)
	print(len(data))
	print(string(data[:])) //这个输出会有点不一样，虚线后面的代码先执行了，printlnbody的代码夹在输出结果中间
	readBuffer :=bufio.NewReader(f)
	for{
		line,err := readBuffer.ReadString('\n')
		if err!=nil ||io.EOF ==err{
			break//这个会出现最后一个空行没有输出的问题，把输出起到这个判断前面会解决这个问题
		}
		fmt.Println(line) //如果前面的读取没有被注释，这个就不会有输出，应该是指针移到最后面了


	}
}