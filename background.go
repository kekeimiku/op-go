package op

func (com *Opsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.op.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *Opsoft) UnBindWindow() int {
	ret, _ := com.op.CallMethod("UnBindWindow")
	return int(ret.Val)
}
