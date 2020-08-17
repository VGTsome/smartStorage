// 自动生成模板SmartStorageProduct
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageProduct struct {
	gorm.Model
	ProductId          string `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:varchar(255)"`
	ProductName        string `json:"productName" form:"productName" gorm:"column:product_name;comment:'';type:varchar(1000)"`
	ProductWeight      int    `json:"productWeight" form:"productWeight" gorm:"column:product_weight;comment:'';type:int(11)"`
	ProductDescription string `json:"productDescription" form:"productDescription" gorm:"column:product_description;comment:'';type:varchar(1000)"`
	ProductImgUrl      string `json:"productImgUrl" form:"productImgUrl" gorm:"column:product_img_url;comment:'';type:varchar(1000)"`
	ProductNumber      int    `json:"productNumber" form:"productNumber" gorm:"column:product_number;comment:'';type:int(11)"`
}

func (SmartStorageProduct) TableName() string {
	return "smart_storage_product"
}
