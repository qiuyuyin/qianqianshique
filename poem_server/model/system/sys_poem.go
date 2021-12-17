package system

import "poem_server/model/global"

type SysPoem struct {
	global.CRM_MODEL
	AuthorID       uint   `json:"authorID" gorm:"comment:作者ID" `
	AuthorNickname string `json:"authorNickname" gorm:"comment:作者姓名" `
	AuthorUsername string `json:"authorUsername" gorm:"comment:作者姓名" `
	Title          string `json:"title" gorm:"comment:标题" `
	Content        string `json:"content"  gorm:"comment:内容信息" `
}
