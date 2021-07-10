package op

import (
	"github.com/go-ole/go-ole"
)

func (com *Opsoft) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Opsoft) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *Opsoft) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *Opsoft) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("OcrEx", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *Opsoft) SetDict(index int, file string) int {
	ret, _ := com.op.CallMethod("SetDict", index, file)
	return int(ret.Val)
}

func (com *Opsoft) UseDict(index int) int {
	ret, _ := com.op.CallMethod("UseDict", index)
	return int(ret.Val)
}

func (com *Opsoft) OcrAuto(x1, y1, x2, y2 int, sim float32) string {
	ret, _ := com.op.CallMethod("OcrAuto", x1, y1, x2, y2, sim)
	return ret.ToString()
}

func (com *Opsoft) OcrFromFile(fileName, colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("OcrFromFile", fileName, colorFormat, sim)
	return ret.ToString()
}

func (com *Opsoft) OcrAutoFromFile(colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("OcrAutoFromFile", colorFormat, sim)
	return ret.ToString()
}
