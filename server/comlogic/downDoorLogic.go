package comlogic

import (
	"encoding/hex"
	"fmt"
	"gin-vue-admin/global"
	"net"
	"strings"
	"time"
)

var doorCommand = map[string]string{
	"openDoor":       "01 05 00 00 FF 00 8C 3A",
	"closeDoor":      "01 05 00 00 00 00 CD CA",
	"getDoorStatus":  "01 01 00 00 00 01 FD CA",
	"getDoorStatusR": "01 01 01 03 11 89",
	"openDoorR":      "01 05 00 00 FF 00 8C 3A",
	"closeDoorR":     "01 05 00 01 FF 00 DD FA",
}

func sendDoorCommand(command string) (body string, err error) {
	body = ""
	commandHex, _ := hex.DecodeString(command)
	conn, err := net.DialTimeout("tcp", global.GVA_CONFIG.System.DoorIPaddr, 2*time.Second)
	tryconn := 1
	for err != nil {
		time.Sleep(time.Second * 1)
		conn, err = net.DialTimeout("tcp", global.GVA_CONFIG.System.DoorIPaddr, 2*time.Second)
		tryconn++
		if tryconn == 20 {
			return "err", err
		}
	}
	_, err = conn.Write(commandHex)
	if err != nil {
		return "write err", err
	}
	body = tcpDoorRecv(conn, command)
	return body, err
}
func tcpDoorRecv(conn net.Conn, command string) string {
	var buf = make([]byte, 500)
	try := 0
	n, err := conn.Read(buf)
	if err != nil {
		e, ok := err.(net.Error)
		if !ok || !e.Temporary() || try >= 3 {
			return "error"
		}
		try++
	}
	hexStr := ""
	for i := 0; i < n; i++ {
		hexStr += fmt.Sprintf("%02x ", buf[i])
	}
	retCmd := strings.Trim(strings.ToUpper(hexStr), " ")
	switch retCmd {
	case doorCommand["openDoorR"]:
		return "openSuccess"
	case doorCommand["closeDoorR"]:
		return "closeSuccess"
	case doorCommand["getDoorStatusR"]:
		return "getDoorStatusR"
	default:
		return "error"

	}

}
func openDoor(try int) string {
	if try > 5 {
		return "error"
	}
	ret, err := sendDoorCommand(doorCommand["openDoor"])
	if err != nil {
		try++
		time.Sleep(time.Second * 1)
		return openDoor(try)
	}
	if ret == "openSuccess" {
		return ret
	} else {
		try++
		time.Sleep(time.Second * 1)
		return openDoor(try)

	}

}
func closeDoor(try int) string {
	if try > 5 {
		return "error"
	}
	ret, err := sendDoorCommand(doorCommand["closeDoor"])
	if err != nil {
		try++
		time.Sleep(time.Second * 1)
		return closeDoor(try)

	}
	if ret == "closeSuccess" {
		return ret
	} else {
		try++
		time.Sleep(time.Second * 1)
		return closeDoor(try)

	}
}
