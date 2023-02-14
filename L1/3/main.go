package main

//Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
import (
	"fmt"
	"time"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	var ans int
	var i int

	for i = 0; i != 5; i++ {
		go SquarePrintSum(nums[i], &ans)
	}
	time.Sleep(time.Second)
	fmt.Printf("%d", ans)
}

func SquarePrintSum(x int, ans *int) {
	*ans += x * x
}
