package models

import "gorm.io/gorm"

// 下载的文件信息

type Files struct {
	gorm.Model
	Name    string `gorm:"index:name_md5" json:"name"`    // 文件名称
	MD5     string `gorm:"index:name_md5" json:"md5"`     // 文件的md5值
	Process string `gorm:"index:name_md5" json:"process"` // 文件下载进度
	Path    string `json:"path"`                          // 文件存储地址
}
