package basic_grammar

import (
	"fmt"
	"net"
	"os"
)

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


func GetLocalIp(){
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}
}