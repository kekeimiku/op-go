package op

/*
函数简介:

绑定指定的窗口,并指定这个窗口的屏幕颜色获取方式,鼠标仿真模式,键盘仿真模式,以及模式设定.

参数定义:

hwnd 整形数: 指定的窗口句柄

display 字符串: 屏幕颜色获取方式 取值有以下几种

normal" : 正常模式,平常我们用的前台截屏模式
"gdi" : gdi模式,用于窗口采用GDI方式刷新时. 此模式占用CPU较大. 参考SetAero
"dx" : dx模式,等同于dx.d3d9
"dx.d3d9" dx模式，使用d3d9渲染
"dx.d3d10" dx模式，使用d3d10渲染
"dx.d3d11" dx模式，使用d3d11渲染
"opengl" opengl模式，使用opengl渲染的窗口，支持最新版雷电模拟器，以及夜神6.1，支持最小化窗口截图
"opengl.nox" opengl模式，针对最新夜神模拟器的渲染方式，测试中。。。
mouse 字符串: 鼠标仿真模式 取值有以下几种

"normal" : 正常模式,平常我们用的前台鼠标模式

"windows": Windows模式,采取模拟windows消息方式 同按键自带后台插件.

keypad 字符串: 键盘仿真模式 取值有以下几种

"normal" : 正常模式,平常我们用的前台键盘模式

"windows": Windows模式,采取模拟windows消息方式 同按键的后台插件.

mode 整形数: 模式

返回值:

整形数: 0:失败 1:成功
*/
func (com *opsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.op.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

/*
函数简介:

解除绑定窗口,并释放系统资源.一般在OnScriptExit调用.

返回值:

整形数: 0:失败 1:成功
*/
func (com *opsoft) UnBindWindow() int {
	ret, _ := com.op.CallMethod("UnBindWindow")
	return int(ret.Val)
}
