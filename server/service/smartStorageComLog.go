package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageComLog
// @description   create a SmartStorageComLog
// @param     smartStorageComLog               model.SmartStorageComLog
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageComLog(smartStorageComLog model.SmartStorageComLog) (err error) {
	err = global.GVA_DB.Create(&smartStorageComLog).Error
	return err
}

// @title    DeleteSmartStorageComLog
// @description   delete a SmartStorageComLog
// @auth                     （2020/04/05  20:22）
// @param     smartStorageComLog               model.SmartStorageComLog
// @return                    error

func DeleteSmartStorageComLog(smartStorageComLog model.SmartStorageComLog) (err error) {
	err = global.GVA_DB.Delete(smartStorageComLog).Error
	return err
}

// @title    DeleteSmartStorageComLogByIds
// @description   delete SmartStorageComLogs
// @auth                     （2020/04/05  20:22）
// @param     smartStorageComLog               model.SmartStorageComLog
// @return                    error

func DeleteSmartStorageComLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageComLog{}, "id in (?)", ids.Ids).Error
	return err
}
func DeleteSmartStorageComLogEmpty() (err error) {
	err = global.GVA_DB.Where("receiveCmd = '' ").Delete(&model.SmartStorageComLog{}).Error
	return err
}

// @title    UpdateSmartStorageComLog
// @description   update a SmartStorageComLog
// @param     smartStorageComLog          *model.SmartStorageComLog
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageComLog(smartStorageComLog *model.SmartStorageComLog) (err error) {
	err = global.GVA_DB.Save(smartStorageComLog).Error
	return err
}

// @title    GetSmartStorageComLog
// @description   get the info of a SmartStorageComLog
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageComLog        SmartStorageComLog

func GetSmartStorageComLog(id uint) (err error, smartStorageComLog model.SmartStorageComLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageComLog).Error
	return
}
func GetSmartStorageComLogByTatoo(tatoo string, funcname string) (err error, smartStorageComLog model.SmartStorageComLog) {
	err = global.GVA_DB.Where("tatoo = ? AND func=?", tatoo, funcname).Order("updated_at DESC").First(&smartStorageComLog).Error
	return
}
func GetSmartStorageComLogNotRecv() (bool, string) {
	var smartStorageComLog model.SmartStorageComLog

	global.GVA_DB.Where("receiveCmd = ''").First(&smartStorageComLog)

	if smartStorageComLog.ID == 0 {

		return true, ""
	} else {

		return false, smartStorageComLog.SendCmd
	}
}

// @title    GetSmartStorageComLogInfoList
// @description   get SmartStorageComLog list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageComLogInfoList(info request.SmartStorageComLogSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageComLog{})
	var smartStorageComLogs []model.SmartStorageComLog
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageComLogs).Error
	return err, smartStorageComLogs, total
}
