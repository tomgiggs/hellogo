package basic_grammar

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func IoDemo() {
	//body:=ReadFile("./basic_grammar/map_demo.go")
	//fmt.Println(body)
	//GetFileList("./")
	ReadLine("./basic_grammar/map_demo.go")

}

func ReadFile(fineName string) string {
	f, err := os.Open(fineName)
	defer f.Close()
	if err != nil {
		fmt.Print("error::")
		fmt.Print(err)
		return ""
	}
	//全部一次性读取
	var body []byte
	body, _ = ioutil.ReadAll(f)
	return string(body)
	//----------------------------------------------
	//能否实现一个像Python那种yield的函数，每次要用的时候再读取？
	//readBuffer :=bufio.NewReader(f)
	//content := ""
	//for{
	//	line,err := readBuffer.ReadString('\n')
	//	if err!=nil ||io.EOF ==err{
	//		break//这个会出现最后一个空行没有输出的问题，把输出起到这个判断前面会解决这个问题
	//	}
	//	content+=line
	//	fmt.Println(line) //如果前面的读取没有被注释，这个就不会有输出，应该是指针移到最后面了
	//}
	//return content
}
func ReadLine(fineName string) interface{} {
	f, err := os.Open("./basic_grammar/map_demo.go")
	defer f.Close()
	if err != nil {
		fmt.Print("error::")
		fmt.Print(err)
		return ""
	}
	//能否实现一个像Python那种yield的函数，每次要用的时候再读取？
	readBuffer := bufio.NewReader(f)
	for {
		lines, _, err := readBuffer.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(lines))
	}

	return nil
}

func GetFileList(dirName string) []interface{} {
	fileList, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil
	}
	for i := 0; i < len(fileList); i++ {
		var fileInfo = fileList[i]
		filePath := dirName + string(os.PathSeparator) + fileInfo.Name()
		fmt.Println(filePath)
		if fileInfo.IsDir() {
			GetFileList(filePath)
			continue
		}
		fmt.Println(fileInfo.Name(), fileInfo.Size())

	}
	return nil
}

func WriteFileDemo() {

	file, err := os.OpenFile("test0001.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()
	csvWriter := bufio.NewWriter(file)
	csvWriter.WriteString("hello this is write by csv writer")
	csvWriter.Flush()
}

func Finfo() {
	fInfo, err := os.Stat("basic_grammar/flag_demo.go")
	if err != nil {
		fmt.Println("file not exist:", err)
		return
	}
	if fInfo.IsDir() {
		fmt.Println("input is a path")
		os.Exit(0)
	}

}
