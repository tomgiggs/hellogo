package basic_grammar

import "fmt"

func SwitchDemo01() {
	cond := 20
	switch cond {
	case 10:
		fmt.Println("case 10 run")
	case 15:
		fmt.Println("case 15 run")
	case 11, 20:
		fmt.Println("mutli value of case run")
	}
}
