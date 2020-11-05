package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCabinetProduct
// @description   create a CabinetProduct
// @param     sscp               model.CabinetProduct
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCabinetProduct(sscp model.CabinetProduct) (err error) {
	err = global.GVA_DB.Create(&sscp).Error
	return err
}

// @title    DeleteCabinetProduct
// @description   delete a CabinetProduct
// @auth                     （2020/04/05  20:22）
// @param     sscp               model.CabinetProduct
// @return                    error

func DeleteCabinetProduct(sscp model.CabinetProduct) (err error) {
	err = global.GVA_DB.Delete(sscp).Error
	return err
}

// @title    DeleteCabinetProductByIds
// @description   delete CabinetProducts
// @auth                     （2020/04/05  20:22）
// @param     sscp               model.CabinetProduct
// @return                    error

func DeleteCabinetProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CabinetProduct{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateCabinetProduct
// @description   update a CabinetProduct
// @param     sscp          *model.CabinetProduct
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCabinetProduct(sscp *model.CabinetProduct) (err error) {
	err = global.GVA_DB.Save(sscp).Error
	return err
}

// @title    GetCabinetProduct
// @description   get the info of a CabinetProduct
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    CabinetProduct        CabinetProduct

func GetCabinetProduct(id uint) (err error, sscp model.CabinetProduct) {
	err = global.GVA_DB.Where("id = ?", id).First(&sscp).Error
	return
}
func GetCabinetProductByCabinetName(cabinetName string) (err error, sscp model.CabinetProduct) {
	var ssc model.SmartStorageCabinet
	err = global.GVA_DB.Where("cabinet_name = ?", cabinetName).First(&ssc).Error
	err = global.GVA_DB.Where("cabinet_id = ?", ssc.CabinetId).First(&sscp).Error
	return
}
func GetCabinetProductByProductId(productId string) (err error, sscp model.CabinetProduct) {

	err = global.GVA_DB.Where("product_id = ?", productId).First(&sscp).Error
	return
}

// @title    GetCabinetProductInfoList
// @description   get CabinetProduct list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCabinetProductList(info request.CabinetProductSearch) (err error, sscps []model.CabinetProduct, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CabinetProduct{})

	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sscps).Error
	for index, _ := range sscps {
		global.GVA_DB.Model(&sscps[index]).Related(&sscps[index].SmartStorageProduct, "SmartStorageProduct")
		global.GVA_DB.Model(&sscps[index]).Related(&sscps[index].SmartStorageCabinet, "SmartStorageCabinet")
	}
	return err, sscps, total
}
func GetAllCabinetProductList() (err error, sscps []model.CabinetProduct, total int) {
	db := global.GVA_DB.Model(&model.CabinetProduct{})

	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Find(&sscps).Error
	return err, sscps, total
}
func GetCabinetProductInfoList(info request.CabinetProductSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CabinetProduct{})
	var sscps []model.CabinetProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sscps).Error
	for index, _ := range sscps {
		global.GVA_DB.Model(&sscps[index]).Related(&sscps[index].SmartStorageProduct, "SmartStorageProduct")
		global.GVA_DB.Model(&sscps[index]).Related(&sscps[index].SmartStorageCabinet, "SmartStorageCabinet")
	}
	return err, sscps, total
}
