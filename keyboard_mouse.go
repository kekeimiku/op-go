package op

import "github.com/go-ole/go-ole"

func (com *Opsoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *Opsoft) GetKeyState(vkCode int) int {
	ret, _ := com.op.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyDown(vkCode string) int {
	ret, _ := com.op.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) KeyPress(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) KeyUp(vkCode int) int {
	ret, _ := com.op.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.op.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) LeftClick() int {
	ret, _ := com.op.CallMethod("LeftClick")
	return int(ret.Val)
}

func (com *Opsoft) LeftDoubleClick() int {
	ret, _ := com.op.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

func (com *Opsoft) LeftDown() int {
	ret, _ := com.op.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *Opsoft) LeftUp() int {
	ret, _ := com.op.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *Opsoft) MiddleClick() int {
	ret, _ := com.op.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *Opsoft) MiddleDown() int {
	ret, _ := com.op.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *Opsoft) MiddleUp() int {
	ret, _ := com.op.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *Opsoft) MoveR(rx, ry int) int {
	ret, _ := com.op.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *Opsoft) MoveTo(x, y int) int {
	ret, _ := com.op.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *Opsoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.op.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

func (com *Opsoft) RightClick() int {
	ret, _ := com.op.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *Opsoft) RightDown() int {
	ret, _ := com.op.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *Opsoft) RightUp() int {
	ret, _ := com.op.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *Opsoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.op.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

func (com *Opsoft) WheelDown() int {
	ret, _ := com.op.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *Opsoft) WheelUp() int {
	ret, _ := com.op.CallMethod("WheelUp")
	return int(ret.Val)
}
