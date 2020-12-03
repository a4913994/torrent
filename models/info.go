package models

import (
	"downloader/db"
	"gorm.io/gorm"
	"log"
)

// 下载信息表

type Info struct {
	gorm.Model
	UserName string `json:"user_name"` // 用户名称
	Torrent  string `json:"torrent"`   // 种子地址
}

func (i Info) GetDownloadFiles(userName string) []Info {
	d := db.GetDB()
	var infos []Info
	result := d.Where(&Info{UserName: userName}).Find(&infos)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return infos
}
