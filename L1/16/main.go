//реализовать quicksort встроенными методами языка
package main

// package mySort

import "fmt"

func myQuickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		myQuickSort(A, p, q-1)
		myQuickSort(A, q+1, r)
	}
}
func partition(A []int, p, r int) int {
	var temp int
	x := A[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i = i + 1
			temp = A[i]
			A[i] = A[j]
			A[j] = temp
		}
	}
	temp = A[i+1]
	A[i+1] = A[r]
	A[r] = temp
	return i + 1
}

func main() {
	mas := []int{5, 1, 9, 10, 12, 2, 3, 4}
	myQuickSort(mas, 0, len(mas)-1)
	fmt.Printf("%d", mas)
}
