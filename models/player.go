package models

import "gin_ranking/dao"

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
}

//// 解决插入表的时候会自动添加复数的问题， 但现在已经可以通过在dao init()中设置SingularTable解决，不用再写TableName()
//func (Player) TableName() string {
//	return "player"
//}

func GetPlayers(aid int) ([]Player, error) {
	var players []Player
	err := dao.DBEngine.Where("aid = ?", aid).Find(&players).Error
	return players, err
}
