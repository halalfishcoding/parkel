package main

import (
	"fmt"
	"os"
	"os/exec"
)

func initContainer () bool {
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


func main () {
	res := initContainer()
	if res {
		fmt.Println("Complete.")
	}
	fmt.Println("Init")
}