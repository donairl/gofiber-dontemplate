package models

import (
	"time"

	"github.com/donairl/gofiber-dontemplate/lib/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname     string `json:"fullname" gorm:"type:varchar(80)"`
	Email        string `json:"-"`
	PasswordHash string `json:"-"`
	Role         uint   `json:"role" gorm:"type:integer"`
	Birthday     *time.Time
}

func UserFindByID(id uint) *User {

	user := &User{}
	database.Connection.First(user, id)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	return user
}

func UserFindByEmail(email uint) *User {

	user := &User{}
	database.Connection.First(user, id)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	return user
}

func UserSave(user User) *gorm.DB {

	result := database.Connection.Create(user)
	return result

}
