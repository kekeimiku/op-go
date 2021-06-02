package op

/*
运行指定的应用程序.

参数定义:

path string类型 可执行程序路径

mode int类型 模式 0普通模式，1加强模式

返回值:

int类型 0失败，1成功
*/
func (com *opsoft) RunApp(path string, mode int) int {
	ret, _ := com.op.CallMethod("RunApp", path, mode)
	return int(ret.Val)
}

/*
运行指定的应用程序.

参数定义:

path string类型 可执行程序路径

dis int类型 模式 0隐藏，1用最近的大小和位置显示,

返回值:

int类型 0失败，1成功
*/
func (com *opsoft) WinExec(path string, dis int) int {
	ret, _ := com.op.CallMethod("WinExec", path, dis)
	return int(ret.Val)
}

/*
运行指定的应用程序.

参数定义：

path string类型 可执行程序路径

time int类型 等待时间 毫秒

返回值:

string类型 cmd输出的字符
*/
func (com *opsoft) GetCmdStr(path string, time int) string {
	ret, _ := com.op.CallMethod("GetCmdStr", path, time)
	return ret.ToString()
}
