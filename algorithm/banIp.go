package algorithm

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
	"fmt"
)
//在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
//每个IP三分钟之内只能访问一次。

type IpBan struct {
	visitIps map[string]time.Time
	lock sync.Mutex
}
func NewIpBan() *IpBan{
	return &IpBan{
		visitIps:make(map[string]time.Time),
	}
}
func (b *IpBan)visit(ip string) bool{
	b.lock.Lock()
	defer b.lock.Unlock()
	if _,ok := b.visitIps[ip];ok {
		return true
	}
	b.visitIps[ip] = time.Now()
	return false
}
func(b *IpBan)removeIp(ip string){
	delete(b.visitIps,ip)
}

func(b *IpBan)clean(ctx context.Context){


	go func() {
		timer2 := time.NewTicker(time.Second)

		//timer := time.NewTimer(time.Second)//timer需要重置，ticker不需要
		for {
			select {
			case <- timer2.C:
				b.lock.Lock()
				for k,v := range b.visitIps{
					if time.Now().Sub(v) > time.Minute*3{
						b.removeIp(k)
					}
				}
				b.lock.Unlock()
				//timer.Reset(time.Second * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
}
func Visit(){
	ipBan := NewIpBan()
	successNum := int32(0)
	wg := sync.WaitGroup{}
	wg.Add(100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ipBan.clean(ctx)

	for count:=0;count<1000;count++{
		for i:=0;i<100;i++{
			go func(ii int) {
				ip := fmt.Sprintf("192.100.1.%d",ii)
				if !ipBan.visit(ip){
					defer wg.Done()
					atomic.AddInt32(&successNum,1)
				}
			}(i)
		}
	}
	wg.Wait()
	fmt.Println("success num:",successNum)
}
