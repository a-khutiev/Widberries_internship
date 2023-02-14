package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
import (
	"fmt"
	"os"
)

func main() {
	var dad *Human
	var err error
	dad, err = New("Bond", "Vitaly", 50, "male")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	fmt.Print(dad, err)

	var dougther Action
	dougther.schoolAdress = "Pushkino 1"
	dougther.Human, err = New("Bond", "Alena", 12, "female")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	fmt.Print(dougther.Human, dougther.schoolAdress)
}

type Human struct {
	surname string
	name    string
	age     int
	sex     string
}
type Action struct {
	*Human
	schoolAdress string
}

func New(surname string, name string, age int, sex string) (*Human, error) {
	return &Human{
		surname: surname,
		name:    name,
		age:     age,
		sex:     sex,
	}, nil
}
