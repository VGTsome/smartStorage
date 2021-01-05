package comlogic

import (
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"strconv"
	"strings"
	"time"
)

//upSetZero 数据库称台清零
func upSetZero(com string, command []string) {
	status := getCmdStatus(command)
	if status != "5a" {
		return
	}

	shelf, box := getCabinetAttr(command)
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	sscp.ProductNumber = 0
	sscp.Weight = 0
	service.UpdateCabinetProduct(&sscp)
	SetLight(com+"-"+shelf+"-"+box, "32")

}

//初始化第一批物品
func upSetFirstProd(com string, command []string) {
	status := getCmdStatus(command)
	if status != "5a" {
		return
	}
	shelf, box := getCabinetAttr(command)
	prodNumString := command[11]
	prodNum := hexstringToNumber(prodNumString)
	weightString := command[9] + command[8] + command[7] + command[6]

	weight := hexstringToNumber(weightString)

	singleWeight := weight / prodNum
	_, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	sscp.ProductNumber = prodNum
	sscp.Weight = singleWeight
	service.UpdateCabinetProduct(&sscp)
	_, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
	ssp.ProductWeight = singleWeight
	ssp.PackageWeight = 0
	service.UpdateSmartStorageProduct(&ssp)
	_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)

	downSetProductInfo(ssc.CabinetName, singleWeight)

}
func getCabinetAttr(command []string) (shelf string, box string) {
	shelf = cabinetMinus30(command[3])
	box = boxSToS(command[4])
	return shelf, box
}

//设置货物参数
func upSetSingWeight(com string, command []string) {
	status := getCmdStatus(command)
	if status != "5a" {
		return
	}
	shelf, box := getCabinetAttr(command)
	// _, sscp := service.GetCabinetProductByCabinetName(com + "-" + shelf + "-" + box)
	// _, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)

	// if ssp.ProductWeight == hexstringToNumber(command[10]+command[9]+command[8]+command[7]) && ssp.PackageWeight == hexstringToNumber(command[14]+command[13]+command[12]+command[11]) {
	SetLight(com+"-"+shelf+"-"+box, "32")
	// }
	//ssp.ProductNumber = hexstringToNumber(command[10] + command[11])

}

//更新整个柜子库存容量
//更新货物库存 保留位01
func upUpdateCabinetProduct(com string, command []string) {
	status := getCmdStatus(command)
	if status != "5a" {
		return
	}
	shelf, _ := getCabinetAttr(command)
	_, ssclist, _ := service.GetSmartStorageCabinetListByShelf(com, shelf)
	for _, ssc := range ssclist {
		_, sscp := service.GetCabinetProductByCabinetName(ssc.CabinetName)
		boxNumStr := strings.Split(ssc.CabinetName, "-")[2]
		sscp.ProductNumber = getBoxNumFromCommand(boxNumStr, command)
		// _, ssp := service.GetSmartStorageProductByProductId(sscp.ProductId)
		// sscp.Weight = ssp.ProductWeight * sscp.ProductNumber
		service.UpdateCabinetProduct(&sscp)
		SetLight(ssc.CabinetName, "30")

		//ssp.ProductNumber = sscp.ProductNumber
		//service.UpdateSmartStorageProduct(&ssp)
	}
	service.SetSystemStatus(1)
}

//检验盘货步骤1发送结果 保留位02
func updatePassWeightCurrentOrder(com string, command []string) {
	status := getCmdStatus(command)
	if status != "5a" && status != "31" {
		return
	}
	shelf, _ := getCabinetAttr(command)
	_, sspwlist, _ := service.GetSmartStoragePassWeightListByShelf(com, shelf)
	for _, sspw := range sspwlist {
		sspw.Pass = 1
		service.UpdateSmartStoragePassWeight(&sspw)
	}
	_, _, total := service.GetSmartStoragePassWeightStatusList(0)
	if total == 0 {
		setCurrentOrderLight("32")
		openDoor(0)
		time.Sleep(time.Second * 5)
		closeDoor(0)
		service.SetSystemStatus(2)
	}
}
func setCurrentOrderLight(lightCode string) {
	_, sscos, _ := service.GetSmartStorageCurrentOrderInfoAllList()
	for _, ssco := range sscos {
		cabinetList := strings.Split(ssco.CabinetList, ",")
		for _, cabinetName := range cabinetList {
			SetLight(cabinetName, lightCode)
		}

	}
}
func getBoxNumFromCommand(boxNumStr string, command []string) (proNum int) {
	boxNum, _ := strconv.Atoi(boxNumStr)
	proNumHex := command[boxNum*2+5] + command[boxNum*2+4]
	proNum = hexstringToNumber(proNumHex)
	return proNum

}

