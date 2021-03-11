package basic_grammar

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
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

func SelectDemo01() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

}

func GetLocalIp() {
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
func WaitGroupDemo01() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func WaitGroupDemo02() {
	runtime.GOMAXPROCS(1) //这个会影响打印顺序，始终最后一个任务先执行
	wg := sync.WaitGroup{}
	wg.Add(10) //这个要跟完成的done次数想匹配，不然会报错，当大于256后，会打算执行
	//当n大于257的时候，应该重新分配了新的队列。所以顺序被打乱了，接着在新队列的某个位置开始加新的任务。 原先的任务被随机分成A和B两堆任务，接着在A后面开始加新的任务。
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("i first: ", i)
	//		wg.Done()
	//	}()
	//}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	//time.Sleep(5 * time.Second)
	wg.Wait()
}
var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex

)
func AtomicDemo01() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait() //等待goroutine结束
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//mutex.Lock()
		atomic.AddInt64(&counter, 1) //安全的对counter加1，LoadInt64，StoreInt64
		runtime.Gosched()// 当调用 runtime.Gosched 函数强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行
	}
	mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
}


