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
		if command[1] == "37" {

			upSetZero(com, command)
		}
		//查重
		if command[1] == "36" {
			//设置第一个货物
			if command[11] == "A1" {
				upSetFirstProd(com, command)
			}
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
	ssp.ProductWeight = singleWeight
	service.UpdateSmartStorageProduct(&ssp)

}
