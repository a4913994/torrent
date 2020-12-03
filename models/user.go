package models

import (
	"downloader/db"
	"downloader/forms"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	BirthDay string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoUrl string `json:"photo_url"`
	Active   bool   `json:"active"`
}

func (h User) Signup(userPayload forms.UserSignup) uint {
	d := db.GetDB()
	user := User{
		Name:     userPayload.Name,
		BirthDay: userPayload.BirthDay,
		Gender:   userPayload.Gender,
		PhotoUrl: userPayload.PhotoURL,
		Active:   true,
	}
	result := d.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user.ID
}

func (h User) GetByID(id uint) (*User, error) {
	d := db.GetDB()
	user := &User{}
	user.ID = id
	result := d.First(user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, result.Error
}
