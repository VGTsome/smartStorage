package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageCurrentOrder
// @description   create a SmartStorageCurrentOrder
// @param     smartStorageCurrentOrder               model.SmartStorageCurrentOrder
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageCurrentOrder(smartStorageCurrentOrder model.SmartStorageCurrentOrder) (err error) {
	err = global.GVA_DB.Create(&smartStorageCurrentOrder).Error
	return err
}

// @title    DeleteSmartStorageCurrentOrder
// @description   delete a SmartStorageCurrentOrder
// @auth                     （2020/04/05  20:22）
// @param     smartStorageCurrentOrder               model.SmartStorageCurrentOrder
// @return                    error

func DeleteSmartStorageCurrentOrder(smartStorageCurrentOrder model.SmartStorageCurrentOrder) (err error) {
	err = global.GVA_DB.Unscoped().Delete(smartStorageCurrentOrder).Error
	return err
}

// @title    DeleteSmartStorageCurrentOrderByIds
// @description   delete SmartStorageCurrentOrders
// @auth                     （2020/04/05  20:22）
// @param     smartStorageCurrentOrder               model.SmartStorageCurrentOrder
// @return                    error

func DeleteSmartStorageCurrentOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]model.SmartStorageCurrentOrder{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageCurrentOrder
// @description   update a SmartStorageCurrentOrder
// @param     smartStorageCurrentOrder          *model.SmartStorageCurrentOrder
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageCurrentOrder(smartStorageCurrentOrder *model.SmartStorageCurrentOrder) (err error) {
	err = global.GVA_DB.Save(smartStorageCurrentOrder).Error
	return err
}

// @title    GetSmartStorageCurrentOrder
// @description   get the info of a SmartStorageCurrentOrder
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageCurrentOrder        SmartStorageCurrentOrder

func GetSmartStorageCurrentOrder(id uint) (err error, smartStorageCurrentOrder model.SmartStorageCurrentOrder) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageCurrentOrder).Error
	return
}

// @title    GetSmartStorageCurrentOrderInfoList
// @description   get SmartStorageCurrentOrder list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageCurrentOrderInfoList(info request.SmartStorageCurrentOrderSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageCurrentOrder{})
	var smartStorageCurrentOrders []model.SmartStorageCurrentOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageCurrentOrders).Error
	return err, smartStorageCurrentOrders, total
}

func GetSmartStorageCurrentOrderInfoAllList() (err error, smartStorageCurrentOrders []model.SmartStorageCurrentOrder, total int) {
	limit := 1000
	offset := 0

	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageCurrentOrder{})
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageCurrentOrders).Error
	return err, smartStorageCurrentOrders, total
}

//TruncateSmartStoragePassWeight 清空表
func TruncateSmartStorageCurrentOrder() {
	db := global.GVA_DB.Model(&model.SmartStorageCurrentOrder{})
	db.Exec("TRUNCATE TABLE smart_storage_current_order")
}
