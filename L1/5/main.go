package main

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
import (
	"fmt"
	"time"
)

func main() {
	strChannel := make(chan string)
	go send(&strChannel)
	go printer(strChannel)
	time.Sleep(time.Second * 5)
}
func send(strChannel *chan string) {
	line := ""
	for i := 1; i < 5; i++ {
		fmt.Scanln(&line)
		*strChannel <- line
	}
	close(*strChannel)
}
func printer(strChannel chan string) {

	for a := range strChannel {
		fmt.Printf("from Channel:%s\n", a)
	}
	close(strChannel)
}
