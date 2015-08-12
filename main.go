package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var (
	allPackages = []string{
		"github.com/volatile/compress",
		"github.com/volatile/core",
		"github.com/volatile/cors",
		"github.com/volatile/i18n",
		"github.com/volatile/log",
		"github.com/volatile/response",
		"github.com/volatile/route",
		"github.com/volatile/secure",
		"github.com/volatile/static",
		"github.com/volatile/volatile",
	}
)

func main() {
	flag.Parse()

	if flag.Args() == nil {
		help()
		return
	}

	switch flag.Arg(0) {
	case "":
		if len(flag.Args()) == 0 && isVolatile() {
			run()
		} else {
			help()
		}
	case "help":
		help()
	case "new":
		new()
	case "run":
		run()
	case "update":
		update()
	default:
		fmt.Printf("volatile: unknown subcommand %q\nRun 'volatile help' for usage.\n", flag.Arg(0))
	}
}

func isVolatile() bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".go" {
			b, err := ioutil.ReadFile(f.Name())
			if err != nil {
				panic(err)
			}

			if bytes.Contains(b, []byte("\tcore.Run()")) {
				return true
			}
		}
	}

	return false
}
