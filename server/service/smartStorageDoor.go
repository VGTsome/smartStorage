package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageDoor
// @description   create a SmartStorageDoor
// @param     smartStorageDoor               model.SmartStorageDoor
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageDoor(smartStorageDoor model.SmartStorageDoor) (err error) {
	err = global.GVA_DB.Create(&smartStorageDoor).Error
	return err
}

// @title    DeleteSmartStorageDoor
// @description   delete a SmartStorageDoor
// @auth                     （2020/04/05  20:22）
// @param     smartStorageDoor               model.SmartStorageDoor
// @return                    error

func DeleteSmartStorageDoor(smartStorageDoor model.SmartStorageDoor) (err error) {
	err = global.GVA_DB.Delete(smartStorageDoor).Error
	return err
}

// @title    DeleteSmartStorageDoorByIds
// @description   delete SmartStorageDoors
// @auth                     （2020/04/05  20:22）
// @param     smartStorageDoor               model.SmartStorageDoor
// @return                    error

func DeleteSmartStorageDoorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageDoor{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageDoor
// @description   update a SmartStorageDoor
// @param     smartStorageDoor          *model.SmartStorageDoor
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageDoor(smartStorageDoor *model.SmartStorageDoor) (err error) {
	err = global.GVA_DB.Save(smartStorageDoor).Error
	return err
}

// @title    GetSmartStorageDoor
// @description   get the info of a SmartStorageDoor
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageDoor        SmartStorageDoor

func GetSmartStorageDoor(id uint) (err error, smartStorageDoor model.SmartStorageDoor) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageDoor).Error
	return
}

// @title    GetSmartStorageDoorInfoList
// @description   get SmartStorageDoor list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageDoorInfoList(info request.SmartStorageDoorSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmartStorageDoor{})
    var smartStorageDoors []model.SmartStorageDoor
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageDoors).Error
	return err, smartStorageDoors, total
}