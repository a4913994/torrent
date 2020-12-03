package models

import "downloader/db"

func Init() {
	d := db.GetDB()
	d.AutoMigrate(&User{}, &Files{}, &Info{})
}
