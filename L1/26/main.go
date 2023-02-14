package main

// 26. Разработать программу, которая проверяет,
//  что все символы в строке уникальные (true — если уникальные, false etc).
//  Функция проверки должна быть регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
import (
	"fmt"
	"strings"
)

func main() {

	var a string = "abCdefAaf"
	temp := strings.ToLower(a)
	mape := make(map[rune]bool)
	var flag string = "true"
	for _, i := range temp {
		if mape[i] == true {
			flag = "false"
		}
		mape[i] = true
	}
	fmt.Printf("%s - %s", a, flag)

}
