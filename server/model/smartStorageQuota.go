// 自动生成模板SmartStorageQuota
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageQuota struct {
	gorm.Model
	AuthorityId         string              `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:'';type:varchar(255)"`
	SysAuthority        SysAuthority        `gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	ProductId           string              `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:varchar(255)"`
	QuotaMonth          int                 `json:"quotaMonth" form:"quotaMonth" gorm:"column:quota_month;comment:'';type:int(11)"`
	SmartStorageProduct SmartStorageProduct `gorm:"ForeignKey:ProductId;AssociationForeignKey:ProductId"`
	UsedMonth           int                 `json:"usedMonth" form:"usedMonth" gorm:"column:used_month;comment:'';type:int(11)"`
	QuotaYear           int                 `json:"quotaYear" form:"quotaYear" gorm:"column:quota_year;comment:'';type:int(11)"`
	UsedYear            int                 `json:"usedYear" form:"usedYear" gorm:"column:used_year;comment:'';type:int(11)"`
}

func (SmartStorageQuota) TableName() string {
	return "smart_storage_quota"
}
