//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
	"time"
)

type pair struct {
	key   int
	value string
}

func main() {
	//

	myMap := make(map[int]byte)
	var a sync.Mutex

	go mapX(&a, myMap)
	go mapY(&a, myMap)
	go mapZ(&a, myMap)

	time.Sleep(time.Second * 3)
	var ch int

	for key, value := range myMap {
		fmt.Printf("%d_ : %d : %c \n", ch, key, value)
		ch++
	}

}

func mapX(a *sync.Mutex, myMap map[int]byte) {
	for i := 0; i < 10; i++ {
		a.Lock()
		myMap[i] = 'X'
		a.Unlock()
	}
}
func mapY(a *sync.Mutex, myMap map[int]byte) {
	for i := 10; i < 20; i++ {
		a.Lock()
		myMap[i] = 'Y'
		a.Unlock()
	}
}
func mapZ(a *sync.Mutex, myMap map[int]byte) {
	for i := 20; i < 31; i++ {
		a.Lock()
		myMap[i] = 'Z'
		a.Unlock()
	}
}
func get(a sync.Mutex) {

}
