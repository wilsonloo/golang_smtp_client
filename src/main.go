// main
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
	"util"
)

func main() {
	fmt.Println("Hello World!")

	conn, err := net.Dial("tcp", "smtp.163.com:25")
	util.CheckErrorOrExit(err)
	defer conn.Close()

	msg_feedback := bytes.NewBuffer(nil)
	var buf [512]byte
	for {

		conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
		transferred, err := conn.Read(buf[0:])

		if err != nil {
			if err == io.EOF {
				break
			}

			// other error handlers
			// ...
			break
		}

		msg_feedback.Write(buf[0:transferred])
	}

	fmt.Println("message from server is: ")
	msg_feedback.WriteTo(os.Stdout)
}
