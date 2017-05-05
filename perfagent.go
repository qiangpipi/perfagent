package main

import (
	"net/http"
	"perfagent/functions"
	. "perfagent/logs"
	"runtime"
)

func main() {
	//Use multiple cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
	//Params
	//pid: id of process
	//pname: name of process
	//pid will be used if pid and pname existing simultaneous
	//"Please provide params" if no params in request
	http.HandleFunc("/ps", functions.Ps)
	http.HandleFunc("/plist", functions.Plist)
	//Params
	//pid: id of process
	//pname: name of process
	//pid will be used if pid and pname existing simultaneous
	//duration: duration to monitor; default 5 minutes
	//freq: freq to fetch data; default 5 seconds
	//"Please provide params" if no params in request
	http.HandleFunc("/monitor", functions.Monitor)
	srv := &http.Server{Addr: "0.0.0.0:9999", Handler: nil}
	err := srv.ListenAndServe()
	if err != nil {
		Error("Start http server failed.")
	}
}
