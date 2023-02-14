package main

// 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.
// var justString string    - лучше не использовать глобальные переменные, усложняют тестирование
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]    - работать лучше со срезом из runes нежели byte, вероятно переполнение массива поэтому можно задатт cap через мейк
// }
// func main() {
//   someFunc()
// }
//

//	"strings"

func createHugeString(i int) string {
	return string(i)
}

func someFunc() string {
	v := createHugeString(1 << 11)
	vRuns := make([]rune, 10, 200)
	copy(vRuns, []rune(v))
	justString := vRuns[:100]
	return string(justString)
}
func main() {
	str_ := someFunc()
	println(str_)
}
