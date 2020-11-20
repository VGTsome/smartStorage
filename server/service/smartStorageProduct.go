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
	err = global.GVA_DB.Unscoped().Delete(smartStorageProduct).Error
	return err
}

// @title    DeleteSmartStorageProductByIds
// @description   delete SmartStorageProducts
// @auth                     （2020/04/05  20:22）
// @param     smartStorageProduct               model.SmartStorageProduct
// @return                    error

func DeleteSmartStorageProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]model.SmartStorageProduct{}, "id in (?)", ids.Ids).Error
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
func GetSmartStorageProductByProductId(productId string) (err error, smartStorageProduct model.SmartStorageProduct) {
	err = global.GVA_DB.Where("productId = ?", productId).First(&smartStorageProduct).Error
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
	var smartStorageProducts []model.SmartStorageProduct
	// var cabinetProducts []model.CabinetProduct
	// 创建db
	//db := global.GVA_DB.Model(&model.CabinetProduct{}).Find(&cabinetProducts)

	db := global.GVA_DB.Model(&model.SmartStorageProduct{})

	// 如果有条件搜索 下方会自动创建搜索语句
	if productName != "" {
		db = db.Where("product_name LIKE ?", "%"+productName+"%")
	}
	err = db.Count(&total).Error

	err = db.Limit(limit).Offset(offset).Find(&smartStorageProducts).Error

	return err, smartStorageProducts, total
}

func GetSmartStorageProductInfoValidList(info request.SmartStorageProductSearch) (err error, list interface{}, total int) {
	productName := info.ProductName
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var smartStorageProducts []model.SmartStorageProduct
	whereStr := ""
	if productName != "" {
		whereStr = "WHERE p.product_name='" + productName + "'"
	}
	// 创建db
	err = global.GVA_DB.Raw("SELECT p.* FROM `smart_storage_product` p  JOIN smart_storage_cabinet_product cp on p.product_id=cp.product_id  "+whereStr+" LIMIT ?,?", offset, limit).Count(&total).Error
	err = global.GVA_DB.Raw("SELECT p.* FROM `smart_storage_product` p  JOIN smart_storage_cabinet_product cp on p.product_id=cp.product_id  "+whereStr+" LIMIT ?,?", offset, limit).Scan(&smartStorageProducts).Error

	return err, smartStorageProducts, total
}
