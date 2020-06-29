package basic_grammar

import (
	"flag"
	"fmt"
	"os"
)

var (
	h          bool
	configPath string
	logFile    string
	logPath    string
	autoScan   bool
	head       int
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: ....Options:`)
	flag.PrintDefaults()
}
func main() {

	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&logFile, "logFile", "", "日志文件")
	flag.StringVar(&configPath, "config", "./config.json", "配置文件路径")
	flag.StringVar(&logPath, "logPath", "./log", "日志目录，和autoscan配合使用")
	flag.BoolVar(&autoScan, "autoScan", false, "是否自动扫描目录")
	flag.IntVar(&head, "head", 10, "展示前N行")
	//flag.Usage = usage
	flag.Parse() //flag provided but not defined: -config
	if h {
		flag.Usage()
	}
	if logFile == "" && logPath == "" {
		fmt.Println("logFile or logPath is required!")
		os.Exit(0)
	}
	//if logFile==""{
	//	autoScan=true
	//}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("restart program...", err)
			main()
		}
	}()
	fmt.Println("log parser running.....")

}
