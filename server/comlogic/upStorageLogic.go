package comlogic

import (
	"gin-vue-admin/service"
	"strconv"
	"strings"
)

//upSetZero 数据库称台清零
func upSetZero(com string, command []string) {
	shelf, box := getCabinetAttr(command)
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	sscp.ProductNumber = 0
	sscp.Weight = 0
	service.UpdateCabinetProduct(&sscp)
}

//初始化第一批物品
func upSetFirstProd(com string, command []string) {
	shelf, box := getCabinetAttr(command)
	prodNumString := command[9] + command[10]
	prodNum := hexstringToNumber(prodNumString)
	weightString := command[4] + command[5] + command[6] + command[7] + command[8]
	weight := hexstringToNumber(weightString)

	singleWeight := weight / prodNum
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	sscp.ProductNumber = prodNum
	sscp.Weight = weight
	service.UpdateCabinetProduct(&sscp)
	_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
	ssp.ProductWeight = 0
	service.UpdateSmartStorageProduct(&ssp)
	_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)

	downSetProductInfo(ssc.CabinetName, singleWeight)

}
func getCabinetAttr(command []string) (shelf string, box string) {
	shelf = cabinetMinus30(command[3])
	box = cabinetMinus30(command[4])
	return shelf, box
}

//设置货物参数
func upSetSingWeight(com string, command []string) {
	shelf, box := getCabinetAttr(command)
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
	ssp.ProductWeight = hexstringToNumber(command[4] + command[5] + command[6])
	ssp.PackageWeight = hexstringToNumber(command[7] + command[8] + command[9])
	//ssp.ProductNumber = hexstringToNumber(command[10] + command[11])
	service.UpdateSmartStorageProduct(&ssp)

}

//更新整个柜子库存容量
//更新货物库存 保留位01
//进门前盘货 保留位02
//出门前盘货 保留位03
func upUpdateCabinetProduct(com string, command []string) {
	shelf, _ := getCabinetAttr(command)
	_, ssclist, _ := service.GetSmartStorageCabinetListByShelf(com, shelf)
	for _, ssc := range ssclist {
		_, sscp := service.GetCabinetProductByCabinetName(ssc.CabinetName)
		boxNumStr := strings.Split(ssc.CabinetName, "-")[1]
		sscp.ProductNumber = getBoxNumFromCommand(boxNumStr, command)
		_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
		sscp.Weight = ssp.ProductWeight * sscp.ProductNumber
		service.UpdateCabinetProduct(&sscp)
		//ssp.ProductNumber = sscp.ProductNumber
		//service.UpdateSmartStorageProduct(&ssp)
	}
}

func getBoxNumFromCommand(boxNumStr string, command []string) (proNum int) {
	boxNum, _ := strconv.Atoi(boxNumStr)
	proNumHex := command[boxNum*2-1+3] + command[boxNum*2-1+4]
	proNum = hexstringToNumber(proNumHex)
	return proNum

}

//进门前做货物检查记录
func beforeEnterDoor() {}

//开门指令
func userEnterDoor() {}

//出门前做货物检查记录
func beforeExitDoor() {}

//出门指令
func userExitDoor() {}
