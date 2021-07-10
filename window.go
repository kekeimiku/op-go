package op

import "github.com/go-ole/go-ole"

func (com *Opsoft) ClientToScreen(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("ClientToScreen", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *Opsoft) EnumProcess(name string) string {
	ret, _ := com.op.CallMethod("EnumProcess", name)
	return ret.ToString()
}

func (com *Opsoft) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.op.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

func (com *Opsoft) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.op.CallMethod("EnumWindowByProcessId", pid, title, className, filter)
	return ret.ToString()
}

func (com *Opsoft) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.op.CallMethod("EnumWindowSuper", spec1, flag1, type1, spec2, flag2, type2, sort)
	return ret.ToString()
}

func (com *Opsoft) FindWindow(class, title string) int {
	ret, _ := com.op.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

func (com *Opsoft) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

func (com *Opsoft) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowByProcessId", processId, class, title)
	return int(ret.Val)
}

func (com *Opsoft) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

func (com *Opsoft) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.op.CallMethod("FindWindowSuper", spec1, flag1, type1, spec2, flag2, type2)
	return int(ret.Val)
}

func (com *Opsoft) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.op.CallMethod("GetClientRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *Opsoft) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.op.CallMethod("GetClientSize", hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	pWidth.Clear()
	pHeight.Clear()
	return int(ret.Val)
}

func (com *Opsoft) GetForegroundFocus() int {
	ret, _ := com.op.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

func (com *Opsoft) GetForegroundWindow() int {
	ret, _ := com.op.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

func (com *Opsoft) GetMousePointWindow() int {
	ret, _ := com.op.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *Opsoft) GetPointWindow(x, y int) int {
	ret, _ := com.op.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *Opsoft) GetProcessInfo(pid int) string {
	ret, _ := com.op.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

func (com *Opsoft) GetSpecialWindow(flag int) int {
	ret, _ := com.op.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

func (com *Opsoft) GetWindowClass(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *Opsoft) GetWindowProcessId(hwnd int) int {
	ret, _ := com.op.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *Opsoft) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *Opsoft) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.op.CallMethod("GetWindowRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *Opsoft) GetWindowState(hwnd, flag int) int {
	ret, _ := com.op.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *Opsoft) GetWindowTitle(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

func (com *Opsoft) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.op.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *Opsoft) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("ScreenToClient", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *Opsoft) SendPaste(hwnd int) int {
	ret, _ := com.op.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

func (com *Opsoft) SendString(hwnd int, str string) int {
	ret, _ := com.op.CallMethod("SendString", hwnd, str)
	return int(ret.Val)
}

func (com *Opsoft) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.op.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *Opsoft) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.op.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *Opsoft) SetWindowState(hwnd, flag int) int {
	ret, _ := com.op.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *Opsoft) SetWindowText(hwnd int, title string) int {
	ret, _ := com.op.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *Opsoft) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.op.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
