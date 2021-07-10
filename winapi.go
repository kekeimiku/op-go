package op

func (com *Opsoft) RunApp(path string, mode int) int {
	ret, _ := com.op.CallMethod("RunApp", path, mode)
	return int(ret.Val)
}

func (com *Opsoft) WinExec(path string, dis int) int {
	ret, _ := com.op.CallMethod("WinExec", path, dis)
	return int(ret.Val)
}

func (com *Opsoft) GetCmdStr(path string, time int) string {
	ret, _ := com.op.CallMethod("GetCmdStr", path, time)
	return ret.ToString()
}
