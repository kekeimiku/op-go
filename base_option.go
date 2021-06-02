package op

/*
函数简介:

获取注册在系统中的op.dll的路径.

返回值:

字符串: 返回op.dll所在路径
*/
func (com *opsoft) GetBasePath() string {
	ret, _ := com.op.CallMethod("GetBasePath")
	return ret.ToString()
}

/*
函数简介:

返回当前大漠对象的ID值，这个值对于每个对象是唯一存在的。可以用来判定两个大漠对象是否一致.

返回值:

整形数: 当前对象的ID值.
*/
func (com *opsoft) GetID() int {
	ret, _ := com.op.CallMethod("GetID")
	return int(ret.Val)
}

/*
函数简介:

获取插件命令的最后错误

返回值:

整形数: 返回值表示错误值。 参考msdn的GetLastError

注: 此函数必须紧跟上一句函数调用，中间任何的语句调用都会改变这个值.
*/
func (com *opsoft) GetLastError() int {
	ret, _ := com.op.CallMethod("GetLastError")
	return int(ret.Val)
}

/*
函数简介:

获取全局路径.(可用于调试)

返回值:

字符串: 以字符串的形式返回当前设置的全局路径
*/
func (com *opsoft) GetPath() string {
	ret, _ := com.op.CallMethod("GetPath")
	return ret.ToString()
}

/*
函数简介:

设置全局路径,设置了此路径后,所有接口调用中,相关的文件都相对于此路径. 比如图片,字库等.

参数定义:

path 字符串: 路径,可以是相对路径,也可以是绝对路径

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetPath(path string) int {
	ret, _ := com.op.CallMethod("SetPath", path)
	return int(ret.Val)
}

/*
函数简介:

设置是否弹出错误信息,默认是打开.

参数定义:

show 整形数: 0表示不打开,1表示打开，2表示将错误信息写入文件

返回值:

整形数: 0: 失败 1: 成功
*/
func (com *opsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.op.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

/*
函数简介:

返回当前插件版本号

返回值:

字符串: 当前插件的版本描述字符串
*/
func (com *opsoft) Ver() string {
	ver, _ := com.op.CallMethod("Ver")
	return ver.ToString()
}

/*
函数简介:

设置是否开启或者关闭插件内部的图片缓存机制. (默认是打开).

参数定义:

enable 整形数: 0 : 关闭 1 : 打开

返回值:

整形数: 0: 失败, 1: 成功

注: 有些时候，系统内存比较吃紧，这时候再打开内部缓存，可能会导致缓存分配在虚拟内存，这样频繁换页，反而导致图色效率下降.这时候就建议关闭图色缓存. 所有图色缓存机制都是对本对象的，也就是说，调用图色缓存机制的函数仅仅对本对象生效. 每个对象都有一个图色缓存队列.
*/
func (com *opsoft) EnablePicCache(enable int) int {
	ret, _ := com.op.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}
