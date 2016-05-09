// main
package main

import (
	"bytes"
	"fmt"
	"net"
	nws "network_service"
	"util"
)

// smtp_msg通道
var smtp_msg_chan chan bytes.Buffer
var smtp_msg_write_chan chan bytes.Buffer

func main() {
	fmt.Println("Hello World!")

	conn, err := net.Dial("tcp", "smtp.163.com:25")
	util.CheckErrorOrExit(err)
	defer conn.Close()

	// 开启独立的读协程
	smtp_msg_chan = make(chan bytes.Buffer, 1)
	smtp_msg_write_chan = make(chan bytes.Buffer, 1)

	go nws.ReadRoutine(conn, &smtp_msg_chan)

	// 独立的写协程
	go nws.WriteToServerRoutine(conn, &smtp_msg_write_chan)

	handle_smtp_msg_routine(&smtp_msg_chan)
}

// 处理 smtp 消息
func handle_smtp_msg_routine(msg_chan *chan bytes.Buffer) {
	// 主协程处理逻辑
	for {
		fmt.Println("waiting msg...")
		msg_raw := <-smtp_msg_chan
		msg := (*bytes.Buffer)(&msg_raw)
		fmt.Println("got msg: ", msg)

		// todo 解析msg的smtp协议

	}
}
