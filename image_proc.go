package op

import "github.com/go-ole/go-ole"

func (com *Opsoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.op.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *Opsoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.op.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *Opsoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Opsoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *Opsoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Opsoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Opsoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Opsoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Opsoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.op.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

func (com *Opsoft) CapturePre(file string) int {
	ret, _ := com.op.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *Opsoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.op.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

func (com *Opsoft) GetScreenData(x1, y1, x2, y2 int) int {
	ret, _ := com.op.CallMethod("GetScreenData", x1, y1, x2, y2)
	return int(ret.Val)
}

func (com *Opsoft) GetScreenDataBmp(x1, y1, x2, y2 int, data, size *int) int {
	d := ole.NewVariant(ole.VT_I4, 0)
	s := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("GetScreenDataBmp", x1, y1, x2, y2, &data, &size)
	*data = int(d.Val)
	*size = int(s.Val)
	d.Clear()
	s.Clear()
	return int(ret.Val)
}

func (com *Opsoft) SetDisplayInput(mode string) int {
	ret, _ := com.op.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}
