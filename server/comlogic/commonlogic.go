package comlogic

import (
	"fmt"
	"strconv"
	"strings"
)

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

func buildWholeCmd(command string) string {

	return "02 " + addCheckSum(strings.Trim(command, " ")) + " 03"
}

func hexAddPreZero(hexstr string, zeroGroup int) string {
	hexstr = "000000000000" + hexstr
	rethex := ""
	for i := zeroGroup; i >= 1; i-- {
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
	totalstr := strconv.Itoa(total)
	hexstr := intStringToHexstring(totalstr)

	addCheck = command + hexAddPreZero(hexstr, 2)
	return addCheck
}
