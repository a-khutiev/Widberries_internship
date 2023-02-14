package main

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	a int
	sync.Mutex
}

func main() {

	a := Counter{
		a: 0,
	}
	for i := 0; i < 10; i++ {
		go func() {
			a.inkrement()
		}()
	}
	time.Sleep(time.Second)
	fmt.Print(a.a)
}
func (a *Counter) inkrement() {
	a.Lock()
	a.a++
	a.Unlock()
}
