package main

// Написать программу,
// которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
import (
	"fmt"
	"time"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	var i int
	for i = 0; i != 5; i++ {
		go SquarePrint(nums[i])
	}
	time.Sleep(time.Second)
}

func SquarePrint(x int) {
	fmt.Printf("%d\n", x*x)
}
