package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStoragePassWeight
// @description   create a SmartStoragePassWeight
// @param     smartStoragePassWeight               model.SmartStoragePassWeight
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStoragePassWeight(smartStoragePassWeight model.SmartStoragePassWeight) (err error) {
	err = global.GVA_DB.Create(&smartStoragePassWeight).Error
	return err
}

// @title    DeleteSmartStoragePassWeight
// @description   delete a SmartStoragePassWeight
// @auth                     （2020/04/05  20:22）
// @param     smartStoragePassWeight               model.SmartStoragePassWeight
// @return                    error

func DeleteSmartStoragePassWeight(smartStoragePassWeight model.SmartStoragePassWeight) (err error) {
	err = global.GVA_DB.Unscoped().Delete(smartStoragePassWeight).Error
	return err
}

// @title    DeleteSmartStoragePassWeightByIds
// @description   delete SmartStoragePassWeights
// @auth                     （2020/04/05  20:22）
// @param     smartStoragePassWeight               model.SmartStoragePassWeight
// @return                    error

func DeleteSmartStoragePassWeightByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]model.SmartStoragePassWeight{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStoragePassWeight
// @description   update a SmartStoragePassWeight
// @param     smartStoragePassWeight          *model.SmartStoragePassWeight
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStoragePassWeight(smartStoragePassWeight *model.SmartStoragePassWeight) (err error) {
	err = global.GVA_DB.Save(smartStoragePassWeight).Error
	return err
}

// @title    GetSmartStoragePassWeight
// @description   get the info of a SmartStoragePassWeight
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStoragePassWeight        SmartStoragePassWeight

func GetSmartStoragePassWeight(id uint) (err error, smartStoragePassWeight model.SmartStoragePassWeight) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStoragePassWeight).Error
	return
}
func GetSmartStoragePassWeightFirst() (err error, smartStoragePassWeight model.SmartStoragePassWeight) {
	err = global.GVA_DB.First(&smartStoragePassWeight).Error
	return
}

// @title    GetSmartStoragePassWeightInfoList
// @description   get SmartStoragePassWeight list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStoragePassWeightInfoList(info request.SmartStoragePassWeightSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStoragePassWeight{})
	var smartStoragePassWeights []model.SmartStoragePassWeight
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStoragePassWeights).Error
	return err, smartStoragePassWeights, total
}

//GetSmartStoragePassWeightListByShelf 获取对应shelf的passlist
func GetSmartStoragePassWeightListByShelf(com string, shelf string) (err error, smartStoragePassWeights []model.SmartStoragePassWeight, total int) {

	// 创建db
	db := global.GVA_DB.Model(&model.SmartStoragePassWeight{})

	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Joins("LEFT JOIN smart_storage_cabinet ON smart_storage_cabinet.cabinet_id=smart_storage_pass_weight.cabinet_id").Where("smart_storage_cabinet.cabinet_name LIKE ?", com+"-"+shelf+"-%").Find(&smartStoragePassWeights).Error
	return err, smartStoragePassWeights, total
}

//GetSmartStoragePassWeightStatusList 获取对应orderID的pass为0列表
func GetSmartStoragePassWeightStatusList(pass int) (err error, smartStoragePassWeights []model.SmartStoragePassWeight, total int) {

	// 创建db
	db := global.GVA_DB.Model(&model.SmartStoragePassWeight{})

	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Where("pass = ?  ", pass).Find(&smartStoragePassWeights).Error
	return err, smartStoragePassWeights, total
}

//TruncateSmartStoragePassWeight 清空表
func TruncateSmartStoragePassWeight() {
	db := global.GVA_DB.Model(&model.SmartStoragePassWeight{})
	db.Exec("TRUNCATE TABLE smart_storage_pass_weight")
}
