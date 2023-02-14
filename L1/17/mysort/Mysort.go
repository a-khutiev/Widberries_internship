package Mysort

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
