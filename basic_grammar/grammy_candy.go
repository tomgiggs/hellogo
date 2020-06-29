package basic_grammar

import "fmt"

type Car interface {
	run()
}

type Honda struct {
}

func (s Honda) run() {

}

var _ Car = Honda{}

func TestCandy() {
	fmt.Printf("print:", 200)

}
