package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSmartStorageQuota
// @description   create a SmartStorageQuota
// @param     ssq               model.SmartStorageQuota
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSmartStorageQuota(ssq model.SmartStorageQuota) (err error) {
	err = global.GVA_DB.Create(&ssq).Error
	return err
}

// @title    DeleteSmartStorageQuota
// @description   delete a SmartStorageQuota
// @auth                     （2020/04/05  20:22）
// @param     ssq               model.SmartStorageQuota
// @return                    error

func DeleteSmartStorageQuota(ssq model.SmartStorageQuota) (err error) {
	err = global.GVA_DB.Delete(ssq).Error
	return err
}

// @title    DeleteSmartStorageQuotaByIds
// @description   delete SmartStorageQuotas
// @auth                     （2020/04/05  20:22）
// @param     ssq               model.SmartStorageQuota
// @return                    error

func DeleteSmartStorageQuotaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmartStorageQuota{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateSmartStorageQuota
// @description   update a SmartStorageQuota
// @param     ssq          *model.SmartStorageQuota
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSmartStorageQuota(ssq *model.SmartStorageQuota) (err error) {
	err = global.GVA_DB.Save(ssq).Error
	return err
}

// @title    GetSmartStorageQuota
// @description   get the info of a SmartStorageQuota
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SmartStorageQuota        SmartStorageQuota

func GetSmartStorageQuota(id uint) (err error, ssq model.SmartStorageQuota) {
	err = global.GVA_DB.Where("id = ?", id).First(&ssq).Error
	return
}

// @title    GetSmartStorageQuotaInfoList
// @description   get SmartStorageQuota list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSmartStorageQuotaInfoList(info request.SmartStorageQuotaSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmartStorageQuota{})
	var ssqs []model.SmartStorageQuota
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ssqs).Error
	for index, _ := range ssqs {
		global.GVA_DB.Model(&ssqs[index]).Select("product_name").Related(&ssqs[index].SmartStorageProduct, "SmartStorageProduct")
		global.GVA_DB.Model(&ssqs[index]).Select("authority_name").Related(&ssqs[index].SysAuthority, "SysAuthority")
	}
	return err, ssqs, total
}
