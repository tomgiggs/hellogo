package basic_grammar

import (
	"fmt"
	"regexp"
	"strings"
)

func RegDemo01() {
	//regStr := `([\d\.]+)\s-\s(.*?)\s\[(.*?)\]\s"(.*?)\s(.*?)\s(.*?)"\s(\d+)\s(\d+)\s"(.*?)"\s"(.*?)"`
	//reg := regexp.MustCompile(regStr)
	//fmt.Printf("%q\n", reg.FindAllString("100.116.222.152 - - [19/Sep/2018:15:28:14 +0800] \"GET /api/child_star/query?classify=2&page=1&page_size=18 HTTP/1.1\" 301 178 \"-\" \"okhttp/3.10.0\"", -1))
	var logStr string = `2020-06-03 00:17:29 [Shop]:23[9223][27],2,0,0,1591114649,1`
	regCommonHeader := regexp.MustCompile(`[0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+:[0-9]+ \[[a-zA-Z]+\]:`)
	//fmt.Printf("%q\n", reg.FindAllString(logStr, -1))

	commonHeaderStr := regCommonHeader.FindString(logStr)
	headerLen := len(commonHeaderStr)
	logBody := logStr[headerLen:]
	commonHeaderInfo := strings.Split(commonHeaderStr, " [")
	logDate := commonHeaderInfo[0]
	//fmt.Println(reflect.TypeOf(logDate))
	logMod := commonHeaderInfo[1]
	logMod = logMod[:len(logMod)-2]
	fmt.Println(logDate, "-------", logMod)
	fmt.Println(strings.Join(strings.SplitN(logBody, ",", -1), "||"))

	fmt.Printf("%q\n", commonHeaderStr)

}
