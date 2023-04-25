package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var s string
	fmt.Scanln(&s)

	e, err := Culc(s)
	if err != nil {
		return
	}

	out(e)

	return
}

// Culc - декодирует строку
func Culc(s string) (string, error) {
	sSlice := []rune(s)
	var err error
	var resultSlice []rune

	for i := 0; i < len(sSlice); i++ {

		if i != 0 && isDigit(sSlice[i]) {

			if !isDigit(sSlice[i-1]) {
				j, _ := strconv.Atoi(string(sSlice[i]))
				for ; j != 1; j-- {
					resultSlice = append(resultSlice, sSlice[i-1])
				}
				continue
			} else if isDigit(sSlice[i-1]) {
				io.WriteString(os.Stderr, ("is non-correct string"))
				return "", err
			}
		}
		resultSlice = append(resultSlice, sSlice[i])
	}

	var finBuf string
	for _, a := range resultSlice {
		finBuf = finBuf + fmt.Sprintf("%c", a)
	}

	return finBuf, err
}
func isDigit(i rune) bool { // проверяет цифра ли символ
	if i >= 48 && i <= 57 {
		return true
	}
	return false
}
func out(finBuf string) { // печать в консоль
	fmt.Fprint(io.Writer(os.Stdout), finBuf)
	return
}
