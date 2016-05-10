package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/whitedevops/colors"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "":
		if len(flag.Args()) == 0 && isVolatile(false) {
			cmdRun()
		} else {
			cmdHelp()
		}
	case "help":
		cmdHelp()
	case "new":
		cmdNew()
	case "run":
		cmdRun()
	case "update":
		cmdUpdate()
	default:
		cmdHelp()
	}
}

func isVolatile(flagsRequired bool) bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		errorExit("fail reading current directory")
	}

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".go" {
			continue
		}

		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			errorExit("fail reading file " + f.Name())
		}

		if bytes.Contains(b, []byte("\nfunc main() {\n")) && bytes.Contains(b, []byte("core.Run()")) {
			if flagsRequired {
				return bytes.Contains(b, []byte("flag.Parse()"))
			}

			return true
		}
	}

	return false
}

func errorExit(err string) {
	fmt.Print("volatile")
	if cmd := flag.Arg(0); cmd != "" {
		fmt.Print(" " + cmd)
	}
	fmt.Printf(": %s%s%s\n", colors.Red, err, colors.ResetAll)

	os.Exit(1)
}
