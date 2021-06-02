package op

import "github.com/go-ole/go-ole"

/*
函数简介:

把窗口坐标转换为屏幕坐标

参数定义:

hwnd 整形数: 指定的窗口句柄
x 变参指针: 窗口X坐标
y 变参指针: 窗口Y坐标

返回值:

int 0 失败，1 成功
*/
func (com *opsoft) ClientToScreen(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("ClientToScreen", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

/*
函数简介:

根据进程名枚举进程

参数定义:

name 字符串:进程名,比如qq.exe

返回值:

字符串 : 返回所有匹配的进程PID,并按打开顺序排序,格式"pid1,pid2,pid3"
*/
func (com *opsoft) EnumProcess(name string) string {
	ret, _ := com.op.CallMethod("EnumProcess", name)
	return ret.ToString()
}

/*
函数简介:

根据指定条件,枚举系统中符合条件的窗口,可以枚举到按键自带的无法枚举到的窗口.

参数定义:

parent 整形数: 获得的窗口句柄是该窗口的子窗口的窗口句柄,取0时为获得桌面句柄

title 字符串: 窗口标题. 此参数是模糊匹配.

class_name 字符串: 窗口类名. 此参数是模糊匹配.

filter整形数: 取值定义如下

1 : 匹配窗口标题,参数title有效

2 : 匹配窗口类名,参数class_name有效.

4 : 只匹配指定父窗口的第一层孩子窗口

8 : 匹配所有者窗口为0的窗口,即顶级窗口

16 : 匹配可见的窗口

32 : 匹配出的窗口按照窗口打开顺序依次排列

这些值可以相加,比如4+8+16就是类似于任务管理器中的窗口列表

返回值:

字符串 : 返回所有匹配的窗口句柄字符串,格式"hwnd1,hwnd2,hwnd3"
*/
func (com *opsoft) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.op.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

/*
函数简介:

根据指定进程pid以及其它条件,枚举系统中符合条件的窗口,可以枚举到按键自带的无法枚举到的窗口

参数定义:

pid 整形数: 进程pid.

title 字符串: 窗口标题. 此参数是模糊匹配.

class_name 字符串: 窗口类名. 此参数是模糊匹配.

filter 整形数: 取值定义如下

1 : 匹配窗口标题,参数title有效

2 : 匹配窗口类名,参数class_name有效

8 : 匹配所有者窗口为0的窗口,即顶级窗口

16 : 匹配可见的窗口

这些值可以相加,比如2+8+16

返回值:

字符串: 返回所有匹配的窗口句柄字符串,格式"hwnd1,hwnd2,hwnd3"
*/
func (com *opsoft) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.op.CallMethod("EnumWindowByProcessId", pid, title, className, filter)
	return ret.ToString()
}

/*
函数简介:

根据两组设定条件来枚举指定窗口.

参数定义:

spec1 字符串: 查找串1. (内容取决于flag1的值)

flag1整形数: 取值如下:

0表示spec1的内容是标题

1表示spec1的内容是程序名字. (比如notepad)

2表示spec1的内容是类名

3表示spec1的内容是程序路径.(不包含盘符,比如\windows\system32)

4表示spec1的内容是父句柄.(十进制表达的串)

5表示spec1的内容是父窗口标题

6表示spec1的内容是父窗口类名

7表示spec1的内容是顶级窗口句柄.(十进制表达的串)

8表示spec1的内容是顶级窗口标题

9表示spec1的内容是顶级窗口类名

type1 整形数: 取值如下

0精确判断

1模糊判断

spec2 字符串: 查找串2. (内容取决于flag2的值)

flag2 整形数: 取值如下:

0表示spec2的内容是标题

1表示spec2的内容是程序名字. (比如notepad)

2表示spec2的内容是类名

3表示spec2的内容是程序路径.(不包含盘符,比如\windows\system32)

4表示spec2的内容是父句柄.(十进制表达的串)

5表示spec2的内容是父窗口标题

6表示spec2的内容是父窗口类名

7表示spec2的内容是顶级窗口句柄.(十进制表达的串)

8表示spec2的内容是顶级窗口标题

9表示spec2的内容是顶级窗口类名

type2 整形数: 取值如下

0精确判断

1模糊判断

sort 整形数: 取值如下

0不排序.

1对枚举出的窗口进行排序,按照窗口打开顺序.

返回值:

字符串: 返回所有匹配的窗口句柄字符串,格式"hwnd1,hwnd2,hwnd3"
*/
func (com *opsoft) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.op.CallMethod("EnumWindowSuper", spec1, flag1, type1, spec2, flag2, type2, sort)
	return ret.ToString()
}

