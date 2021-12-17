package sys_responce

import "poem_server/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}
