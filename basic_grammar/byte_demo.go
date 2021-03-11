package basic_grammar

import (
	"fmt"
	"unsafe"
)

func ByteDemo01(){
	//xx := DemoX{}
	//xx.ReadUint32()
	//-----------
	aa := make(map[int]int,0)
	MapParam01(aa)
	MapParam02(&aa)
	fmt.Println(aa)// map[3:10 4:10]
	//-----------
	//bb := []int{1,2,3,4,5}
	//SliceParam(&bb)
	//fmt.Println(bb)//值已经被函数改变，输出：[1 2 10 4 5]
	//-----------
	//basic := 20
	//BasicParam01(basic)
	//fmt.Println(basic)//不会被改变
	//BasicParam(&basic)
	//fmt.Println(basic)//被改变了
}
type DemoX struct {

}
func(x *DemoX) ReadUint32(){
	//b :=[]byte("4112")
	b := [4]byte{}
	//b[0] = uint8(4)
	//b[1] = uint8(1)
	//b[2] = uint8(1)
	//b[3] = uint8(2)
	fmt.Println(b)
	fmt.Printf("%p\n",&b)
	fmt.Println(len(b[:]))
	a := b[:]
	fmt.Printf("%p")
	if err := x.Read(a); err != nil {
		fmt.Println(err)
	}
	//a := 200
	//addr := unsafe.Pointer(&a)
	//fmt.Println(addr,*(*int64)(addr))
	size := *(*uint32)(unsafe.Pointer(&b))
	fmt.Println("size is:",size)
}

func(x *DemoX) Read(bs []byte) error {
	var err error
	var bytesRead int
	bytesToRead := len(bs)
	for err == nil && bytesRead < bytesToRead {
		var readLen int
		fmt.Printf("%p\n",&bs)
		bs = []byte("1123")//这样赋值和下面的复制方式不一样
		//fmt.Printf("%p\n",bs)

		//fmt.Println(len(bs))
		//bs[0] = uint8(4)
		//bs[1] = uint8(1)
		//bs[2] = uint8(1)
		//bs[3] = uint8(2)
		readLen = 4
		if readLen <= 0 {
			break
		}
		bytesRead += readLen
	}
	return nil
}

func BasicParam01(m int){
	m = 10
}
func BasicParam(m *int){
	*m = 10
}

func MapParam01(m map[int]int){
	fmt.Printf("%p\n",m)
	m[4] = 10
}
func MapParam02(m *map[int]int){
	fmt.Printf("%p\n",m)
	//*m[3] = 10//语法错误
	cc := *m
	cc[3] = 10
}


func SliceParam(m *[]int){
	//m[10] = 10//如果参数是指针报错，无法修改
	xx := *m
	xx[2] = 10//曲线赋值成功
	for _,v := range *m{//可以对指针进行迭代，但是无法赋值
		fmt.Println(v)
	}
}