/*
函数简介:

根据指定的进程名字，来查找可见窗口.

参数定义:

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是模糊匹配.

title 字符串: 窗口标题,如果为空，则匹配所有.这里的匹配是模糊匹配.

返回值:

整形数: 整形数表示的窗口句柄，没找到返回0.
*/
func (com *opsoft) FindWindow(class, title string) int {
	ret, _ := com.op.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

/*
函数简介:

根据指定的进程名字，来查找可见窗口.

参数定义:

process_name 字符串: 进程名. 比如(notepad.exe).这里是精确匹配,但不区分大小写.

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是模糊匹配.

title 字符串: 窗口标题,如果为空，则匹配所有.这里的匹配是模糊匹配.

返回值:

整形数: 整形数表示的窗口句柄，没找到返回0.
*/
func (com *opsoft) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

/*
函数简介:

根据指定的进程Id，来查找可见窗口.

参数定义:

process_id 整形数: 进程id.

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是模糊匹配.

title 字符串: 窗口标题,如果为空，则匹配所有.这里的匹配是模糊匹配.

返回值:

整形数: 整形数表示的窗口句柄，没找到返回0.
*/
func (com *opsoft) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowByProcessId", processId, class, title)
	return int(ret.Val)
}

/*
函数简介:

查找符合类名或者标题名的顶层可见窗口,如果指定了parent,则在parent的第一层子窗口中查找.

参数定义:

parent 整形数: 父窗口句柄，如果为空，则匹配所有顶层窗口

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是模糊匹配.

title 字符串: 窗口标题,如果为空，则匹配所有. 这里的匹配是模糊匹配.

返回值:

整形数: 整形数表示的窗口句柄，没找到返回0.
*/
func (com *opsoft) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.op.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

/*
函数简介:

根据两组设定条件来查找指定窗口.

参数定义:

spec1 字符串: 查找串1. (内容取决于flag1的值)

flag1整形数: 取值如下:

0表示spec1的内容是标题

1表示spec1的内容是程序名字. (比如notepad)

2表示spec1的内容是类名

3表示spec1的内容是程序路径.(不包含盘符,比如\windows\system32)

4表示spec1的内容是父句柄.(十进制表达的串)

5表示spec1的内容是父窗口标题

6表示spec1的内容是父窗口类名

7表示spec1的内容是顶级窗口句柄.(十进制表达的串)

8表示spec1的内容是顶级窗口标题

9表示spec1的内容是顶级窗口类名

type1 整形数: 取值如下

0精确判断

1模糊判断

spec2 字符串: 查找串2. (内容取决于flag2的值)

flag2 整形数: 取值如下:

0表示spec2的内容是标题

1表示spec2的内容是程序名字. (比如notepad)

2表示spec2的内容是类名

3表示spec2的内容是程序路径.(不包含盘符,比如\windows\system32)

4表示spec2的内容是父句柄.(十进制表达的串)

5表示spec2的内容是父窗口标题

6表示spec2的内容是父窗口类名

7表示spec2的内容是顶级窗口句柄.(十进制表达的串)

8表示spec2的内容是顶级窗口标题

9表示spec2的内容是顶级窗口类名

type2 整形数: 取值如下

0精确判断

1模糊判断

返回值:

整形数: 整形数表示的窗口句柄，没找到返回0
*/
func (com *opsoft) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.op.CallMethod("FindWindowSuper", spec1, flag1, type1, spec2, flag2, type2)
	return int(ret.Val)
}

