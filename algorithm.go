package op

func (com *Opsoft) AStarFindPath(mapWidth, mapHeight int, disable_points string, beginX, beginY, endX, endY int) int {
	ret, _ := com.op.CallMethod("AStarFindPath", mapWidth, mapHeight, disable_points, beginX, beginY, endX, endY)
	return int(ret.Val)
}
