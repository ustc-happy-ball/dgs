package aoi

import "dgs/configs"

func InitTowers() []*Tower {
	var towers []*Tower
	for i := int32(0); i < configs.TowerRows*configs.TowerCols; i++ {
		towers = append(towers, InitTower(i))
	}
	return towers
}
