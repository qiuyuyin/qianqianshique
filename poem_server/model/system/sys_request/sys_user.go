package sys_request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Login struct {
	Username string `json:"username" binding:"required,max=50"`       // 用户名
	Password string `json:"password" binding:"required,min=6,max=50"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

type Register struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Nickname string `json:"nickname" binding:"required,min=1,max=50"`
	Mail     string `json:"email"`
}

//相当于删去密码加上一个jwt的User
type UserToken struct {
	ID          uint
	UUID        uuid.UUID
	Username    string
	NickName    string
	Mail        string
	AuthorityId string
	BufferTime  int64

	jwt.StandardClaims
}
