package op

func (com *Opsoft) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.op.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *Opsoft) SetMouseDelay(types string, delay int) int {
	ret, _ := com.op.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *Opsoft) SetMouseSpeed(speed int) int {
	ret, _ := com.op.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *Opsoft) SetSimMode(mode int) int {
	ret, _ := com.op.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}
