package main

//22. Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"strings"
)

func main() {

	var a, b float64
	var operation string
	fmt.Println("Введите а: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Incorrect input")
		return
	}
	fmt.Println("Введите b: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	fmt.Println("Операция:  (* / + -)")
	_, err = fmt.Scanln(&operation)
	operation = strings.TrimSpace(operation)
	switch operation {
	case "*":
		fmt.Print(a * b)
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "/":
		fmt.Print(a / b)
	}

}
