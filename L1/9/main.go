//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"time"
)

func main() {
	oneChan := make(chan int)
	twoChan := make(chan int)

	massive := []int{1, 92, 4, 24, 13, 5, 4, 4, 12, 35, 362, 34, 123, 562, 33, 4332, 34, 99, 234, 0, -34, 324, -33}

	go square(oneChan, twoChan)

	for i := range massive {
		oneChan <- i
		fmt.Printf("%d^2 = %d\n", i, <-twoChan)
	}
	close(twoChan)
	time.Sleep(time.Second * 1)
}
func square(oneChan chan int, twoChan chan int) {
	for a := range oneChan {
		twoChan <- a * a
	}
	close(oneChan)
}
