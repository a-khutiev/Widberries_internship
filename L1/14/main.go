package main

import "fmt"

//Разработать программу,
//которая в рантайме способ на определить тип переменной: int, string, bool, channel из переменной типа interface{}.
type typee interface{}

func main() {
	itg := 2
	str := "man"
	var channelItf interface{} = make(chan int)

	fmt.Printf("%s - %s - %s", typeOfMy(itg), typeOfMy(str), typeOfMy(channelItf))
}
func typeOfMy(a typee) string {
	switch a.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	case string:
		return "string"
	case chan int:
		return "chan int"
	default:
		return "<nil>"
	}
	return "nil"
}