/*
函数简介:

获取窗口客户区域在屏幕上的位置.

参数定义:

hwnd 整形数: 指定的窗口句柄

x1 变参指针: 返回窗口客户区左上角X坐标

y1 变参指针: 返回窗口客户区左上角Y坐标

x2 变参指针: 返回窗口客户区右下角X坐标

y2 变参指针: 返回窗口客户区右下角Y坐标

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
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

/*
函数简介:

获取窗口客户区域的宽度和高度.

参数定义:

hwnd 整形数: 指定的窗口句柄

width 变参指针: 宽度

height 变参指针: 高度

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.op.CallMethod("GetClientSize", hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	pWidth.Clear()
	pHeight.Clear()
	return int(ret.Val)
}

/*
函数简介:

获取顶层活动窗口中具有输入焦点的窗口句柄.

返回值:

整形数: 返回整型表示的窗口句柄
*/
func (com *opsoft) GetForegroundFocus() int {
	ret, _ := com.op.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

/*
函数简介:

获取顶层活动窗口,可以获取到按键自带插件无法获取到的句柄

返回值:

整形数: 返回整型表示的窗口句柄
*/
func (com *opsoft) GetForegroundWindow() int {
	ret, _ := com.op.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

/*
函数简介:

获取鼠标指向的可见窗口句柄,可以获取到按键自带的插件无法获取到的句柄.

返回值:

整形数: 返回整型表示的窗口句柄
*/
func (com *opsoft) GetMousePointWindow() int {
	ret, _ := com.op.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

/*
函数简介:

获取给定坐标的可见窗口句柄,可以获取到按键自带的插件无法获取到的句柄.

参数定义:

X 整形数: 屏幕X坐标.

Y 整形数: 屏幕Y坐标.

返回值:

整形数: 返回整型表示的窗口句柄.
*/
func (com *opsoft) GetPointWindow(x, y int) int {
	ret, _ := com.op.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

/*
函数简介:

根据指定的pid获取进程详细信息,(进程名,进程全路径,CPU占用率(百分比),内存占用量(字节))

参数定义:

pid 整形数: 进程pid.

返回值:

字符串: 格式"进程名|进程路径|cpu|内存"
*/
func (com *opsoft) GetProcessInfo(pid int) string {
	ret, _ := com.op.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

/*
函数简介:

获取特殊窗口.

参数定义:

Flag 整形数: 取值定义如下

0 : 获取桌面窗口

1 : 获取任务栏窗口

返回值:

整形数: 以整型数表示的窗口句柄
*/
func (com *opsoft) GetSpecialWindow(flag int) int {
	ret, _ := com.op.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

/*
函数简介:

获取窗口的类名.

参数定义:

hwnd 整形数: 指定的窗口句柄

返回值:

字符串: 窗口的类名
*/
func (com *opsoft) GetWindowClass(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

/*
函数简介:

获取指定窗口所在的进程ID.

参数定义:

hwnd 整形数: 窗口句柄

返回值:

整形数: 返回整型表示的是进程ID
*/
func (com *opsoft) GetWindowProcessId(hwnd int) int {
	ret, _ := com.op.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

/*
函数简介:

获取指定窗口所在的进程的exe文件全路径.

参数定义:

hwnd 整形数: 窗口句柄

返回值:

字符串: 返回字符串表示的是exe全路径名
*/
func (com *opsoft) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

/*
函数简介:

获取窗口在屏幕上的位置

参数定义:

hwnd 整形数: 指定的窗口句柄

x1 变参指针: 返回窗口左上角X坐标

y1 变参指针: 返回窗口左上角Y坐标

x2 变参指针: 返回窗口右下角X坐标

y2 变参指针: 返回窗口右下角Y坐标

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
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

/*
函数简介:

获取指定窗口的一些属性

参数定义:

hwnd 整形数: 指定的窗口句柄

flag 整形数: 取值定义如下

0 : 判断窗口是否存在

1 : 判断窗口是否处于激活

2 : 判断窗口是否可见

3 : 判断窗口是否最小化

4 : 判断窗口是否最大化

5 : 判断窗口是否置顶

6 : 判断窗口是否无响应

7 : 判断窗口是否可用(灰色为不可用)

8 : 另外的方式判断窗口是否无响应,如果6无效可以尝试这个

返回值:

整形数: 0: 不满足条件 1: 满足条件
*/
func (com *opsoft) GetWindowState(hwnd, flag int) int {
	ret, _ := com.op.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

获取窗口的标题

参数定义:

hwnd 整形数: 指定的窗口句柄

返回值:

字符串: 窗口的标题
*/
func (com *opsoft) GetWindowTitle(hwnd int) string {
	ret, _ := com.op.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

/*
函数简介:

移动指定窗口到指定位置

参数定义:

hwnd 整形数: 指定的窗口句柄

x 整形数: X坐标

y 整形数: Y坐标

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.op.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

/*
函数简介:

把屏幕坐标转换为窗口坐标

参数定义:

hwnd 整形数: 指定的窗口句柄

x 变参指针: 屏幕X坐标

y 变参指针: 屏幕Y坐标

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.op.CallMethod("ScreenToClient", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

/*
函数简介:

向指定窗口发送粘贴命令. 把剪贴板的内容发送到目标窗口.

参数定义:

hwnd 整形数: 指定的窗口句柄

返回值:

整形数: 0: 失败 1: 成功

注:剪贴板是公共资源，多个线程同时设置剪贴板时,会产生冲突，必须用互斥信号保护.
*/
func (com *opsoft) SendPaste(hwnd int) int {
	ret, _ := com.op.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

/*
函数简介:

向指定窗口发送文本数据

参数定义:

hwnd 整形数: 指定的窗口句柄

str 字符串: 发送的文本数据

返回值:

整形数: 0: 失败 1: 成功

注： 有时候发送中文，可能会大部分机器正常，少部分会乱码。这种情况一般有两个可能

系统编码没有设置为GBK

目标程序里可能安装了改变当前编码的软件，比如常见的是输入法. （尝试卸载）
*/
func (com *opsoft) SendString(hwnd int, str string) int {
	ret, _ := com.op.CallMethod("SendString", hwnd, str)
	return int(ret.Val)
}

/*
函数简介:

设置窗口客户区域的宽度和高度

参数定义:

hwnd 整形数: 指定的窗口句柄

width 整形数: 宽度

height 整形数: 高度

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.op.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的大小

参数定义:

hwnd 整形数: 指定的窗口句柄

width 整形数: 宽度

height 整形数: 高度

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.op.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的状态

参数定义:

hwnd 整形数: 指定的窗口句柄

flag 整形数: 取值定义如下

0 : 关闭指定窗口

1 : 激活指定窗口

2 : 最小化指定窗口,但不激活

3 : 最小化指定窗口,并释放内存,但同时也会激活窗口.(释放内存可以考虑用FreeProcessMemory函数)

4 : 最大化指定窗口,同时激活窗口.

5 : 恢复指定窗口 ,但不激活

6 : 隐藏指定窗口

7 : 显示指定窗口

8 : 置顶指定窗口

9 : 取消置顶指定窗口

10 : 禁止指定窗口

11 : 取消禁止指定窗口

12 : 恢复并激活指定窗口

13 : 强制结束窗口所在进程.

14 : 闪烁指定的窗口

15 : 使指定的窗口获取输入焦点

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetWindowState(hwnd, flag int) int {
	ret, _ := com.op.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的标题

参数定义:

hwnd 整形数: 指定的窗口句柄

titie 字符串: 标题

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetWindowText(hwnd int, title string) int {
	ret, _ := com.op.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的透明度.

参数定义:

hwnd 整形数: 指定的窗口句柄

trans 整形数: 透明度取值(0-255) 越小透明度越大 0为完全透明(不可见) 255为完全显示(不透明)

返回值:

整形数: 0: 失败 1: 成功

注 : 此接口不支持WIN98
*/
func (com *opsoft) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.op.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
