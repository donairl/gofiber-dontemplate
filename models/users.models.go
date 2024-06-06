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

func UserFindByEmail(email string) *User {

	user := &User{}
	database.Connection.Where("email = ?", email).First(user)
	//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	return user
}

func UserCreate(user *User) *gorm.DB {

	result := database.Connection.Create(user)
	return result

}

func UserUpdate(id uint, user User) *gorm.DB {
	result := database.Connection.Model(&User{}).Where("id = ?", id).Updates(user)
	return result
}

func UserDelete(id uint) *gorm.DB {
	result := database.Connection.Delete(&User{}, id)
	return result
}

// Function to display all users
func UserFindAll() ([]User, error) {
	users := []User{}
	err := database.Connection.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
