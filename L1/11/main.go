package main

//Реализовать пересечение двух неупорядоченных множеств.
import "fmt"

func main() {
	a := []int{3, 4, 5, 6, 7}

	b := []int{1, 2, 6, 4, 5}
	fmt.Print(cap(a, b))

}

func cap(a, b []int) []int {
	mape := make(map[int]int)
	for _, i := range a {
		mape[i] = 1
	}

	var acapb []int
	for _, j := range b {
		if mape[j] == 1 {
			acapb = append(acapb, j)
		}
	}
	return acapb
}
