package main

// реализовать функцию Sleep
import (
	"time"
)

func sleepMySleep(d time.Duration) {
	hL, mL, sL := time.Now().Clock()

	hL = hL + int(d.Hours())
	mL = mL + int(d.Minutes())
	sL = sL + int(d.Seconds())

	for {
		durH, durM, durS := time.Now().Clock()

		if hL == durH && durM == mL && durS == sL {
			return
		}
	}
}

func main() {
	sleepMySleep(time.Second * 2)
	print("aaa")
}
