package middleware

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	//"github.com/layeh/gopher-luar"
	"layeh.com/gopher-luar"
	"log"
	//"github.com/fwhezfwhez/errorx"
)


func LuaDemo() {
	ls := lua.NewState()
	defer ls.Close()
	//if err := ls.DoString(`print("hello, this is run by golang")`); err != nil {
	//	panic("run lua code error")
	//}

	if err := ls.DoFile("basic_grammar/demo.lua"); err != nil {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		log.Println(err)
		//print(err.(errorx.Error).StackTrace())

		panic("run lua file error")
	}
	err := ls.CallByParam(lua.P{
		Fn:      ls.GetGlobal("luaWorld"), //指定调用函数名
		NRet:    1,                        //指定返回值数量
		Protect: true,                     //出现异常是panic还是返回err
	}, lua.LNumber(20)) //传递入参
	if err != nil {
		panic("call func error")
	}
	ret := ls.Get(-1) //获取返回值
	ls.Pop(1)
	res, ok := ret.(lua.LNumber)
	if ok {
		fmt.Println("get res from lua:", res)
	}
}

type LuaZygote struct{
	ls       *lua.LState
}

func(lz *LuaZygote) Add(a,b int) int{
	fmt.Println("add func called,result is: ",a+b)
	return a+b
}

func LuaDemo02(){
	lz := new(LuaZygote)
	lz.ls = lua.NewState()
	lz.ls.SetGlobal("Add", luar.New(lz.ls, lz.Add))// 将方法暴露给lua脚本调用
	if err := lz.ls.DoString(`print(Add(2,3))`); err != nil {
		fmt.Println("run lua code error: ",err.Error())
	}
}