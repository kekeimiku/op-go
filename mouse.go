package op

func (com *opsoft) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.op.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *opsoft) SetMouseDelay(types string, delay int) int {
	ret, _ := com.op.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *opsoft) SetMouseSpeed(speed int) int {
	ret, _ := com.op.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *opsoft) SetSimMode(mode int) int {
	ret, _ := com.op.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}
