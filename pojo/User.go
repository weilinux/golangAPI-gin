package pojo

import (
	"github.com/weilinux/golangAPI-gin/database"
	"log"
)

//结构体定义及其在database命名规则
type User struct { //DB_name: users
	Id int64 `json:"UserId"` //默认字段名称: DB_TABLE_fields: id, UserId DB_TABLE_fields : user_id
	Name string `json:"UserName"` //设置字段名称 `gorm:"Column:username"`
	Password string `json:"UserPassword"`
	Email string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var Users []User
	database.DBConnect.Find(&Users)
	return Users
}

func FindUserById(userId string) User {
	var user User
	database.DBConnect.Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user User) User {
	database.DBConnect.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBConnect.Where("id = ?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId string, user User) User {
	database.DBConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}