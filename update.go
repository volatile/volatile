package main

import (
	"os"
	"os/exec"
)

func cmdUpdate() {
	exec.Command("go", "get", "-u",
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
	).Run()

	os.Exit(0)
}
