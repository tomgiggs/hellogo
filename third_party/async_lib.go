package third_party

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"sync"
	"testing"
)

type RunFunc func(in interface{}, out interface{}) error

func (f RunFunc) GetName() string {
	return reflect.TypeOf(f).Name()
}

type Async struct {
	mu        sync.Mutex
	wait      sync.WaitGroup
	tasks     []RunFunc
	err       chan error
	paramMap  map[string]interface{}
	resultMap map[string]interface{}
	results   chan interface{}
	callback func(map[string]interface{}) error
}

func (a *Async) AddFunc(f RunFunc) {
	a.mu.Lock()
	a.tasks = append(a.tasks, f)
	a.mu.Unlock()
}
func (a *Async) AddCalalbackFunc(callback func(map[string]interface{})error)  {
	a.callback = callback
}

func (a *Async) Parallel()error {
	a.wait.Add(len(a.tasks))
	for _, f := range a.tasks {
		go func(fn RunFunc) {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 1024)
					runtime.Stack(buf, false)
					log.Error(string(buf))
				}
			}()
			defer a.wait.Done()
			fnName := reflect.TypeOf(fn).Name()
			a.err <- fn(a.paramMap[fnName],a.resultMap[fnName])
		}(f)
	}
	return  a.callback(a.resultMap)
}

func AsyncTest(t *testing.T){

}