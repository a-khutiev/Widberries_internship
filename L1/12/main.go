package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
import "fmt"

func main() {
	a := make(map[string]struct{})
	m := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, i := range m {
		a[i] = struct{}{}
	}
	mas := []string{}
	for j, _ := range a {
		mas = append(mas, j)
	}
	fmt.Println(mas)
}
