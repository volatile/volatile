package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Print(`Volatile is a tool for managing Volatile app source code.

Usage:

    volatile [command] [arguments]

With no command, 'volatile' recompiles and reruns the app in the current directory every time a file change.

The commands are:

    new [type] [name]  bootstrap and run an app of desired type ('web' or 'api') and name
    run [args...]      recompile and rerun app with args every time a file change
    update             get updated versions of the Volatile Core and all official packages

`)
	os.Exit(0)
}
