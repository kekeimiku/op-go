package op

func (com *Opsoft) GetBasePath() string {
	ret, _ := com.op.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *Opsoft) GetID() int {
	ret, _ := com.op.CallMethod("GetID")
	return int(ret.Val)
}

func (com *Opsoft) GetLastError() int {
	ret, _ := com.op.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *Opsoft) GetPath() string {
	ret, _ := com.op.CallMethod("GetPath")
	return ret.ToString()
}

func (com *Opsoft) SetPath(path string) int {
	ret, _ := com.op.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *Opsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.op.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *Opsoft) Ver() string {
	ver, _ := com.op.CallMethod("Ver")
	return ver.ToString()
}

func (com *Opsoft) EnablePicCache(enable int) int {
	ret, _ := com.op.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}
