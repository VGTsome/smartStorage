package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageProduct
// @description   create a SmartStorageProduct
// @param     smartStorageProduct               model.SmartStorageProduct
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageProduct(smartStorageProduct model.SmartStorageProduct) (err error) {
	err = global.GVA_DB.Create(&smartStorageProduct).Error
	return err
}

// @title    DeleteSmartStorageProduct
// @description   delete a SmartStorageProduct
// @auth                     （2020/04/05  20:22）
// @param     smartStorageProduct               model.SmartStorageProduct
// @return                    error

func DeleteSmartStorageProduct(smartStorageProduct model.SmartStorageProduct) (err error) {
	err = global.GVA_DB.Delete(smartStorageProduct).Error
	return err
}

// @title    DeleteSmartStorageProductByIds
// @description   delete SmartStorageProducts
// @auth                     （2020/04/05  20:22）
// @param     smartStorageProduct               model.SmartStorageProduct
// @return                    error

func DeleteSmartStorageProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageProduct{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageProduct
// @description   update a SmartStorageProduct
// @param     smartStorageProduct          *model.SmartStorageProduct
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageProduct(smartStorageProduct *model.SmartStorageProduct) (err error) {
	err = global.GVA_DB.Save(smartStorageProduct).Error
	return err
}

// @title    GetSmartStorageProduct
// @description   get the info of a SmartStorageProduct
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageProduct        SmartStorageProduct

func GetSmartStorageProduct(id uint) (err error, smartStorageProduct model.SmartStorageProduct) {
	err = global.GVA_DB.Where("id = ?", id).First(&smartStorageProduct).Error
	return
}

// @title    GetSmartStorageProductInfoList
// @description   get SmartStorageProduct list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageProductInfoList(info request.SmartStorageProductSearch) (err error, list interface{}, total int) {
	productName := info.ProductName
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageProduct{})
	var smartStorageProducts []model.SmartStorageProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if productName != "" {
		db = db.Where("product_name LIKE ?", "%"+productName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&smartStorageProducts).Error
	return err, smartStorageProducts, total
}
