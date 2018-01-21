package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"utils"
)

const (
	SERVER_PORT = "8008"
	SERVER_LIST = "serverlist.prop"
)

func main() {
	ipList := []string{}
	file, _ := os.Open(SERVER_LIST)
	scanner := bufio.NewScanner(file)

	//Compile list of ip address from serverlist.prop
	for scanner.Scan() {
		var ip = scanner.Text()
		ip = ip + ":" + SERVER_PORT
		ipList = append(ipList, ip)
	}

	t0 := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("ERROR: Not enough arguments.")
		fmt.Println("Usage: client.go -options keywordToSearch")
		fmt.Println("		-options: available in linux grep command")
		os.Exit(1)
	}
	 
	c := make(chan string)

	serverInput := os.Args 

	// Connect to every server in serverlist.prop
	for _, v := range ipList {
		go utils.SendToServer(v, serverInput, c)
	}

	// Print results from server
	for i := range ipList {
		serverResult := <-c
		fmt.Println(serverResult)
		fmt.Printf("END----------%d\n", i)
	
	}

	t1 := time.Now()
	fmt.Print("Grep search took: ")
	fmt.Println(t1.Sub(t0))
}

