package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func new() {
	update()

	appType := "web"
	appName := "app"

	if len(flag.Args()) >= 2 {
		switch flag.Arg(1) {
		case "api":
			appType = "api"
			if len(flag.Args()) >= 3 {
				appName = flag.Arg(2) // API app name
			}
		case "web":
			if len(flag.Args()) >= 3 {
				appName = flag.Arg(2) // Web app name
			}
		default:
			appName = flag.Arg(1) // Web app name
		}
	}

	dirPath := filepath.Clean("." + string(filepath.Separator) + appName)

	// Check destination is empty.
	dst, _ := os.Open(dirPath)
	if names, _ := dst.Readdirnames(1); len(names) != 0 {
		fmt.Println("volatile new: destination must be empty")
		os.Exit(1)
	}
	dst.Close()

	exec.Command("mkdir", "-p", dirPath).Run()
	exec.Command("cp", "-R", os.Getenv("GOPATH")+"/src/github.com/volatile/volatile/bootstrap/"+appType+"/", dirPath).Run()
	if err := os.Chdir(dirPath); err != nil {
		panic(err)
	}
	run()
}