//checkPassWeight 用户出门检查货物数量
func checkPassWeight(com string, command []string) {
	status := getCmdStatus(command)
	shelf, _ := getCabinetAttr(command)
	shelfNum, _ := strconv.Atoi(shelf)
	if status == "30" {
		UpdateAllProd("04", shelfNum)
		return
	} else if status == "35" {
		UpdateAllProd("04", shelfNum)
		return
	} else if status != "5a" {
		return
	}

	_, sspws, _ := service.GetSmartStoragePassWeightListByShelf(com, shelf)

	for _, sspw := range sspws {

		_, ssc := service.GetSmartStorageCabinetByCabinetID(sspw.CabinetId)
		boxNumStr := strings.Split(ssc.CabinetName, "-")[2]

		proNum := getBoxNumFromCommand(boxNumStr, command)
		sspw.ProductNewNumber = proNum
		sspw.Pass = 2
		service.UpdateSmartStoragePassWeight(&sspw)

	}
	_, _, total := service.GetSmartStoragePassWeightStatusList(1)

	if total == 0 {
		productOld := make(map[string]int)
		productNew := make(map[string]int)
		_, sspws, _ := service.GetSmartStoragePassWeightStatusList(2)
		for _, sspw := range sspws {
			_, sscp := service.GetCabinetProductByCabinetId(sspw.CabinetId)
			if _, ok := productOld[sscp.ProductId]; !ok {
				productOld[sscp.ProductId] = sspw.ProductNumber
			} else {
				productOld[sscp.ProductId] += sspw.ProductNumber
			}
			if _, ok := productNew[sscp.ProductId]; !ok {
				productNew[sscp.ProductId] = sspw.ProductNewNumber
			} else {
				productNew[sscp.ProductId] += sspw.ProductNewNumber
			}
		}
		canPass := true
		Comments := ""
		for key, value := range productNew {

			_, sso := service.GetSmartStorageOrderByOrderIDProductId(sspws[0].OrderId, key)

			if sso.OrderNumber+value != productOld[key] {
				_, ssp := service.GetSmartStorageProductByProductId(key)
				Comments += ssp.ProductName + " 数量异常\r\n"
				canPass = false
			}

		}
		if canPass {
			service.SetSystemComments(Comments)
			openDoor(0)
			time.Sleep(time.Second * 5)
			closeDoor(0)
			updateStorageByPassWeight(sspws)
			updateOrderStatusByOrderId(sspws[0].OrderId)
			setCurrentOrderLight("30")
			service.TruncateSmartStorageCurrentOrder()
			service.TruncateSmartStoragePassWeight()
			service.SetSystemStatus(1)
		} else {
			service.SetSystemComments(Comments)
			UpdateAllProd("04", shelfNum)
			//allShelfSendInit()
		}

	}

}

func updateStorageByPassWeight(sspws []model.SmartStoragePassWeight) {
	for _, sspw := range sspws {
		_, sscp := service.GetCabinetProductByCabinetId(sspw.CabinetId)
		sscp.ProductNumber = sspw.ProductNewNumber
		service.UpdateCabinetProduct(&sscp)

	}
}

func updateOrderStatusByOrderId(orderId string) {
	_, ssos := service.GetSmartStoragesOrderByOrderID(orderId)
	for _, sso := range ssos {
		sso.OrderStatus = 10
		service.UpdateSmartStorageOrder(&sso)
	}
}
