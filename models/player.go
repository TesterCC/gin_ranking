package models

import (
	"gin_ranking/dao"
	"gorm.io/gorm"
)

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

//func GetPlayers(aid int) ([]Player, error) {
//	var players []Player
//	err := dao.DBEngine.Where("aid = ?", aid).Find(&players).Error
//	return players, err
//}

func GetPlayers(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.DBEngine.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

func GetPlayerInfo(id int) (Player, error) {
	var player Player
	err := dao.DBEngine.Where("id = ?", id).First(&player).Error
	return player, err
}

func UpdatePlayerScore(id int) {
	var player Player

	// https://gorm.io/zh_CN/docs/update.html#%E4%BD%BF%E7%94%A8-SQL-%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%9B%B4%E6%96%B0
	// db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	dao.DBEngine.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
