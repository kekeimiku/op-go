package op

//op version
func (com *opsoft) Ver() string {
	ver, _ := com.op.CallMethod("Ver")
	return ver.ToString()
}

//
