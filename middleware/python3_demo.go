package middleware

import (
	"fmt"
	"github.com/DataDog/go-python3"
)

func ImportGlobalModule() {
	sysModule := python3.PyImport_ImportModule("sys")                                                         // import sys
	path := sysModule.GetAttrString("path")                                                                   // path = sys.path
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString("D:\\program\\Python37\\Lib\\site-packages")) // path.insert(0, dir) 将路径添加到python 包搜索路径里面
}

func ImportModule(dir, name string) *python3.PyObject {
	sysModule := python3.PyImport_ImportModule("sys")                 // import sys
	path := sysModule.GetAttrString("path")                           // path = sys.path
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir)) // path.insert(0, dir)
	return python3.PyImport_ImportModule(name)                        // return __import__(name)
}

func RunPython3Code() {
	// 初始化go-python
	fmt.Println("init ...")
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	osMod := python3.PyImport_ImportModule("os")

	if osMod == nil {
		fmt.Println("could not import 'os'")
		return
	}
	// osMod := python3.PyImport_AddModule("os")
	// defer osMod.DecRef()
	fmt.Println("find os mod")
	// listdir := osMod.GetAttrString("listdir") //获取 listdir 函数
	// if listdir == nil {
	// 	fmt.Println("could not find 'os.listdir'")
	// 	return
	// }
	flist := osMod.CallMethodArgs("listdir", python3.PyUnicode_FromString("./"))
	if flist == nil {
		return
	}
	info, err := pythonRepr(flist)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(info)
}

func Python3Demo() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	python3.PyRun_SimpleString("print('hello from python in go')")
	ImportGlobalModule() // 导入全局依赖包

	res := ImportModule("./pycode", "test")
	helloFunc := res.GetAttrString("main")
	if helloFunc == nil {
		return
	}
	fmt.Println("hellofunc is not nil")
	bArgs := python3.PyTuple_New(1)
	python3.PyTuple_SetItem(bArgs, 0, python3.PyUnicode_FromString("world"))
	// 调用函数
	result := helloFunc.Call(bArgs, python3.Py_None)
	// result := helloFunc.Call(python3.PyUnicode_FromString("jobs"), python3.Py_None) // 无效
	// result := helloFunc.Call(python3.Py_None, python3.Py_None) // 无效
	if result == nil {
		return
	}
	fmt.Println("call func success", result)
	// 列表测试
	fmt.Println(python3.PyUnicode_AsUTF8(result))
	err := printList()
	if err != nil {
		fmt.Printf("failed to print the python list: %s\n", err)
	}
}

func printList() error {
	list := python3.PyList_New(5)
	python3.PyList_SetItem(list, 0, python3.PyLong_FromGoInt(20))

	if exc := python3.PyErr_Occurred(); list == nil && exc != nil {
		return fmt.Errorf("Fail to create python list object")
	}
	defer list.DecRef()

	repr, err := pythonRepr(list)
	if err != nil {
		return fmt.Errorf("fail to get representation of object list")
	}
	fmt.Printf("python list: %s\n", repr)

	return nil
}

func pythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}

	s := o.Repr()
	if s == nil {
		python3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer s.DecRef()

	return python3.PyUnicode_AsUTF8(s), nil
}
