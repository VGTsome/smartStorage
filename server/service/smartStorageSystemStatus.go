package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageSystemStatus
// @description   create a SmartStorageSystemStatus
// @param     ssss               model.SmartStorageSystemStatus
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageSystemStatus(ssss model.SmartStorageSystemStatus) (err error) {
	err = global.GVA_DB.Create(&ssss).Error
	return err
}

// @title    DeleteSmartStorageSystemStatus
// @description   delete a SmartStorageSystemStatus
// @auth                     （2020/04/05  20:22）
// @param     ssss               model.SmartStorageSystemStatus
// @return                    error

func DeleteSmartStorageSystemStatus(ssss model.SmartStorageSystemStatus) (err error) {
	err = global.GVA_DB.Delete(ssss).Error
	return err
}

// @title    DeleteSmartStorageSystemStatusByIds
// @description   delete SmartStorageSystemStatuss
// @auth                     （2020/04/05  20:22）
// @param     ssss               model.SmartStorageSystemStatus
// @return                    error

func DeleteSmartStorageSystemStatusByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageSystemStatus{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageSystemStatus
// @description   update a SmartStorageSystemStatus
// @param     ssss          *model.SmartStorageSystemStatus
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageSystemStatus(ssss *model.SmartStorageSystemStatus) (err error) {
	err = global.GVA_DB.Save(ssss).Error
	return err
}

// @title    GetSmartStorageSystemStatus
// @description   get the info of a SmartStorageSystemStatus
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageSystemStatus        SmartStorageSystemStatus

func GetSmartStorageSystemStatus(id uint) (err error, ssss model.SmartStorageSystemStatus) {
	err = global.GVA_DB.Where("id = ?", id).First(&ssss).Error
	return
}

//SetSystemStatus 1:正常 2:管理员理货 3:取货中 4:异常状态
func SetSystemComments(Comments string) {
	_, ssss := GetSmartStorageSystemStatus(1)
	ssss.Comments = Comments
	UpdateSmartStorageSystemStatus(&ssss)
}

//GetSystemStatus 1:正常 2:管理员理货 3:取货中 4:异常状态
func GetSystemComments() string {
	_, ssss := GetSmartStorageSystemStatus(1)
	return ssss.Comments
}

//SetSystemStatus 1:正常 2:管理员理货 3:取货中 4:异常状态
func SetSystemStatus(status int) {
	_, ssss := GetSmartStorageSystemStatus(1)
	ssss.SystemStatus = status
	UpdateSmartStorageSystemStatus(&ssss)
}

//GetSystemStatus 1:正常 2:管理员理货 3:取货中 4:异常状态
func GetSystemStatus() int {
	_, ssss := GetSmartStorageSystemStatus(1)
	return ssss.SystemStatus
}

// @title    GetSmartStorageSystemStatusInfoList
// @description   get SmartStorageSystemStatus list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageSystemStatusInfoList(info request.SmartStorageSystemStatusSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageSystemStatus{})
	var sssss []model.SmartStorageSystemStatus
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sssss).Error
	return err, sssss, total
}
