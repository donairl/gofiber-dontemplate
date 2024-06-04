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

func FindUserByID(id uint) *User {

	user := &User{}
	database.Connection.First(user, id)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	return user
}
