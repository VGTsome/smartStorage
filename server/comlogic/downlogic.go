package comlogic

import (
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"net/url"
	"strings"
)

func sendCommand(com string, command string) (body string, err error) {
	command = strings.Trim(command, " ")
	body, err = utils.GetData("http://" + global.GVA_CONFIG.System.ComIPaddr + "/Service?comname=" + com + "&command=" + url.QueryEscape(command))
	return body, err

}
func splitCabinetName(cabinetName string) (com string, shelf string, box string) {
	cabinet := strings.Split(cabinetName, "-")
	com = cabinet[0]
	shelf = cabinetAdd30(cabinet[1])
	box = cabinetAdd30(cabinet[2])
	return com, shelf, box
}

//SetZero 清空置0操作
func SetZero(cabinetName string) {

	com, shelf, box := splitCabinetName(cabinetName)
	retCommand := "37 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "00 00 00 00 00 00 00 00 00 00 00 00 00 00 "
	retCommand = buildWholeCmd(retCommand)
	sendCommand(com, retCommand)

}

//InitProduct 初始化产品
func InitProduct(cabinetName string, productNum int) {
	com, shelf, box := splitCabinetName(cabinetName)

	retCommand := "36 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "00 00 00 00 00"
	retCommand += hexAddPreZero(intToHexString(productNum), 2)
	retCommand += " A1 00 00 00 00 00 01"
	retCommand = buildWholeCmd(retCommand)
	sendCommand(com, retCommand)
}

//设置货物参数
func downSetProductInfo(cabinetName string, singleWeight int) {
	com, shelf, box := splitCabinetName(cabinetName)

	retCommand := "42 "
	retCommand += shelf + " "
	retCommand += box + " "
	//单重
	retCommand += hexAddPreZero(intToHexString(singleWeight), 2)
	//皮重
	retCommand += hexAddPreZero(intToHexString(0), 2)
	//满盒数量
	retCommand += hexAddPreZero(intToHexString(1), 1)
	retCommand += hexAddPreZero(intToHexString(0), 4)
	retCommand = buildWholeCmd(retCommand)
	sendCommand(com, retCommand)

}

//获取所有货物数量
//更新货物库存 保留位01
//进门前盘货 保留位02
//出门前盘货 保留位03
func UpdateAllProd(reserv string) {
	_, smartStorageCabinetList, _ := service.GetSmartStorageCabinetList()
	var shelflist []string
	var COMret string
	for _, smartStorageCabinet := range smartStorageCabinetList {
		COM, shelf, _ := splitCabinetName(smartStorageCabinet.CabinetName)
		shelflist = append(shelflist, shelf)
		COMret = COM

	}
	newshelflist := RemoveReplicaSliceString(shelflist)
	for _, shelfRet := range newshelflist {

		retCommand := "31 " + shelfRet + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)

	}

}

func GetAllProductNumber() {}
