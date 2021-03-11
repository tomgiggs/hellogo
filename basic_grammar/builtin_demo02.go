package basic_grammar

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
)

func StringDemo() {
	var stringDemo = strings.NewReader("addMoney 1000 200")
	var cmdNmae string
	fmt.Fscanf(stringDemo, "%s", &cmdNmae)
	fmt.Printf("cmd name is: %s\n", cmdNmae)
	var num = int64(0)
	n, _ := fmt.Fscanf(stringDemo, "%d", &num)
	if n < 1 {
		print("n is:", n)
	}
	fmt.Printf("the number is: %d\n", num)
	var num2 = int64(0)
	fmt.Fscanf(stringDemo, "%d", &num2)
	fmt.Printf("the number2 is: %d", num2)
}

func MathDemo() {
	//p := rand.Int31()
	p := rand.Int31n(200) //返回[0,n)中的一个整数
	print(p)
}

type UserAges struct {
	ages map[string] int//ages没有暴露给外部，导致调用者在外部无法初始化ages，两种修改，一种是变成大写，一种是写一个new方法给外部调用
	sync.Mutex
}
func NewUserAge() *UserAges{
	return &UserAges{
		ages:make(map[string]int),
	}
}


func (u *UserAges)Add(name string,age int)  {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}

func (u *UserAges)Get(name string)int{
	if age,ok:=u.ages[name];ok{
		return age
	}
	return -1
}

func TypeDemo02(){
	type Myint int
	var i int =1
	var j Myint = Myint(i)
	fmt.Println(j)
}
func BasicDemo02(){
	//重复使用短变量声明时，需要保证短变量声明语句中至少要声明一个新的变量，否则直接使用赋值语句 = 就可以了
	aa, err := os.Open("")
	if err != nil {

	}
	bb, err := os.Open("")
	fmt.Println(aa,bb)
}