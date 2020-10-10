package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageWeightLog
// @description   create a SmartStorageWeightLog
// @param     smartStorageWeightLog               model.SmartStorageWeightLog
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageWeightLog(smartStorageWeightLog model.SmartStorageWeightLog) (err error) {
	err = global.GVA_DB.Create(&smartStorageWeightLog).Error
	return err
}

// @title    DeleteSmartStorageWeightLog
// @description   delete a SmartStorageWeightLog
// @auth                     （2020/04/05  20:22）
// @param     smartStorageWeightLog               model.SmartStorageWeightLog
// @return                    error

func DeleteSmartStorageWeightLog(smartStorageWeightLog model.SmartStorageWeightLog) (err error) {
	err = global.GVA_DB.Delete(smartStorageWeightLog).Error
	return err
}

// @title    DeleteSmartStorageWeightLogByIds
// @description   delete SmartStorageWeightLogs
// @auth                     （2020/04/05  20:22）
// @param     smartStorageWeightLog               model.SmartStorageWeightLog
// @return                    error

func DeleteSmartStorageWeightLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageWeightLog{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageWeightLog
// @description   update a SmartStorageWeightLog
// @param     smartStorageWeightLog          *model.SmartStorageWeightLog
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageWeightLog(smartStorageWeightLog *model.SmartStorageWeightLog) (err error) {
	err = global.GVA_DB.Save(smartStorageWeightLog).Error
	return err
}

// @title    GetSmartStorageWeightLog
// @description   get the info of a SmartStorageWeightLog
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageWeightLog        SmartStorageWeightLog

func GetSmartStorageWeightLog(id uint) (err error, smartStorageWeightLog model.SmartStorageWeightLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageWeightLog).Error
	return
}

// @title    GetSmartStorageWeightLogInfoList
// @description   get SmartStorageWeightLog list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageWeightLogInfoList(info request.SmartStorageWeightLogSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmartStorageWeightLog{})
    var smartStorageWeightLogs []model.SmartStorageWeightLog
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageWeightLogs).Error
	return err, smartStorageWeightLogs, total
}