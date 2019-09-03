package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

type Status struct {
	isRunning bool
}

func main() {

	status := Status{isRunning: true}

	go func(currentStatus *Status) {
		for true {
			if currentStatus.isRunning {
				robotgo.MouseClick("left", true)
				//fmt.Println("click left button")
			} else {
				//fmt.Println("stop click button")
			}
		}
	}(&status)
	for true {
		mRight := robotgo.AddEvent("mright")
		if mRight {
			status.isRunning = false
			fmt.Println("you press... ", "mouse right button")
		}
		center := robotgo.AddEvent("center")
		if center {
			status.isRunning = true
			fmt.Println("you press... ", "mouse left button")
		}

	}
}
