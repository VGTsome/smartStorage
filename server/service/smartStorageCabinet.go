package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageCabinet
// @description   create a SmartStorageCabinet
// @param     smartStorageCabinet               model.SmartStorageCabinet
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageCabinet(smartStorageCabinet model.SmartStorageCabinet) (err error) {
	var stemp model.SmartStorageCabinet
	db := global.GVA_DB.Model(&model.SmartStorageCabinet{})
	db.Unscoped().Select("cabinet_id").Order("cabinet_id desc").First(&stemp)
	smartStorageCabinet.CabinetId = stemp.CabinetId + 1
	err = global.GVA_DB.Create(&smartStorageCabinet).Error
	return err
}

// @title    DeleteSmartStorageCabinet
// @description   delete a SmartStorageCabinet
// @auth                     （2020/04/05  20:22）
// @param     smartStorageCabinet               model.SmartStorageCabinet
// @return                    error

func DeleteSmartStorageCabinet(smartStorageCabinet model.SmartStorageCabinet) (err error) {
	err = global.GVA_DB.Delete(smartStorageCabinet).Error
	return err
}

// @title    DeleteSmartStorageCabinetByIds
// @description   delete SmartStorageCabinets
// @auth                     （2020/04/05  20:22）
// @param     smartStorageCabinet               model.SmartStorageCabinet
// @return                    error

func DeleteSmartStorageCabinetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageCabinet{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageCabinet
// @description   update a SmartStorageCabinet
// @param     smartStorageCabinet          *model.SmartStorageCabinet
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageCabinet(smartStorageCabinet *model.SmartStorageCabinet) (err error) {
	err = global.GVA_DB.Save(smartStorageCabinet).Error
	return err
}

// @title    GetSmartStorageCabinet
// @description   get the info of a SmartStorageCabinet
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageCabinet        SmartStorageCabinet

func GetSmartStorageCabinet(id uint) (err error, smartStorageCabinet model.SmartStorageCabinet) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageCabinet).Error
	return
}

func GetSmartStorageCabinetByCabinetID(cabinetID int) (err error, smartStorageCabinet model.SmartStorageCabinet) {
	err = global.GVA_DB.Where("cabinet_id = ?", cabinetID).First(&smartStorageCabinet).Error
	return
}

// @title    GetSmartStorageCabinetInfoList
// @description   get SmartStorageCabinet list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageCabinetInfoList(info request.SmartStorageCabinetSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageCabinet{})
	var smartStorageCabinets []model.SmartStorageCabinet
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageCabinets).Error
	return err, smartStorageCabinets, total
}

func GetSmartStorageCabinetList() (err error, list []model.SmartStorageCabinet, total int) {

	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageCabinet{})
	var smartStorageCabinets []model.SmartStorageCabinet
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Find(&smartStorageCabinets).Error
	return err, smartStorageCabinets, total
}
