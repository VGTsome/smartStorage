package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageOrder
// @description   create a SmartStorageOrder
// @param     smartStorageOrder               model.SmartStorageOrder
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageOrder(smartStorageOrder model.SmartStorageOrder) (err error) {
	err = global.GVA_DB.Create(&smartStorageOrder).Error
	return err
}

// @title    DeleteSmartStorageOrder
// @description   delete a SmartStorageOrder
// @auth                     （2020/04/05  20:22）
// @param     smartStorageOrder               model.SmartStorageOrder
// @return                    error

func DeleteSmartStorageOrder(smartStorageOrder model.SmartStorageOrder) (err error) {
	err = global.GVA_DB.Delete(smartStorageOrder).Error
	return err
}

// @title    DeleteSmartStorageOrderByIds
// @description   delete SmartStorageOrders
// @auth                     （2020/04/05  20:22）
// @param     smartStorageOrder               model.SmartStorageOrder
// @return                    error

func DeleteSmartStorageOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageOrder{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageOrder
// @description   update a SmartStorageOrder
// @param     smartStorageOrder          *model.SmartStorageOrder
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageOrder(smartStorageOrder *model.SmartStorageOrder) (err error) {
	err = global.GVA_DB.Save(smartStorageOrder).Error
	return err
}

// @title    GetSmartStorageOrder
// @description   get the info of a SmartStorageOrder
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageOrder        SmartStorageOrder

func GetSmartStorageOrder(id uint) (err error, smartStorageOrder model.SmartStorageOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageOrder).Error
	return
}

// @title    GetSmartStorageOrderInfoList
// @description   get SmartStorageOrder list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageOrderInfoList(info request.SmartStorageOrderSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmartStorageOrder{})
    var smartStorageOrders []model.SmartStorageOrder
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageOrders).Error
	return err, smartStorageOrders, total
}