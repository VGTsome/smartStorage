// 自动生成模板CabinetProduct
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type CabinetProduct struct {
	gorm.Model
	CabinetId           int                 `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(11)"`
	SmartStorageCabinet SmartStorageCabinet `gorm:"ForeignKey:CabinetId;AssociationForeignKey:CabinetId"`
	ProductId           string              `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:varchar(255)"`
	SmartStorageProduct SmartStorageProduct `gorm:"ForeignKey:ProductId;AssociationForeignKey:ProductId"`
	Weight              int                 `json:"weight" form:"weight" gorm:"column:weight;comment:'';type:int(11)"`
	ProductNumber       int                 `json:"productNumber" form:"productNumber" gorm:"column:product_number;comment:'';type:int(11)"`
}

type Shelf struct {
	ShelfID string `json:"shelfID" form:"shelfID" "`
	Action  int    `json:"action" form:"action" "`
}

func (CabinetProduct) TableName() string {
	return "smart_storage_cabinet_product"
}
