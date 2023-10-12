package models

import (
	"gin_ranking/dao"
)

type User struct {
	Id         int `json:"id"`
	Username   string `json:"username"`

	//Id         string `gorm:"type:int(20); not null" json:"id"`
	//Name       string `gorm:"type:varchar(20); not null" json:"name"`
	//Password   string `json:"password"`
	//AddTime    int64  `json:"addTime"`
	//UpdateTime int64  `json:"updateTime"`
	//Status     string `gorm:"type:varchar(20); not null" json:"status" binding:"required"`
	//Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
	//Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
	//Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
}

// custom generate table name by dao model, but now can just set SingularTable in dao.go
func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	// report error
	err := dao.DBEngine.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.DBEngine.Create(&user).Error
	return user.Id, err
}
