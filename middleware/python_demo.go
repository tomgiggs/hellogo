package middleware
//
//import (
//	"fmt"
//	"github.com/sbinet/go-python"
//)
//
//var (
//	PyStr = python.PyString_FromString
//	GoStr = python.PyString_AS_STRING
//)
//
///*
//go-python需要用到gcc，Windows上可以使用choco快速安装，命令：choco install mingw
//还需要添加python头文件：
//export CPLUS_INCLUDE_PATH=/usr/local/include/python2.7:$CPLUS_INCLUDE_PATH
//export C_INCLUDE_PATH=/usr/local/include/python2.7:$C_INCLUDE_PATH
//set C_INCLUDE_PATH=/usr/local/include/python2.7:$C_INCLUDE_PATH
//
//找不到python的头文件会报错：./go-python.h:4:10: fatal error: Python.h: No such file or directory
//
//-- 然后会遇到（python3.9会出现，python2.7不会） could not determine kind of name for C.METH_OLDARGS
//还需要添加动态链接库搜索路径：
//Linux下要设置这个变量 LD_LIBRARY_PATH=D:\program\Python27\libs
//set LIBRARY_PATH=D:\program\Python27\libs
//不然会出现这个问题：ld.exe: cannot find -lpython27，ld returned 1 exit status
//
//*/
//
////RunPythonCode 直接执行python代码
//func RunPythonCode() {
//	// 初始化go-python
//	err := python.Initialize()
//	if err != nil {
//		fmt.Errorf("init python runtime error,err: %v", err.Error())
//		return
//	}
//	gostr := "world"                           //定义goloang字符串
//	pystr := python.PyString_FromString(gostr) //将golang字符串专程python字符串
//	str := python.PyString_AsString(pystr)     //将python字符串，再转为golang字符串。
//	fmt.Printf("hello %s \n", str)
//
//	osMod := python.PyImport_ImportModule("os") //导入os模块
//	if osMod == nil {
//		fmt.Println("could not import 'os'")
//		return
//	}
//
//	listdir := osMod.GetAttrString("listdir") //获取 listdir 函数
//	if listdir == nil {
//		fmt.Println("could not find 'os.listdir'")
//		return
//	}
//	//res := listdir.Call(python.PyString_FromString("./"),python.Py_None)//这个调用没用
//	flist := listdir.CallFunction(python.PyString_FromString("./"))
//	fmt.Println(flist)
//}
//
////CallPythonCustomLib 调用自己写的python包
//func CallPythonCustomLib() {
//	err := python.Initialize()
//	if err != nil {
//		fmt.Errorf("init python runtime error,err: %v", err.Error())
//		return
//	}
//	//demo := ImportModule("D:\\workspace\\hellogo", "requests")//如果python环境没有安装对应的依赖包，就会包找不到模块
//	//demo := ImportModule("D:\\workspace\\hellogo", "pycode") // 为什么这种方式不行了呢，因为少了pkg-config这个二进制，在这里可以找到https://www.cnblogs.com/qing123/p/12893111.html
//	demo := ImportModule("./pycode","demo")
//	// 也可以在这安装 choco install pkgconfiglite
//	//demo:= python.PyImport_ImportModule("demo")
//	if demo == nil{
//		fmt.Println("module not find")
//		return
//	}
//	fmt.Printf("[MODULE] repr(demo) = %s\n", GoStr(demo.Repr()))
//
//	// 获取变量
//
//	//获取函数
//	helloFunc := demo.GetAttrString("hello")
//	fmt.Printf("[FUNC] hello = %#v\n", helloFunc)
//
//	//构造函数入参
//	bArgs := python.PyTuple_New(1)
//	python.PyTuple_SetItem(bArgs, 0, PyStr("world"))
//	// 调用函数
//	res := helloFunc.Call(bArgs, python.Py_None)
//	fmt.Printf("[CALL] helloFunc('world') = %s\n", GoStr(res))
//
//	// sklearn
//	sklearn := demo.GetAttrString("sklearn")
//	skVersion := sklearn.GetAttrString("__version__")
//	fmt.Printf("[IMPORT] sklearn = %s\n", GoStr(sklearn.Repr()))
//	fmt.Printf("[IMPORT] sklearn version =  %s\n", GoStr(skVersion.Repr()))
//
//}
//func PythonRequestDemo(){
//	err := python.Initialize()
//	if err != nil {
//		fmt.Errorf("init python runtime error,err: %v", err.Error())
//		return
//	}
//	request := ImportModule("D:\\workspace\\hellogo", "requests") //文件夹里面需要放一个__init__.py文件，这说明这个是一个包
//	//request:= python.PyImport_ImportModule("request")
//	if request == nil{
//		fmt.Println("module not find")
//		return
//	}
//	fmt.Printf("[MODULE] repr(request) = %s\n", GoStr(request.Repr()))
//
//	a := request.GetAttrString("request")
//	fmt.Printf("[VARS] a = %#v\n", python.PyInt_AsLong(a))
//	reqArgs := python.PyTuple_New(2)
//	python.PyTuple_SetItem(reqArgs,0,PyStr("get"))
//	python.PyTuple_SetItem(reqArgs,1,PyStr("http://www.baidu.com"))
//	resp := a.Call(reqArgs,python.Py_None)
//	fmt.Println(GoStr(resp.GetAttrString("content")))
//}
//
//// ImportModule will import python module from given directory
//func ImportModule(dir, name string) *python.PyObject {
//	sysModule := python.PyImport_ImportModule("sys") // import sys
//	path := sysModule.GetAttrString("path")          // path = sys.path
//	python.PyList_Insert(path, 0, PyStr("D:\\program\\Python27\\Lib\\site-packages"))        // path.insert(0, dir) 将路径添加到python 包搜索路径里面
//	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
//	fmt.Println(GoStr(path.Repr()))
//	return python.PyImport_ImportModule(name)        // return __import__(name)
//}
