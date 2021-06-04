package op

/*
函数简介:

A星算法（最短路径）

参数定义:

mapWidth 整形数:区域的左上X坐标

mapHeight 整形数:区域的左上Y坐标

disable_points 字符串:不可通行的坐标，以"|"分割，例如:"10,15|20,30"

beginX 整形数:源坐标X

beginY 整形数:源坐标Y

endX 整形数:目的坐标X

endY 整形数:目的坐标Y

返回值:

int 0 失败，1 成功
*/
func (com *Opsoft) AStarFindPath(mapWidth, mapHeight int, disable_points string, beginX, beginY, endX, endY int) int {
	ret, _ := com.op.CallMethod("AStarFindPath", mapWidth, mapHeight, disable_points, beginX, beginY, endX, endY)
	return int(ret.Val)
}
