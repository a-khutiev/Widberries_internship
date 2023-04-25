package main

import (
	"fmt"

	"github.com/a-khutiev/Widberries_internship/L2/develop/dev01/task"
)

func main() {
	err := task.Ain()
	if err != nil {
		fmt.Print(err)
	}
}
