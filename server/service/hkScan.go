package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHkScan
// @description   create a HkScan
// @param     hkScan               model.HkScan
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHkScan(hkScan model.HkScan) (err error) {
	err = global.GVA_DB.Create(&hkScan).Error
	return err
}

// @title    DeleteHkScan
// @description   delete a HkScan
// @auth                     （2020/04/05  20:22）
// @param     hkScan               model.HkScan
// @return                    error

func DeleteHkScan(hkScan model.HkScan) (err error) {
	err = global.GVA_DB.Delete(hkScan).Error
	return err
}

// @title    DeleteHkScanByIds
// @description   delete HkScans
// @auth                     （2020/04/05  20:22）
// @param     hkScan               model.HkScan
// @return                    error

func DeleteHkScanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HkScan{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateHkScan
// @description   update a HkScan
// @param     hkScan          *model.HkScan
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHkScan(hkScan *model.HkScan) (err error) {
	err = global.GVA_DB.Save(hkScan).Error
	return err
}

// @title    GetHkScan
// @description   get the info of a HkScan
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    HkScan        HkScan

func GetHkScan(id uint) (err error, hkScan model.HkScan) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkScan).Error
	return
}

// @title    GetHkScanInfoList
// @description   get HkScan list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHkScanInfoList(info request.HkScanSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.HkScan{})
    var hkScans []model.HkScan
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hkScans).Error
	return err, hkScans, total
}