package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"bufio"
	"encoding/gob"
	"utils"
)

const (
	BUF_LEN 	= 1024
	TEST_PORT   = "8009"
)
var(
	localIp string
)

func main() {
	
	listener, err := net.Listen("tcp", ":" + TEST_PORT)
	if err != nil {
		println("error listening:", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Testing server listening on port :" + TEST_PORT)
	
	localIp = utils.GetLocalIP()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept:", err.Error())
			return
		}
		
		receiveTestClientReq(conn)
	}
}

func receiveTestClientReq(conn net.Conn) {
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
	
	logName := "../machine."+localIp+".log"
	strs2 := strings.Join(strs," ")
	fmt.Println("Join String : " + strs2)
	if(strings.Contains(strs2, "writeLogs")){
		generateLogs(localIp)
		results = "Logs generated on "  + localIp
	}else if (strings.Contains(strs2, "runGrep r")){
		results = utils.ExecGrep([]string{"-c","rare"}, logName, localIp)
	}else if (strings.Contains(strs2, "runGrep f")){		
		results = utils.ExecGrep([]string{"-c", "frequent"}, logName, localIp)
	}else if (strings.Contains(strs2, "runGrep s")){
		results = utils.ExecGrep([]string{"-c", "somewhat"}, logName, localIp)			
	}
	fmt.Println("Result ", results)	
	//Send data to test client	
	sendBuf := make([]byte, len(results))
	copy(sendBuf, string(results))
	conn.Write(sendBuf)
	conn.Close()
}

func generateLogs(ip string) {
	fmt.Println("Genearate Logs")
	rare := "rare"
	frequent := "frequent"
	sometimes := "somewhat"
	random := "random"

	var lines = []string{}

	for i := 0; i < 100; i++ {
		lines = append(lines, frequent)
		lines = append(lines,"\n")
		lines = append(lines, random)
		lines = append(lines, "\n")
		if i%4 == 0 {
			lines = append(lines, sometimes)
			lines = append (lines,"\n")
		}
	}
	lines = append(lines, rare)
	lines = append (lines,"\n")
	writeLines(lines, "../machine."+ip+".log")
}

/*
 * Creates a test log file and appends the required strings
 * @param lines a slice of lines that we will append to the given file name
 * @param fileName of the test log file we are going to create
 */
func writeLines(lines []string, fileName string) error {
	fmt.Println("Write lines to file")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
