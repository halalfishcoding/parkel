package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// type Container struct {
// 	id string
// 	owner string
// 	language string
// }

func CreateContainerAuth (w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
}

// func _retrieveContainerId () string {
// 	//open file here 
// }



func InitContainer () bool {
	result := &exec.Cmd {
		Path: "sh",
		Args: []string {"sh", "/scripts/init_container.sh"}	,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if result.Stderr == nil {
		//return container id but alr
		return true
	}
	return false
}



func HttpHandler () {
	http.HandleFunc("/CreateContainer", CreateContainerAuth)
    log.Fatal(http.ListenAndServe(":10000", nil))
}


func main () {
	res := initContainer()
	if res {
		fmt.Println("Complete.")
	}
	fmt.Println("Init")
}