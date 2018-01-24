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

/*
 * Test to generate some random logs on the server. 
 * The log generated will have frequent/somewhatfreq/rare keywords
 */
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

/*
 * Test will exec multiple grep on the logs generated in Test1 and compare the returned result.
 */
func TestGrep(t *testing.T) {
	cases := []struct {
        name     string
        keyword  string
        expected string
    }{
        {"frequent", "runGrep f", "\n100"},
        {"somewhat", "runGrep s", "\n25"},
        {"rare", "runGrep r", "\n1"},
    }
	
	c := make(chan string)
	ipList = getIpList()
	for _,ip := range ipList {
	    for _, tc := range cases {
	        t.Run(tc.name, func(t *testing.T) {
		        go utils.SendToServer(ip, []string{"test", tc.keyword}, c)
				serverResult := <- c
				t.Log("Server Result: " , serverResult)
	            if !strings.Contains(serverResult,string(tc.expected)) {
	                t.Fatalf("expected wordcount from machine %s is %d, but got %s", ip, tc.expected, serverResult)
	            }
	        })
	    }
	}    
}

/*
 * Get a list of ip address from file   
 */ 
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
