package utils

import (
	"fmt"
	"gin-vue-admin/global"
	"strconv"
	"strings"
)

func NumberToHexstring(ten int) string {
	m := 0
	hex := make([]int, 0)
	for {
		m = ten % 16
		ten = ten / 16
		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	hexStr := []string{}
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	return strings.Join(hexStr, "")
}
func hexstringToNumber(hexst string) int {
	n, err := strconv.ParseUint(hexst, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)

}

func sendCommand(com string, command string) (body string, err error) {
	body, err = GetData("http://" + global.GVA_CONFIG.System.ComIPaddr + "/Service?comname=" + com + "&command=" + command)
	return body, err

}
func cmdRoute(upcmd string) {}
func setZero(upcmd string)  {}
