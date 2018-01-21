package tests

import (
    "testing"
    "os"
    "bufio"
    "strings"
    "utils"
)

const (
	SERVER_LIST = "serverlist.prop"
	TEST_PORT  =  "8009"
)
var(
	ipList []string
)

func TestGenerateLogs(t *testing.T) {
	ipList = getIpList()
	
	c := make(chan string)
	
	for _, v := range ipList {
		go utils.SendToServer(v, []string{"test","writeLogs"}, c)
	}	
	
	for i:=0; i<len(ipList) ;i++{
		serverResult := <- c  // just a dummy assignment.
		t.Log("Server Result", serverResult)
	}	
}

func TestGrep(t *testing.T) {
	cases := []struct {
        name     string
        keyword  string
        expected int
    }{
        {"frequent", "runGrep f", 100},
        {"somewhat", "runGrep s", 25},
        {"rare", "runGrep r", 1},
    }
	
	c := make(chan string)
	ipList = getIpList()
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
	        go utils.SendToServer(ipList[0], []string{"test", tc.keyword}, c)
			serverResult := <- c
            if !strings.Contains(serverResult,string(tc.expected)) {
                t.Fatalf("expected wordcount from machine %s is %d, but got %s", ipList[0], tc.expected, serverResult)
            }
        })
    }
}

func getIpList() []string{
	ipList := []string{}
	file, _ := os.Open(SERVER_LIST)
	scanner := bufio.NewScanner(file)

	//Compile list of ip address from serverlist.prop
	for scanner.Scan() {
		var ip = scanner.Text()
		ip = ip + ":" + TEST_PORT
		ipList = append(ipList, ip)
	}
	
	return ipList
}
