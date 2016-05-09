// nws
package network_service

import (
	"bytes"
	"fmt"
	"net"
	"util"
)

// 读操作
func ReadRoutine(conn *net.Conn, reader *chan bytes.Buffer) {
	// todo

	msg_feedback := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		// 因为采用了独立的读协程，此处采用阻塞读
		transferred, err := (*conn).Read(buf[0:])
		util.CheckErrorOrExit(err)

		msg_feedback.Write(buf[0:transferred])

		// todo 这里假设没有粘包或者包不足的情况
		*reader <- *msg_feedback

		fmt.Println("reseted")
		// 处理完一个msg，需要重置
		msg_feedback.Reset()
	}
}

func WriteToServerRoutine(conn *net.Conn, writer *chan bytes.Buffer) {
	// todo

	for {
		msg := <-*writer
		msg2 := (bytes.Buffer)(msg)
		(*conn).Write(msg2.Bytes())
	}
}
