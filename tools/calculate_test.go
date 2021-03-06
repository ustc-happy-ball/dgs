package tools

import (
	"dgs/configs"
	"fmt"
	"math"
	"testing"
)

func TestGet(t *testing.T) {
	configs.MapMinX = 0
	configs.MapMaxX = 300
	configs.MapMinY = 0
	configs.MapMaxY = 300
	configs.TileSize = 1                                                                                       //地图瓦片大小，一个瓦片对应一个地图坐标
	configs.TowerRadius = 50                                                                                   // 灯塔AOI半径
	configs.TowerDiameter = configs.TowerRadius * 2                                                            //灯塔AOI直径
	configs.PlayerRange = 100                                                                                  // 玩家视野半径
	configs.TowerCols = int32(math.Ceil(float64((configs.MapMaxX - configs.MapMinX) / configs.TowerDiameter))) // 整个地图中有多少列Tower 从1开始
	configs.TowerRows = int32(math.Ceil(float64((configs.MapMaxY - configs.MapMinY) / configs.TowerDiameter))) // 整个地图中有多少行Tower 从1开始

	result := GetOtherTowers(8)
	fmt.Println(result)
}
