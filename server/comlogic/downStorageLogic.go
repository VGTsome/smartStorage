package comlogic

import (
	"encoding/hex"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"net"
	"strconv"
	"strings"
	"time"
)

func tcpRecv(conn net.Conn) {

	var buf = make([]byte, 500)
	try := 0
	n, err := conn.Read(buf)
	if err != nil {
		e, ok := err.(net.Error)
		if !ok || !e.Temporary() || try >= 3 {
			return
		}
		try++
	}
	hexStr := ""
	for i := 0; i < n; i++ {
		hexStr += fmt.Sprintf("%02x ", buf[i])
	}
	tatoo := fmt.Sprintf("%02x%02x", buf[n-6], buf[n-5])
	funcname := fmt.Sprintf("%02x", buf[2])
	_, sscl := service.GetSmartStorageComLogByTatoo(tatoo, funcname)
	if sscl.ID == 0 {
		return
	}
	sscl.ReceiveCmd = hexStr
	service.UpdateSmartStorageComLog(&sscl)
	CmdRoute(comdict["cabinet"], hexStr)
	buf = nil
	//conn.Close()

}
func resent(command string) {
	time.Sleep(time.Microsecond * 200)
	conn, err := net.DialTimeout("tcp", global.GVA_CONFIG.System.StorageIPaddr, 2*time.Second)
	for err != nil {
		time.Sleep(time.Second * 1)
		conn, err = net.DialTimeout("tcp", global.GVA_CONFIG.System.StorageIPaddr, 2*time.Second)

	}

	command = strings.Replace(command, " ", "", -1)

	data, _ := hex.DecodeString(command)
	_, err = conn.Write(data)
	if err != nil {
		return
	}
	var sscl model.SmartStorageComLog
	sscl.SendCmd = command
	sscl.Func = fmt.Sprintf("%02x", data[2])
	sscl.Tatoo = fmt.Sprintf("%02x%02x", data[len(data)-5], data[len(data)-4])
	sscl.Isret = fmt.Sprintf("%02x", data[4])
	//service.CreateSmartStorageComLog(sscl)
	go tcpRecv(conn)
}
func isInitGetNumber(command string, lastcmd string) bool {

	if lastcmd != "" {
		commandHex, _ := hex.DecodeString(command)
		lastcmdHex, _ := hex.DecodeString(lastcmd)
		if commandHex[2] == 0x37 && lastcmdHex[2] == 0x37 && commandHex[4] == 0x30 && lastcmdHex[4] == 0x30 {
			if commandHex[3] != lastcmdHex[3] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
func sendCommand(com string, command string) (body string, err error) {
	command = strings.Replace(command, " ", "", -1)
	commandHex, _ := hex.DecodeString(command)
	try := 1
	for canpass, lastcmd := service.GetSmartStorageComLogNotRecv(); !canpass; try++ {
		time.Sleep(time.Second * 1)
		canpass, lastcmd = service.GetSmartStorageComLogNotRecv()
		if isInitGetNumber(command, lastcmd) {
			break
		}
		if try%5 == 0 {
			resent(lastcmd)

		}
		if try == 20 {
			service.DeleteSmartStorageComLogEmpty()
			break
		}
	}
	time.Sleep(time.Microsecond * 200)
	conn, err := net.DialTimeout("tcp", global.GVA_CONFIG.System.StorageIPaddr, 2*time.Second)
	tryconn := 1
	for err != nil {
		time.Sleep(time.Second * 1)
		conn, err = net.DialTimeout("tcp", global.GVA_CONFIG.System.StorageIPaddr, 2*time.Second)
		tryconn++
		if tryconn == 20 {
			return "err", err
		}
	}

	_, err = conn.Write(commandHex)
	if err != nil {
		return "write err", err
	}
	var sscl model.SmartStorageComLog
	sscl.SendCmd = command
	sscl.Func = fmt.Sprintf("%02x", commandHex[2])
	sscl.Tatoo = fmt.Sprintf("%02x%02x", commandHex[len(commandHex)-5], commandHex[len(commandHex)-4])
	sscl.Isret = fmt.Sprintf("%02x", commandHex[5])
	service.CreateSmartStorageComLog(sscl)

	go tcpRecv(conn)

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

//Set1000 1000g标定
func Set1000(cabinetName string) {

	com, shelf, box := splitCabinetName(cabinetName)
	retCommand := "34 "
	retCommand += shelf + " "
	retCommand += box + " "
	retCommand += "31" + hexAddPreZeroRerverse(intToHexString(500), 4) + " 00 00 00 00 00"
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
		time.Sleep(time.Second * 2)
		turnOffLight(cabinetName)
	} else {

		turnOffLight(cabinetName)
		time.Sleep(time.Second * 2)
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
	retCommand += hexAddPreZero(intToFloatHexString(singleWeight), 4)
	//皮重
	retCommand += hexAddPreZeroRerverse(intToFloatHexString(0), 4)
	//满盒数量
	retCommand += hexAddPreZeroRerverse(intToHexString(100), 4)
	retCommand += hexAddPreZeroRerverse(intToHexString(0), 5)
	retCommand = buildWholeCmd(retCommand)

	sendCommand(com, retCommand)

}

//UpdateAllProd 获取所有货物数量
//预备更新 保留位00
//更新货物库存 保留位01
//进门前发送指令1 保留位02
//出门前盘货 保留位03
//04 领错物料重新init
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
		precedure = "30"
		isret = "31"

		retCommand := "37 " + shelfRet + " " + precedure + " " + isret + " " + reserv + " 00 00 00 00"
		retCommand = buildWholeCmd(retCommand)
		sendCommand(COMret, retCommand)
	} else if reserv == "04" {
		precedure = "30"
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
