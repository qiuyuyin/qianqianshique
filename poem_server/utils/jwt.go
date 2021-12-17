package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"poem_server/global"
	"poem_server/model/system/sys_request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.POEM_CONFIG.JWT.SigningKey),
	}
}

// 创建一个token
func (j JWT) CreateToken(claims sys_request.UserToken) (string, error) {
	// 根据加密方式新建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//根据签名新建tokenString
	return token.SignedString(j.SigningKey)
}

// 更换新的token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims sys_request.UserToken) (string, error) {
	// 更新签名
	v, err, _ := global.POEM_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

//解析token
func (j JWT) ParseToken(tokenString string) (*sys_request.UserToken, error) {
	// 通过claims进行解析
	token, err := jwt.ParseWithClaims(tokenString, &sys_request.UserToken{}, func(token *jwt.Token) (i interface{}, e error) {
		// 解析返回签名
		return j.SigningKey, nil
	})
	// 进行验证
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*sys_request.UserToken); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}
