package middleware

import (
	"fmt"
	"github.com/sbinet/go-python"
)

func PythonDemo(){
	// 初始化go-python
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	gostr := "foo"  //定义goloang字符串
	pystr := python.PyString_FromString(gostr)  //将golang字符串专程python字符串
	str := python.PyString_AsString(pystr)     //将python字符串，再转为golang字符串。
	fmt.Println("hello [", str, "]")

	pickle := python.PyImport_ImportModule("os")  //导入os模块
	if pickle == nil {
		panic("could not import 'os'")
	}

	listdir := pickle.GetAttrString("listdir")   //获取 listdir 函数
	if listdir == nil {
		panic("could not retrieve 'os.listdir'")
	}
	//res := listdir.Call(python.PyString_FromString("./"),python.Py_None)//这个调用没用
	res := listdir.CallFunction(python.PyString_FromString("./"))
	fmt.Println(res)
}
var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

func PythonDemo2(){
	// 初始化go-python
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	//InsertBeforeSysPath("D:\\program\\python27\\Lib\\site-packages")
	//InsertBeforeSysPath("/usr/local/lib/python2.7/dist-packages/site-packages")
	demo := ImportModule("./middleware", "demo")
	fmt.Printf("[MODULE] repr(hello) = %s\n", GoStr(demo.Repr()))

	a := demo.GetAttrString("a")
	fmt.Printf("[VARS] a = %#v\n", python.PyInt_AsLong(a))

	hello := demo.GetAttrString("hello")
	fmt.Printf("[FUNC] b = %#v\n", hello)
	bArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(bArgs, 0, PyStr("xixi"))
	res := hello.Call(bArgs, python.Py_None)
	fmt.Printf("[CALL] hello('xixi') = %s\n", GoStr(res))

	// sklearn
	sklearn := demo.GetAttrString("sklearn")
	skVersion := sklearn.GetAttrString("__version__")
	fmt.Printf("[IMPORT] sklearn = %s\n", GoStr(sklearn.Repr()))
	fmt.Printf("[IMPORT] sklearn version =  %s\n", GoStr(skVersion.Repr()))

}
// InsertBeforeSysPath will add given dir to python import path
func InsertBeforeSysPath(p string) string {
	sysModule := python.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	python.PyList_Insert(path, 0, PyStr(p))
	return GoStr(path.Repr())
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")                    // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)
	return python.PyImport_ImportModule(name)            // return __import__(name)
}