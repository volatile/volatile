package main

import (
	"fmt"
	"os"
)

func cmdHelp() {
	fmt.Print(`Volatile is a tool for managing Volatile app source code.

Usage:

	volatile [command] [arguments]

Without command, "volatile" recompiles and reruns the app in the current directory every time a file change.

Commands:

	new [type] [name]	bootstrap an app of type ("web" or "api") and name
	run [arguments]		recompile and rerun app with arguments every time a file change
	update			get updated versions of the Volatile CLI, the core and all official packages

`)

	os.Exit(0)
}
