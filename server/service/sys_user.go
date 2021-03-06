package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"

	uuid "github.com/satori/go.uuid"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notRegister := global.GVA_DB.Where("username = ? or scan_id= ? ", u.Username, u.ScanID).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名或门禁ID已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}
func UpdateUser(u request.RegisterStruct) (err error, userInter model.SysUser) {
	var oriuser model.SysUser
	var checkuser model.SysUser
	var checkScanID model.SysUser
	global.GVA_DB.Where("id = ?", u.ID).First(&oriuser)
	existName := global.GVA_DB.Where("username = ? ", u.Username).First(&checkuser).RecordNotFound()
	existScanID := global.GVA_DB.Where("scan_id = ? ", u.ScanID).First(&checkScanID).RecordNotFound()

	// notRegister为false表明读取到了 不能注册

	if !existName && oriuser.Username != u.Username {
		return errors.New("用户名重复"), userInter
	} else if !existScanID && oriuser.ScanID != u.ScanID {
		return errors.New("门禁ID重复"), userInter
	} else {
		if u.Password != "" {
			u.Password = utils.MD5V([]byte(u.Password))
		}

		global.GVA_DB.Model(model.SysUser{}).Where("id = ?", u.ID).Updates(&u)
	}
	return err, userInter
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser

func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

//GetUserByScanID 根据scanID获取用户信息
func GetUserByScanID(scanID string) (err error, user model.SysUser) {

	err = global.GVA_DB.Where("scan_id = ?", scanID).First(&user).Error
	return err, user
}

func GetUserByID(ID uint) (err error, user model.SysUser) {

	err = global.GVA_DB.Where("id = ?", ID).First(&user).Error
	return err, user
}
