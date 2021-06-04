package main

import (
	"github.com/kekeimiku/op-go"
)

func main() {
	op, _ := op.Load()
	op.BindWindow(op.FindWindow("", "Aim Hero"), "normal", "windows", "windows", 0)
	var x, y int = 0, 0
	for {
		if op.FindPic(163, 244, 638, 607, "3.bmp", "000000", 0.7, 0, &x, &y) != -1 {
			op.MoveTo(x, y)
			op.LeftClick()
		}
	}
}
