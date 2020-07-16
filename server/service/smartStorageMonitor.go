package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageMonitor
// @description   create a SmartStorageMonitor
// @param     smartStorageMonitor               model.SmartStorageMonitor
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageMonitor(smartStorageMonitor model.SmartStorageMonitor) (err error) {
	err = global.GVA_DB.Create(&smartStorageMonitor).Error
	return err
}

// @title    DeleteSmartStorageMonitor
// @description   delete a SmartStorageMonitor
// @auth                     （2020/04/05  20:22）
// @param     smartStorageMonitor               model.SmartStorageMonitor
// @return                    error

func DeleteSmartStorageMonitor(smartStorageMonitor model.SmartStorageMonitor) (err error) {
	err = global.GVA_DB.Delete(smartStorageMonitor).Error
	return err
}

// @title    DeleteSmartStorageMonitorByIds
// @description   delete SmartStorageMonitors
// @auth                     （2020/04/05  20:22）
// @param     smartStorageMonitor               model.SmartStorageMonitor
// @return                    error

func DeleteSmartStorageMonitorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageMonitor{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageMonitor
// @description   update a SmartStorageMonitor
// @param     smartStorageMonitor          *model.SmartStorageMonitor
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageMonitor(smartStorageMonitor *model.SmartStorageMonitor) (err error) {
	err = global.GVA_DB.Save(smartStorageMonitor).Error
	return err
}

// @title    GetSmartStorageMonitor
// @description   get the info of a SmartStorageMonitor
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageMonitor        SmartStorageMonitor

func GetSmartStorageMonitor(id uint) (err error, smartStorageMonitor model.SmartStorageMonitor) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageMonitor).Error
	return
}

// @title    GetSmartStorageMonitorInfoList
// @description   get SmartStorageMonitor list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageMonitorInfoList(info request.SmartStorageMonitorSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmartStorageMonitor{})
    var smartStorageMonitors []model.SmartStorageMonitor
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageMonitors).Error
	return err, smartStorageMonitors, total
}