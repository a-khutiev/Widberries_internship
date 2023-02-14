package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	a := []string{"a", "b", "c", "e", "d"}

	// var i int

	// 1. Без сохранения порядка эл-тов

	// 1. Копировать последний элемент в индекс i.
	// a[i] = a[len(a)-1]
	// 2. Удалить i-ый элемент из слайса.
	// a[len(a)-1] = ""
	// 3. Усечь срез
	// a = a[:len(a)-1]

	// 2. Сохраненеие порядка

	// 1. Сдвинуть
	// copy(a[i:], a[i+1:])
	// 2. Удалить последний элемент
	// a[len(A)-1] = ""
	// 3. Усечь срез.
	// a = a[a:len(a)-1]
	fmt.Println(a)
}
