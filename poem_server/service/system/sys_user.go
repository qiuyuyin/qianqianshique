package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"poem_server/global"
	"poem_server/model/system"
	"poem_server/model/system/sys_request"
	"poem_server/utils"
)

type UserService struct {
}

// 用户注册
func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.POEM_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.UUID = uuid.NewV4()
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.POEM_DB.Create(&u).Error
	return err, u
}

// 用户登录
func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	if global.POEM_DB.Where("username = ?", u.Username).Preload("Authority").First(&user).Error != nil {
		return errors.New("用户名不存在"), &user
	}
	if user.Password != u.Password {
		return errors.New("密码不正确"), &user
	}
	return nil, &user
}

// 删除一个用户
func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.SysUser
	err = global.POEM_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return err
}

// 设置用户的信息ID
func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.POEM_DB.Updates(&reqUser).Error
	return err, reqUser
}

//查找用户的信息UUID
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.POEM_DB.Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

//根据Id查找user
func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.POEM_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// 根据UUID查找user
func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.POEM_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

// 根据UUID查找user
func (userService *UserService) FindUserPoemByPage(info sys_request.SearchUserPoem) (err error, list interface{}, count int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.POEM_DB.Model(&system.SysPoem{})
	var poemList []system.SysPoem
	err = db.Count(&count).Error
	if err != nil {
		return err, poemList, count
	}
	db = db.Limit(limit).Offset(offset)
	err = db.Order("created_at desc").Find(&poemList).Error
	if err != nil {
		return err, poemList, count
	}
	return nil, poemList, count
}

// 根据UUID查找user
func (userService *UserService) PostUserPoem(info system.SysPoem) (err error) {
	return global.POEM_DB.Create(&info).Error
}
