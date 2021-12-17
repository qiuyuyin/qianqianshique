package system

import (
	"github.com/satori/go.uuid"
	"poem_server/model/global"
)

// SysUser 用户的详细信息
type SysUser struct {
	global.CRM_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`                    // 用户UUID
	Username    string       `json:"userName" gorm:"comment:用户登录名"`                 // 用户登录名
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`                      // 用户登录密码
	NickName    string       `json:"nickName" gorm:"default:新用户;comment:用户昵称"`      // 用户昵称
	Mail        string       `json:"mail" gorm:"comment:用户邮箱"`                      // 用户昵称
	HeaderImg   string       `json:"headerImg" gorm:"comment:用户头像"`                 // 用户头像
	AuthorityId string       `json:"authorityId" gorm:"default:250;comment:用户角色ID"` // 用户角色ID
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}
