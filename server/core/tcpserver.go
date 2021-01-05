package core

import (
	"fmt"

	"github.com/fwhezfwhez/tcpx"
)

func InitTcp() {
	srv := tcpx.NewTcpX(nil)
	srv.UseGlobal(func(c *tcpx.Context) {
		fmt.Println("before raw message in")
	})

	srv.HandleRaw = response
	srv.ListenAndServeRaw("tcp", ":666")

}
func response(c *tcpx.Context) {
	var buf = make([]byte, 500)
	var n int
	var e error
	for {
		n, e = c.ConnReader.Read(buf)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println("TCP receive:%x", buf[:n])
		c.ConnWriter.Write([]byte("hello,I am server."))
	}
}
