package comlogic

import (
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"net/url"
	"strings"
)

func sendCommand(com string, command string) (body string, err error) {
	command = strings.Trim(command, " ")
	body, err = utils.GetData("http://" + global.GVA_CONFIG.System.ComIPaddr + "/Service?comname=" + com + "&command=" + url.QueryEscape(command))
	return body, err

}

//SetZero 清空置0操作
func SetZero(cabinetName string) {
	cabinet := strings.Split(cabinetName, "-")
	retCommand := "37 "
	retCommand += cabinetAdd30(cabinet[1]) + " "
	retCommand += cabinetAdd30(cabinet[2]) + " "
	retCommand += "00 00 00 00 00 00 00 00 00 00 00 00 00 00 "
	retCommand = buildWholeCmd(retCommand)
	sendCommand(cabinet[0], retCommand)

}

//InitProduct 初始化产品
func InitProduct(cabinetName string, productNum int) {
	cabinet := strings.Split(cabinetName, "-")
	retCommand := "36 "
	retCommand += cabinetAdd30(cabinet[1]) + " "
	retCommand += cabinetAdd30(cabinet[2]) + " "
	retCommand += "00 00 00 00 00"
	retCommand += hexAddPreZero(intToHexString(productNum), 2)
	retCommand += " A1 00 00 00 00 00 01"
	retCommand = buildWholeCmd(retCommand)
	sendCommand(cabinet[0], retCommand)
}

func UpdateAllProd() {}
