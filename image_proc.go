package op

import "github.com/go-ole/go-ole"

/*
函数简介:

抓取指定区域(x1, y1, x2, y2)的图像,保存为file(24位位图)

参数定义:

x1 int 区域的左上X坐标.

y1 int 区域的左上Y坐标.

x2 int 区域的右下X坐标.

y2 int 区域的右下Y坐标.

file string 保存的文件名.

保存的地方一般为SetPath中设置的目录

当然这里也可以指定全路径名.

返回值:

int 0:失败 1:成功
*/
func (com *opsoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.op.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

/*
函数简介:

比较指定坐标点(x,y)的颜色

参数定义:

x int X坐标

y int Y坐标

color string 颜色字符串,可以支持偏色,多色,例如 "ffffff-202020|000000-000000" 这个表示白色偏色为202020,和黑色偏色为000000.颜色最多支持10种颜色组合. 注意，这里只支持RGB颜色.

sim float 相似度(0.1-1.0)

返回值:

int: 0: 颜色匹配 1: 颜色不匹配
*/
func (com *opsoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.op.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

/*
函数简介:

查找指定区域内的颜色,颜色格式"RRGGBB-DRDGDB",注意,和按键的颜色格式相反

参数定义:

x1 int 区域的左上X坐标

y1 int 区域的左上Y坐标

x2 int 区域的右下X坐标

y2 int 区域的右下Y坐标

color string 颜色 格式为"RRGGBB-DRDGDB",比如"123456-000000|aabbcc-202020".注意，这里只支持RGB颜色.

sim float32 相似度,取值范围0.1-1.0

dir int 查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左,从下到上 4：从中心往外查找 5: 从上到下,从左到右 6: 从上到下,从右到左 7: 从下到上,从左到右 8: 从下到上,从右到左.

intX *int 返回X坐标

intY *int 返回Y坐标

返回值:

int 0:没找到 1:找到
*/
func (com *opsoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

/*
函数简介:

查找指定区域内的所有颜色,颜色格式"RRGGBB-DRDGDB",注意,和按键的颜色格式相反

参数定义:

x1 int 区域的左上X坐标

y1 int 区域的左上Y坐标

x2 int 区域的右下X坐标

y2 int 区域的右下Y坐标 color 字符串:颜色 格式为"RRGGBB-DRDGDB" 比如"aabbcc-000000|123456-202020".注意，这里只支持RGB颜色.

sim float32 相似度,取值范围0.1-1.0

dir int 查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左,从下到上 5: 从上到下,从左到右 6: 从上到下,从右到左 7: 从下到上,从左到右 8: 从下到上,从右到左

返回值:

string 返回所有颜色信息的坐标值,然后通过GetResultCount等接口来解析 (由于内存限制,返回的颜色数量最多为1800个左右)
*/
func (com *opsoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

/*
函数简介:

根据指定的多点查找颜色坐标

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

first_color 字符串:颜色格式为"RRGGBB-DRDGDB|RRGGBB-DRDGDB|…………",比如"123456-000000"

这里的含义和按键自带Color插件的意义相同，只不过我的可以支持偏色和多种颜色组合

所有的偏移色坐标都相对于此颜色.注意，这里只支持RGB颜色.

offset_color 字符串: 偏移颜色可以支持任意多个点 格式和按键自带的Color插件意义相同, 只不过我的可以支持偏色和多种颜色组合

格式为"x1|y1|RRGGBB-DRDGDB|RRGGBB-DRDGDB……,……xn|yn|RRGGBB-DRDGDB|RRGGBB-DRDGDB……"

比如"1|3|aabbcc|aaffaa-101010,-5|-3|123456-000000|454545-303030|565656"等任意组合都可以，支持偏色

还可以支持反色模式，比如"1|3|-aabbcc|-334455-101010,-5|-3|-123456-000000|-353535|454545-101010","-"表示除了指定颜色之外的颜色.

sim 双精度浮点数:相似度,取值范围0.1-1.0 dir 整形数:查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左, 从下到上 .

intX 变参指针:返回X坐标(坐标为first_color所在坐标) .

intY 变参指针:返回Y坐标(坐标为first_color所在坐标)

返回值:

整形数: 0:没找到 1:找到
*/
func (com *opsoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

/*
函数简介:

根据指定的多点查找所有颜色坐标

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

first_color 字符串:颜色 格式为"RRGGBB-DRDGDB|RRGGBB-DRDGDB|…………",比如"123456-000000"

这里的含义和按键自带Color插件的意义相同，只不过我的可以支持偏色和多种颜色组合

所有的偏移色坐标都相对于此颜色.注意，这里只支持RGB颜色.

offset_color 字符串: 偏移颜色 可以支持任意多个点 格式和按键自带的Color插件意义相同, 只不过我的可以支持偏色和多种颜色组合

格式为"x1|y1|RRGGBB-DRDGDB|RRGGBB-DRDGDB……,……xn|yn|RRGGBB-DRDGDB|RRGGBB-DRDGDB……"

比如"1|3|aabbcc|aaffaa-101010,-5|-3|123456-000000|454545-303030|565656"等任意组合都可以，支持偏色

还可以支持反色模式，比如"1|3|-aabbcc|-334455-101010,-5|-3|-123456-000000|-353535|454545-101010","-"表示除了指定颜色之外的颜色.

sim 双精度浮点数:相似度,取值范围0.1-1.0 dir 整形数:查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左, 从下到上

返回值:

字符串: 返回所有颜色信息的坐标值,然后通过GetResultCount等接口来解析(由于内存限制,返回的坐标数量最多为1800个左右)

坐标是first_color所在的坐标
*/
func (com *opsoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

/*
函数简介:

查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.

这个函数可以查找多个图片,只返回第一个找到的X Y坐标.

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

pic_name 字符串:图片名,可以是多个图片,比如"test.bmp|test2.bmp|test3.bmp"

delta_color 字符串:颜色色偏比如"203040" 表示RGB的色偏分别是20 30 40 (这里是16进制表示)

sim 双精度浮点数:相似度,取值范围0.1-1.0

dir 整形数:查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左, 从下到上

intX 变参指针:返回图片左上角的X坐标

intY 变参指针:返回图片左上角的Y坐标

返回值:

整形数: 返回找到的图片的序号,从0开始索引.如果没找到返回-1
*/
func (com *opsoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

/*
函数简介:

查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.

这个函数可以查找多个图片,并且返回所有找到的图像的坐标.

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

pic_name 字符串:图片名,可以是多个图片,比如"test.bmp|test2.bmp|test3.bmp"

delta_color 字符串:颜色色偏比如"203040" 表示RGB的色偏分别是20 30 40 (这里是16进制表示)

sim 双精度浮点数:相似度,取值范围0.1-1.0

dir 整形数:查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左, 从下到上

返回值:

字符串: 返回的是所有找到的坐标格式如下:"id,x,y|id,x,y..|id,x,y" (图片左上角的坐标)

比如"0,100,20|2,30,40" 表示找到了两个,第一个,对应的图片是图像序号为0的图片,坐标是(100,20),第二个是序号为2的图片,坐标(30,40) (由于内存限制,返回的图片数量最多为1500个左右)
*/
func (com *opsoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.op.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

/*
函数简介:

获取(x,y)的颜色,颜色返回格式"RRGGBB",注意,和按键的颜色格式相反

参数定义:

x 整形数:X坐标

y 整形数:Y坐标

返回值:

字符串: 颜色字符串(注意这里都是小写字符，和工具相匹配)

比如"0,100,20|2,30,40" 表示找到了两个,第一个,对应的图片是图像序号为0的图片,坐标是(100,20),第二个是序号为2的图片,坐标(30,40) (由于内存限制,返回的图片数量最多为1500个左右)
*/
func (com *opsoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.op.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

/*
函数简介:

抓取上次操作的图色区域，保存为file(32位位图)

参数定义:

file 字符串:保存的文件名,保存的地方一般为SetPath中设置的目录

返回值:

整形数: 0:失败 1:成功
*/
func (com *opsoft) CapturePre(file string) int {
	ret, _ := com.op.CallMethod("CapturePre", file)
	return int(ret.Val)
}

/*
函数简介:

开启图色调试模式，此模式会稍许降低图色和文字识别的速度.默认不开启.

参数定义:

enable_debug 整形数: 0 为关闭 1 为开启

返回值:

整形数: 0:失败 1:成功
*/
func (com *opsoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.op.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

/*
函数简介:

获取指定区域的图像,用二进制数据的方式返回,（不适合按键使用）方便二次开发.

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

返回值:

整形数: 返回的是指定区域的二进制颜色数据地址,每个颜色是4个字节,表示方式为(BBGGRR00)
*/
func (com *opsoft) GetScreenData(x1, y1, x2, y2 int) int {
	ret, _ := com.op.CallMethod("GetScreenData", x1, y1, x2, y2)
	return int(ret.Val)
}

/*
函数简介:

获取指定区域的图像,用24位位图的数据格式返回,方便二次开发.（或者可以配合SetDisplayInput的mem模式）

参数定义:

x1 整形数:区域的左上X坐标

y1 整形数:区域的左上Y坐标

x2 整形数:区域的右下X坐标

y2 整形数:区域的右下Y坐标

data 变参指针:返回图片的数据指针

size 变参指针:返回图片的数据长度

返回值:

整形数: 0 : 失败 1 : 成功
*/
func (com *opsoft) GetScreenDataBmp(x1, y1, x2, y2 int, data, size *int) int {
	d := ole.NewVariant(ole.VT_I4, 0)
	s := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("GetScreenDataBmp", x1, y1, x2, y2, &data, &size)
	*data = int(d.Val)
	*size = int(s.Val)
	d.Clear()
	s.Clear()
	return int(ret.Val)
}

/*
函数简介:

设定图色的获取方式，默认是显示器或者后台窗口(具体参考BindWindow)

参数定义:

mode 字符串: 图色输入模式取值有以下几种

"screen" 这个是默认的模式，表示使用显示器或者后台窗口

"pic:file" 指定输入模式为指定的图片,如果使用了这个模式，则所有和图色相关的函数均视为对此图片进行处理，比如文字识别查找图片 颜色 等等一切图色函数.

需要注意的是，设定以后，此图片就已经加入了缓冲，如果更改了源图片内容，那么需要 释放此缓冲，重新设置.

"mem:addr" 指定输入模式为指定的图片,此图片在内存当中. addr为图像内存地址,一般是GetScreenDataBmp的返回值（前54字节为位图信息，后面的是像素数据），注意与大漠的区别.
如果使用了这个模式，则所有和图色相关的函数,均视为对此图片进行处理.比如文字识别 查找图片 颜色 等等一切图色函数所有坐标都相对此图片，如果不想受到影响，调用GetScreenDataBmp时应时整个窗口的大小.

返回值:

整形数: 0 : 失败 1 : 成功
*/
func (com *opsoft) SetDisplayInput(mode string) int {
	ret, _ := com.op.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}
