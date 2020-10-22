package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageComReceive
// @description   create a SmartStorageComReceive
// @param     smartStorageComReceive               model.SmartStorageComReceive
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageComReceive(smartStorageComReceive model.SmartStorageComReceive) (err error) {
	err = global.GVA_DB.Create(&smartStorageComReceive).Error
	return err
}

// @title    DeleteSmartStorageComReceive
// @description   delete a SmartStorageComReceive
// @auth                     （2020/04/05  20:22）
// @param     smartStorageComReceive               model.SmartStorageComReceive
// @return                    error

func DeleteSmartStorageComReceive(smartStorageComReceive model.SmartStorageComReceive) (err error) {
	err = global.GVA_DB.Delete(smartStorageComReceive).Error
	return err
}

// @title    DeleteSmartStorageComReceiveByIds
// @description   delete SmartStorageComReceives
// @auth                     （2020/04/05  20:22）
// @param     smartStorageComReceive               model.SmartStorageComReceive
// @return                    error

func DeleteSmartStorageComReceiveByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageComReceive{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageComReceive
// @description   update a SmartStorageComReceive
// @param     smartStorageComReceive          *model.SmartStorageComReceive
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageComReceive(smartStorageComReceive *model.SmartStorageComReceive) (err error) {
	err = global.GVA_DB.Save(smartStorageComReceive).Error
	return err
}

// @title    GetSmartStorageComReceive
// @description   get the info of a SmartStorageComReceive
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageComReceive        SmartStorageComReceive

func GetSmartStorageComReceive(id uint) (err error, smartStorageComReceive model.SmartStorageComReceive) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageComReceive).Error
	return
}

// @title    GetSmartStorageComReceiveInfoList
// @description   get SmartStorageComReceive list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageComReceiveInfoList(info request.SmartStorageComReceiveSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmartStorageComReceive{})
    var smartStorageComReceives []model.SmartStorageComReceive
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageComReceives).Error
	return err, smartStorageComReceives, total
}