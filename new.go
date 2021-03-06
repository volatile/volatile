package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
)

func cmdNew() {
	cmdUpdate()

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

	dirPath := filepath.Clean(appName)

	// Check destination is empty.
	dst, _ := os.Open(dirPath)
	if names, _ := dst.Readdirnames(1); len(names) != 0 {
		errorExit("destination is not empty")
	}
	dst.Close()

	exec.Command("mkdir", "-p", dirPath).Run()
	exec.Command("cp", "-r", os.Getenv("GOPATH")+"/src/github.com/volatile/volatile/bootstrap/"+appType+"/.", dirPath+"/.").Run()

	os.Exit(0)
}
