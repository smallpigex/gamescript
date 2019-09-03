package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	flag := false
	for true {
		mright := robotgo.AddEvent("mright")
		if mright {
			flag = true
			fmt.Println("you press... ", "mouse right button")
		}
		if flag {
			fmt.Println("stop click button")
		} else {
			//robotgo.MouseClick("left", true)
			fmt.Println("click left button")
		}

	}
	//for true {
	//	robotgo.MouseClick("left", true)
}
