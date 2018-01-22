package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"encoding/gob"
	"utils"
)

const (
	BUF_LEN 		= 1024
	PORT     		= "8008"
)
var(
	localIp string
)

func main() {
	fmt.Println("Logging server listening on port :" + PORT)

	listener, err := net.Listen("tcp", ":" + PORT)
	if err != nil {
		println("error listening:", err.Error())
		os.Exit(1)
	}
	
	localIp = utils.GetLocalIP()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept:", err.Error())
			return
		}
		
		grepLog(conn)
	}
}


func grepLog(conn net.Conn) {
	recvBuf := make([]byte, BUF_LEN)
	_, err := conn.Read(recvBuf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	
	strs := []string{}
    gob.NewDecoder(bytes.NewReader(recvBuf)).Decode(&strs)
    fmt.Println("Received String: ", strs)

	var results string
	// exec the grep
	results = utils.ExecGrep(strs, os.Args[1], localIp)
	
	sendBuf := make([]byte, len(results))
	copy(sendBuf, string(results))
	conn.Write(sendBuf)
	conn.Close()
}

