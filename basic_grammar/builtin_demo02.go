package basic_grammar

import (
	"fmt"
	"math/rand"
	"strings"
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
