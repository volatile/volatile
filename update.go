package main

import "os/exec"

func update() {
	cmdArgs := append([]string{"get", "-u"}, allPackages...)
	exec.Command("go", cmdArgs...).Run()
}
