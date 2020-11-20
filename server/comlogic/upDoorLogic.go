package comlogic

import (
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"strconv"
	"strings"
)

func EnterScanFaceID(com string, command []string) {
	_, _, total := service.GetSmartStorageCurrentOrderInfoAllList()
	if total > 0 {
		return
	}
	scanID := "111"
	_, user := service.GetUserByScanID(scanID)

	_, ssos := service.GetSmartStorageOrderListByUserId(user.ID)
	err, sspw := service.GetSmartStoragePassWeightFirst()
	shelflist := make(map[string]bool)
	//判断是否有数据在passweight中
	if err != nil {
		if sspw.OrderId == ssos[0].OrderId {
			_, sspws, _ := service.GetSmartStoragePassWeightStatusList(0)
			for _, sspw := range sspws {
				_, ssc := service.GetSmartStorageCabinetByCabinetID(sspw.CabinetId)
				shelf := strings.Split(ssc.CabinetName, "-")[1]
				if _, ok := shelflist[shelf]; !ok {
					shelflist[shelf] = true
					shelfNum, _ := strconv.Atoi(shelf)
					UpdateAllProd("02", shelfNum)
				}
			}

		}
		return
	}

	//创建CurrentOrder
	for _, ssoEach := range ssos {
		_, ssp := service.GetSmartStorageProductByProductId(ssoEach.ProductId)
		_, sscps := service.GetCabinetProductListByProductId(ssoEach.ProductId)

		cabinetNameListStr := ""
		for _, sscp := range sscps {
			_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)
			cabinetNameListStr += ssc.CabinetName + ","
		}
		cabinetNameListStr = strings.Trim(cabinetNameListStr, ",")
		ssco := model.SmartStorageCurrentOrder{OrderId: ssoEach.OrderId, UserId: user.ID, ProductId: ssoEach.ProductId, ProductName: ssp.ProductName, CabinetList: cabinetNameListStr, OrderNumber: ssoEach.OrderNumber, OrderStatus: ssoEach.OrderStatus}
		service.CreateSmartStorageCurrentOrder(ssco)
	}
	//创建PassWeight
	_, sscps, _ := service.GetAllCabinetProductList()

	for _, sscp := range sscps {

		sspw := model.SmartStoragePassWeight{OrderId: ssos[0].OrderId, CabinetId: sscp.CabinetId, ProductNumber: sscp.ProductNumber, ProductNewNumber: 0, Pass: 0}
		service.CreateSmartStoragePassWeight(sspw)
		_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)
		shelf := strings.Split(ssc.CabinetName, "-")[1]
		if _, ok := shelflist[shelf]; !ok {
			shelflist[shelf] = true
			shelfNum, _ := strconv.Atoi(shelf)
			UpdateAllProd("02", shelfNum)
		}

	}

}
func UpExitDoor(com string, command []string) {
	_, sscps, _ := service.GetAllCabinetProductList()
	shelflist := make(map[string]bool)
	for _, sscp := range sscps {
		_, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)
		shelf := strings.Split(ssc.CabinetName, "-")[1]
		if _, ok := shelflist[shelf]; !ok {
			shelflist[shelf] = true
			shelfNum, _ := strconv.Atoi(shelf)
			UpdateAllProd("03", shelfNum)
		}

	}
}
