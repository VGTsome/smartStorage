package comlogic

import (
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"strings"
)

func enterScanFaceID(com string, command []string) {
	scanID := "111"
	_, user := service.GetUserByScanID(scanID)
	_, sso := service.GetSmartStorageOrder(user.ID)
	//service.TruncateSmartStoragePassWeight()
	_, sscps, _ := service.GetAllCabinetProductList()
	for _, sscp := range sscps {
		sspw := model.SmartStoragePassWeight{OrderId: sso.OrderId, CabinetId: sscp.CabinetId, ProductNumber: sscp.ProductNumber, ProductWeight: sscp.Weight, Pass: 0}
		service.CreateSmartStoragePassWeight(sspw)
	}
	openDoor()

}

//UpdateAllProd 获取所有货物数量
//更新货物库存 保留位01
//进门前盘货 保留位02
//出门前盘货 保留位03
func upExitDoor(com string, command []string) {
	shelf, _ := getCabinetAttr(command)
	_, sspws, _ := service.GetSmartStoragePassWeightListByShelf(com, shelf)

	for _, sspw := range sspws {
		_, ssc := service.GetSmartStorageCabinetByCabinetID(sspw.CabinetId)
		boxNumStr := strings.Split(ssc.CabinetName, "-")[2]
		proNum := getBoxNumFromCommand(boxNumStr, command)
		if proNum != sspw.ProductNumber {
			_, ssos := service.GetSmartStorageOrderByOrderID(sspw.OrderId)
			for _, sso := range ssos {
				_, sscp := service.GetCabinetProductByProductId(sso.ProductId)
				if sscp.CabinetId == sspw.CabinetId {
					if proNum+sso.OrderNumber == sspw.ProductNumber {
						sspw.Pass = 1
						service.UpdateSmartStoragePassWeight(&sspw)
					}
				}
			}

		} else {
			sspw.Pass = 1
			service.UpdateSmartStoragePassWeight(&sspw)
		}

	}
	_, _, total := service.GetSmartStoragePassWeightZeroList(sspws[0].OrderId)
	if total == 0 {
		openDoor()
	}

}
