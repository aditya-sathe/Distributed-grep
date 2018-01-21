# Distributed-grep
Log quering on distributed machines

This project queries the log files which are distributed over multiple machines. Steps to run the project on linux machines. We tested on Amazon EC2 instances. 

###### Start the servers
Follow the following steps on **all machines**
- Install Golang on 
- Upload all files to location say `projects/src` 
- Set the GOPATH the the location of projects folder example `export GOPATH=$HOME/projects`
- Start the server program. Command
```
projects/src/server] go run server.go <logFileLocation>
```
- Please specify the complete path of log file for `LogFileLocation`

###### Start the client 
Follow these steps on any one machine
- Edit the `client/serverlist.prop` and specify ip addresses of all server machines
- Change directory to `projects/src/client`
- Run command
```
go run client.go <options> <searchKeyword>
```
- Specify the options available for `grep` for <options>
- Specify the keyword to search for the logs
- Eg to search for word count of `error` in all logs
```
go run client.go -c error
```  
