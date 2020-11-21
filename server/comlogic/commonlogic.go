package comlogic

import (
	"fmt"
	"strconv"
	"strings"
)

var comdict = map[string]string{
	"COM1":    "cabinet",
	"COM2":    "door",
	"cabinet": "COM1",
	"door":    "COM2",
}

//CmdRoute 上行命令路由
func CmdRoute(com string, upcmd string) {
	command := strings.Split(upcmd, " ")
	switch comdict[com] {
	case "cabinet":
		//标零
		if command[2] == "33" {

			upSetZero(com, command)
		}
		//查重
		if command[2] == "32" {
			//设置第一个货物
			if command[11] == "01" {
				upSetFirstProd(com, command)
			}
		}
		//设置货品参数
		if command[2] == "36" {
			upSetSingWeight(com, command)
		}
		//更新货架货物数量
		if command[2] == "37" {
			if command[len(command)-9] == "01" {
				upUpdateCabinetProduct(com, command)
			}
			if command[len(command)-9] == "02" {

				updatePassWeightCurrentOrder(com, command)
			}
			if command[len(command)-9] == "03" {
				checkPassWeight(com, command)
			}
		}
		break
	case "door":

		break

	}
}
func intToHexString(ten int) string {
	i := int64(ten)
	s := strconv.FormatInt(i, 16)
	return s
}
func intStringToHexstring(tenstring string) string {
	ten, _ := strconv.Atoi(tenstring)
	hexStr := intToHexString(ten)
	return hexStr
}
func hexstringToNumber(hexst string) int {
	n, err := strconv.ParseUint(hexst, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)

}
func hexStringToNumberString(hexst string) string {
	num := hexstringToNumber(hexst)
	return strconv.Itoa(num)
}

func cabinetAdd30(cabinetNum string) string {
	cnum, _ := strconv.Atoi(cabinetNum)
	cnum += 48

	return intToHexString(cnum)
}

func cabinetMinus30(cabinetNum string) string {
	cnum := hexstringToNumber(cabinetNum)
	cnum -= 48
	return strconv.Itoa(cnum)
}
func getCmdStatus(cmd []string) string {
	return cmd[len(cmd)-4]
}

func buildWholeCmd(command string) string {
	command = addCheckSum(strings.Trim(command, " ")) + " "
	cmdlen := len(command) / 3
	cmdlen += 2
	cmdlenstr := hexAddPreZero(intToHexString(cmdlen), 2)
	return "02 " + cmdlenstr + command + " 03"
}
func RemoveReplicaSliceString(slc []string) []string {
	/*
	   slice(string类型)元素去重
	*/
	result := make([]string, 0)
	tempMap := make(map[string]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}
func hexAddPreZero(hexstr string, zeroGroup int) string {
	hexstr = "000000000000000000" + hexstr
	rethex := ""
	for i := 1; i <= zeroGroup; i++ {
		rethex += " " + hexstr[len(hexstr)-i*2:len(hexstr)-2*(i-1)]
	}
	return rethex

}
func addCheckSum(command string) (addCheck string) {
	command = strings.Trim(command, " ")
	commandlist := strings.Split(command, " ")
	total := 0
	for i := 1; i < len(commandlist); i++ {
		total += hexstringToNumber(commandlist[i])
	}

	hexstr := intToHexString(total)

	addCheck = command + hexAddPreZero(hexstr, 4)
	return addCheck
}
