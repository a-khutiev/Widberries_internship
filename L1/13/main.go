package main

import "fmt"

//поменять местами два числа без создания временной переменной

func main() {
	var A int = 1
	var B int = 2
	fmt.Printf("%d %d\n", A, B)
	A, B = B, A
	fmt.Print(A, B)

}
