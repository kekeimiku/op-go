package op

import (
	"github.com/go-ole/go-ole"
)

/*
在屏幕范围(x1,y1,x2,y2)内,查找string(可以是任意个字符串的组合),并返回符合color_format的坐标位置,相似度sim同Ocr接口描述.

参数定义:

x1 int 区域的左上X坐标.

y1 int 区域的左上Y坐标.

x2 int 区域的右下X坐标.

y2 int 区域的右下Y坐标.

str string 待查找的字符串,可以是字符串组合，比如"长安|洛阳|大雁塔",中间用"|"来分割字符串.

color_format string 颜色格式串, 可以包含换行分隔符,语法是","后加分割字符串. 具体可以查看下面的示例.

注意，RGB和HSV格式都支持.

sim 双精度浮点数:相似度,取值范围0.1-1.0

intX 变参指针:返回X坐标没找到返回-1

intY 变参指针:返回Y坐标没找到返回-1

返回值:

返回字符串的索引

int 没找到返回-1, 比如"长安|洛阳",若找到长安，则返回0
*/
func (com *opsoft) FindStr(x1, y1, x2, y2 int, str string, color_format string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.op.CallMethod("FindStr", x1, y1, x2, y2, str, color_format, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

/*
在屏幕范围(x1,y1,x2,y2)内,查找string(可以是任意个字符串的组合),并返回符合color_format的坐标位置,相似度sim同Ocr接口描述.

参数定义:

x1 int 区域的左上X坐标.

y1 int 区域的左上Y坐标.

x2 int 区域的右下X坐标.

y2 int 区域的右下Y坐标.

str string 待查找的字符串,可以是字符串组合，比如"长安|洛阳|大雁塔",中间用"|"来分割字符串.

color_format string 颜色格式串, 可以包含换行分隔符,语法是","后加分割字符串. 具体可以查看下面的示例.

注意，RGB和HSV格式都支持.

sim 双精度浮点数:相似度,取值范围0.1-1.0


返回值:

返回所有找到的坐标集合,格式如下: "id,x0,y0|id,x1,y1|......|id,xn,yn"

比如"0,100,20|2,30,40" 表示找到了两个,第一个,对应的是序号为0的字符串,坐标是(100,20),第二个是序号为2的字符串,坐标(30,40)

int 没找到返回-1, 比如"长安|洛阳",若找到长安，则返回0
*/
func (com *opsoft) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.op.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

/*
设置字库文件

参数定义:

index int 字库的序号,取值为0-19,目前最多支持20个字库
file string 字库文件名

返回值:

int 0成功，1失败

注: 此函数速度很慢，全局初始化时调用一次即可，切换字库用UseDict
*/
func (com *opsoft) SetDict(index int, file string) int {
	ret, _ := com.op.CallMethod("SetDict", index, file)
	return int(ret.Val)
}
