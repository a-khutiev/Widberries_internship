package main

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
import (
	"bufio"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()
	println(reverseStrs(line))
}

func reverseStrs(a string) string {
	s := strings.Split(a, " ")

	for i, j := 0, len(s)-1; i != len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	res := strings.Join(s, " ")
	return res
}
