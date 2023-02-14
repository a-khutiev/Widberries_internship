package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//  Символы могут быть unicode.

func main() {
	var s string
	fmt.Scanln(&s)
	print(reverse(s))
}

func reverse(s string) string {

	a := strings.Split(s, "")
	for j, i := len(a)-1, 0; i != len(a)/2; j-- {
		a[i], a[j] = a[j], a[i]
		i++
	}
	s = strings.Join(a, "")
	return s
}
