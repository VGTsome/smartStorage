package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
		model.CabinetProduct{},
		model.SmartStorageCabinet{},
		model.SmartStorageDoor{},
		model.SmartStorageOrder{},
		model.SmartStorageMonitor{},
		model.SmartStorageProduct{},
		model.SmartStorageQuota{},
		model.SmartStoragePassWeight{},
		model.SmartStorageCurrentOrder{},
		model.SmartStorageSystemStatus{},
		model.SmartStorageComLog{},
		model.HkScan{},
	)
	global.GVA_LOG.Debug("register table success")
}
