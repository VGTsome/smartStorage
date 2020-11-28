package comlogic

import (
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"net/url"
	"strconv"
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
	box = boxSToHex(cabinet[2])
	return com, shelf, box
}

//SetZero 清空置0操作
func SetZero(cabinetName string) {

	com, shelf, box := splitCabinetName(cabinetName)
	retCommand := "33 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "31 00 00 00 00 00"
	retCommand = buildWholeCmd(retCommand)
	//

	sendCommand(com, retCommand)
	SetLight(cabinetName, "31")

}
func turnOffLight(cabinetName string) {
	com, shelf, box := splitCabinetName(cabinetName)
	retCommand := "35 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "31 30 00 00 00 00 00"
	retCommand = buildWholeCmd(retCommand)
	sendCommand(com, retCommand)
}

//SetLight 设置灯 30灭 31红 32绿
func SetLight(cabinetName string, light string) {
	if light == "30" {
		turnOffLight(cabinetName)
	} else {
		turnOffLight(cabinetName)
		com, shelf, box := splitCabinetName(cabinetName)
		retCommand := "35 "
		retCommand += shelf + " "
		retCommand += box + " "
		retCommand += "31 " + light + " 00 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(com, retCommand)
	}
}

//InitProduct 初始化产品
//保留1 功能号01
//保留2 产品数<15
func InitProduct(cabinetName string, productNum int) {
	com, shelf, box := splitCabinetName(cabinetName)

	retCommand := "32 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "31 01"
	retCommand += hexAddPreZeroRerverse(intToHexString(productNum), 1)
	retCommand += " 00 00 00"
	retCommand = buildWholeCmd(retCommand)

	sendCommand(com, retCommand)
	SetLight(cabinetName, "31")
}

//设置货物参数
//保留1 00
func downSetProductInfo(cabinetName string, singleWeight int) {
	com, shelf, box := splitCabinetName(cabinetName)

	retCommand := "36 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "31 31"
	//单重
	retCommand += hexAddPreZeroRerverse(intToFloatHexString(singleWeight), 4)
	//皮重
	retCommand += hexAddPreZeroRerverse(intToFloatHexString(0), 4)
	//满盒数量
	retCommand += hexAddPreZeroRerverse(intToHexString(1), 4)
	retCommand += hexAddPreZeroRerverse(intToHexString(0), 5)
	retCommand = buildWholeCmd(retCommand)

	sendCommand(com, retCommand)

}

//UpdateAllProd 获取所有货物数量
//预备更新 保留位00
//更新货物库存 保留位01
//进门前发送指令1 保留位02
//出门前盘货 保留位03
func UpdateAllProd(reserv string, shelfID int) {
	precedure := "30"
	isret := "30"
	COMret := comdict["cabinet"]
	shelfRet := cabinetAdd30(strconv.Itoa(shelfID))
	if reserv == "00" {
		precedure = "30"
		isret = "31"

		retCommand := "37 " + shelfRet + " " + precedure + " " + isret + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)
		_, ssclist, _ := service.GetSmartStorageCabinetListByShelf(COMret, strconv.Itoa(shelfID))
		for _, ssc := range ssclist {
			SetLight(ssc.CabinetName, "31")

		}
	} else if reserv == "01" {
		precedure = "31"
		isret = "31"
		_, ssclist, _ := service.GetSmartStorageCabinetListByShelf(COMret, strconv.Itoa(shelfID))

		retCommand := "37 " + shelfRet + " " + precedure + " " + isret + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)
		for _, ssc := range ssclist {
			SetLight(ssc.CabinetName, "32")

		}
	} else if reserv == "02" {
		precedure = "31"
		isret = "31"

		retCommand := "37 " + shelfRet + " " + precedure + " " + isret + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)
	} else if reserv == "03" {
		precedure = "31"
		isret = "31"

		retCommand := "37 " + shelfRet + " " + precedure + " " + isret + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)
	}

}
