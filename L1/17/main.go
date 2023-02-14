package main

import "fmt"

func MyQuickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		MyQuickSort(A, p, q-1)
		MyQuickSort(A, q+1, r)
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

func BinSearch(key int, a []int) int {
	return BinSearchProc(key, a, 0, len(a)-1)
}

func BinSearchProc(a int, mas []int, L int, R int) int {
	if L > R {
		return -1
	}

	medNumber := (L + R) / 2
	if a == mas[medNumber] {
		return medNumber
	} else if a < mas[medNumber] {
		R = medNumber - 1
	} else if a > mas[medNumber] {
		L = medNumber + 1
	}

	return BinSearchProc(a, mas, L, R)
}

func main() {
	a := []int{1, 2}
	MyQuickSort(a, 0, len(a)-1)
	b := BinSearch(2, a)
	fmt.Print(b)
}
