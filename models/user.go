package models

import (
	"gin_ranking/dao"
)

type User struct {
	//gorm.Model        // 默认会添加主键等字段
	Id   string `gorm:"type:int(20); not null" json:"id"`
	Name string `gorm:"type:varchar(20); not null" json:"name"`
	//Status     string `gorm:"type:varchar(20); not null" json:"status" binding:"required"`
	//Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
	//Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
	//Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
}

//// custom generate table name by dao model, but now can just set SingularTable in dao.go
//func (User) TableName() string {
//	return "user"
//}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.DB.Where("id = ?", id).First(&user).Error
	//dao.DB.Select("id").Where("id=?", id).Find(&user)
	return user, err
}
