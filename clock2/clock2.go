// Package clock2 is a TCP server that periodically writes the time. Adapted
// from https://github.com/adonovan/gopl.io/tree/master/ch8/clock2.
package clock2

import (
	"io"
	"net"
	"time"
)

func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(time.Second)
	}
}
