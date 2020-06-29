package basic_grammar

import (
	"fmt"
	"log"
	"runtime"
)

func Update() {
	for i := 0; i < 100; i++ {
		//这个并不能如期恢复
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 8*1024)
				len := runtime.Stack(buf, false)
				log.Println("recover panic;")
				log.Println("%v:%s", r, string(buf[:len]))
			}
		}()
		//............
		panic("recover 01")
	}
}

var ii int = 0

func Update2() {

	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 8*1024)
			len := runtime.Stack(buf, false)
			log.Println("recover panic;")
			log.Println("%v:%s", r, buf[:len])
			Update2()
		}
	}()

	for ii = 0; ii < 100; ii++ {
		//............
		fmt.Println("curr num is:", ii)
		panic("for loop")

	}
}

func Update3() {
	//goto MAINLOOP //goto不能再defer中使用
	//MAINLOOP:
	//fmt.Println("mi")
	for ii = 0; ii < 100; ii++ {
		//............
		//panic("for loop")
		fmt.Println("gooood")
		go Recovery()
	}

}
func Recovery() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 8*1024)
			len := runtime.Stack(buf, false)
			fmt.Println("recover panic;")
			fmt.Println("%v:%s", r, buf[:len])
		}
	}()
	panic("for loop")
}
