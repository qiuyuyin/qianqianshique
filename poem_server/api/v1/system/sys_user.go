package system

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"poem_server/global"
	"poem_server/model/common/responce"
	"poem_server/model/system"
	"poem_server/model/system/sys_request"
	"poem_server/model/system/sys_responce"
	"poem_server/utils"
	"time"
)

type BaseApi struct {
}

func (b BaseApi) Login(c *gin.Context) {
	var login sys_request.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		responce.FailWithMessage("密码不正确", c)
		return
	}
	user := &system.SysUser{Username: login.Username, Password: login.Password}

	if err, user := userService.Login(user); err != nil {
		global.POEM_LOG.Error("登录失败！用户名不存在或者密码错误!", zap.Any("err", err))
		responce.FailWithMessage(err.Error(), c)
	} else {
		b.getToken(c, *user)
	}
}

func (b BaseApi) getToken(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.POEM_CONFIG.JWT.SigningKey)}
	claims := sys_request.UserToken{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		Mail:        user.Mail,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.POEM_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.POEM_CONFIG.JWT.ExpiresTime,
			Issuer:    "yili",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.POEM_LOG.Error("获取token失败", zap.Any("err", err))
		responce.FailWithMessage("获取token失败", c)
		return
	}
	// 不使用多个服务器进行部署
	if !global.POEM_CONFIG.System.UseMultipoint {
		responce.OkWithDetailed(sys_responce.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func (b BaseApi) Register(c *gin.Context) {
	var r sys_request.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		global.POEM_LOG.Error("注册失败", zap.Any("err", err))
		responce.FailWithMessage("输入参数错误", c)
		return
	}
	user := system.SysUser{Username: r.Username, Password: r.Password, NickName: r.Nickname, Mail: r.Mail, AuthorityId: "250"}
	if err, userReturn := userService.Register(user); err != nil {
		global.POEM_LOG.Error("注册失败", zap.Any("err", err))
		responce.FailWithMessage(err.Error(), c)
	} else {
		responce.OkWithDetailed(sys_responce.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

func (b BaseApi) FindUserPoemByPage(c *gin.Context) {
	var r sys_request.SearchUserPoem
	if err := c.ShouldBindJSON(&r); err != nil {
		global.POEM_LOG.Error("注册失败", zap.Any("err", err))
		responce.FailWithMessage("输入参数错误", c)
		return
	}
	if err, list, count := userService.FindUserPoemByPage(r); err != nil {
		global.POEM_LOG.Error("参数错误", zap.Any("err", err))
		responce.FailWithMessage(err.Error(), c)
	} else {
		responce.OkWithDetailed(responce.PageResult{
			List:     list,
			Total:    count,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "获取成功", c)
	}
}
func (b BaseApi) PostUserPoem(c *gin.Context) {
	var r system.SysPoem
	if err := c.ShouldBindJSON(&r); err != nil {
		global.POEM_LOG.Error("参数错误", zap.Any("err", err))
		responce.FailWithMessage("输入参数错误", c)
		return
	}
	if err := userService.PostUserPoem(r); err != nil {
		global.POEM_LOG.Error("插入信息失败", zap.Any("err", err))
		responce.FailWithMessage(err.Error(), c)
	} else {
		responce.OkWithMessage("创建成功", c)
	}
}
