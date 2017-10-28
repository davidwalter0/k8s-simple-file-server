package main

import (
	"fmt"
	strftime "github.com/jehiah/go-strftime"
	"net/http"
	"os"
	"strings"
	"time"
)

var Build string
var Commit string

// var Info *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
// var Info *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func Now() string {
	format := "%Y.%m.%d.%H.%M.%S.%z"
	now := time.Now()
	return strftime.Format(format, now)
}

// var prefix = Now()
var prefix = ""

func init() {
	array := strings.Split(os.Args[0], "/")
	me := array[len(array)-1]
	fmt.Println(me, "version built at:", Build, "commit:", Commit)
}

func main() {

	fmt.Println("This can be configured with the following environment variables.")
	fmt.Println("SIMPLE_FILE_SERVER_PATH to set the path to serve. The default path is /target.")
	fmt.Println("SIMPLE_FILE_SERVER_PORT to set the port to serve on. The default port is 8080")
	fmt.Println("SIMPLE_FILE_SERVER_HOST to set the host interface to serve. The default is all.")

	path := os.Getenv("SIMPLE_FILE_SERVER_PATH")
	port := os.Getenv("SIMPLE_FILE_SERVER_PORT")
	host := os.Getenv("SIMPLE_FILE_SERVER_HOST")
	if len(path) == 0 || strings.Contains(path, "../") { // || path == "/" {
		fmt.Println("SIMPLE_FILE_SERVER_PATH not set or using .. to subvert paths, using /target instead")
		path = "/var/opt/simple-file-server/data"
	} else {
		fmt.Println("SIMPLE_FILE_SERVER_PATH=" + path)
	}
	if len(port) == 0 {
		fmt.Println("SIMPLE_FILE_SERVER_PORT not set, using 8080")
		port = "8080"
	} else {
		fmt.Println("SIMPLE_FILE_SERVER_PORT=" + port)
	}
	if len(host) == 0 {
		fmt.Println("SIMPLE_FILE_SERVER_HOST not set, default bind all")
		host = "0.0.0.0"
	} else {
		fmt.Println("SIMPLE_FILE_SERVER_HOST=" + host)
	}
	listen := host + ":" + port

	http.Handle("/", http.FileServer(http.Dir(path)))

	fmt.Println("PATH to serve  " + ":" + path)
	fmt.Println("PORT on which  " + ":" + port)
	fmt.Println("HOST interface " + ":" + host)
	fmt.Println("listening on " + listen + " and serving " + path)

	err := http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Println(err)
	}
}
