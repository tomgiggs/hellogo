package algorithm

import (
	"fmt"
	"sync"
)

func PrintNumAndLetter(){

	letter,num := make(chan bool),make(chan  bool)
	wg := sync.WaitGroup{}
	begin := 1
	wg.Add(1	)
	go func() {
		for {
			select {
			case <-letter:
				fmt.Print(begin)
				begin++
				fmt.Print(begin)
				begin++
				num <- true
				break
			default:
				break
			}
		}
	}()

	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		strLen := len(str)
		i := 0
		for {
			select {
			case <-num:
				if i> strLen-2{
					wg.Done()
					return
				}
				fmt.Print(str[i:i+2])
				i += 2
				letter <- true
				break
			}
		}
	}()
	letter <- true
	wg.Wait()
}
