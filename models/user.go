package models

import (
	"downloader/db"
	"downloader/forms"
	"log"
	"time"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoUrl  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active"`
	UpdatedAt int64  `json:"updated_at_time"`
}

func (h User) Signup(userPayload forms.UserSignup) string {
	db := db.GetDB()
	user := User{
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gendor,
		PhotoUrl:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user.ID
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	user := &User{
		ID: id,
	}
	result := db.First(user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return user, result.Error
}
