package comlogic

import (
	"gin-vue-admin/service"
	"strings"
)

var comdict = map[string]string{
	"COM1": "cabinet",
	"COM2": "door",
}

//CmdRoute 上行命令路由
func CmdRoute(com string, upcmd string) {
	command := strings.Split(upcmd, " ")
	switch comdict[com] {
	case "cabinet":
		//标零
		if command[2] == "37" {

			upSetZero(com, command)
		}
		//查重
		if command[2] == "36" {
			//设置第一个货物
			if command[11] == "A1" {
				upSetFirstProd(com, command)
			}
		}
		//设置货品参数
		if command[2] == "42" {
			upSetSingWeight(com, command)
		}
		break
	case "door":

		break

	}
}

//upSetZero 数据库称台清零
func upSetZero(com string, command []string) {
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + cabinetMinus30(command[2]) + "-" + cabinetMinus30(command[3]))
	sscp.ProductNumber = 0
	sscp.Weight = 0
	service.UpdateCabinetProduct(&sscp)
}

//初始化第一批物品
func upSetFirstProd(com string, command []string) {
	prodNumString := command[9] + command[10]
	prodNum := hexstringToNumber(prodNumString)
	weightString := command[4] + command[5] + command[6] + command[7] + command[8]
	weight := hexstringToNumber(weightString)

	singleWeight := weight / prodNum
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + cabinetMinus30(command[2]) + "-" + cabinetMinus30(command[3]))
	sscp.ProductNumber = prodNum
	sscp.Weight = weight
	service.UpdateCabinetProduct(&sscp)
	_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
	ssp.ProductWeight = 0
	service.UpdateSmartStorageProduct(&ssp)
	_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)

	downSetProductInfo(ssc.CabinetName, singleWeight)

}

//设置货物参数
func upSetSingWeight(com string, command []string) {
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + cabinetMinus30(command[2]) + "-" + cabinetMinus30(command[3]))
	_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
	ssp.ProductWeight = hexstringToNumber(command[4] + command[5] + command[6])
	ssp.PackageWeight = hexstringToNumber(command[7] + command[8] + command[9])
	ssp.ProductNumber = hexstringToNumber(command[10] + command[11])
	service.UpdateSmartStorageProduct(&ssp)

}

//进门前做货物检查记录
func beforeEnterDoor() {}

//开门指令
func userEnterDoor() {}

//出门前做货物检查记录
func beforeExitDoor() {}

//出门指令
func userExitDoor() {}
