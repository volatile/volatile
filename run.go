package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/whitedevops/atexit"
)

var (
	lastMod     time.Time
	buildScript = "./build"

	errIsMod = errors.New("app is modified")
)

func cmdRun() {
	if !isVolatile(false) {
		errorExit("no runnable app detected")
	}

	name, err := os.Getwd()
	if err != nil {
		errorExit("faild getting current directory name")
	}
	name = filepath.Base(name)

	atexit.Use(func() {
		os.Remove(name)
	})

	var args []string
	if len(flag.Args()) > 1 {
		args = flag.Args()[1:]
	}
	build(name, args...)
}

func build(name string, args ...string) {
	// Run executable local script
	fi, err := os.Stat(buildScript)
	if err == nil && !fi.IsDir() {
		script := exec.Command(buildScript)
		script.Stdout = os.Stdout
		script.Stderr = os.Stderr
		if err = script.Run(); err != nil {
			log.Println(err)
		}
	}

	// Prepare building
	buildCmd := exec.Command("go", "build", "-o", name)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	// Prepare running
	runCmd := exec.Command("./"+name, args...)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr

	// Build and run
	running := false
	if err = buildCmd.Run(); err == nil {
		runCmd.Start()
		running = true
	}

	// Modification detection
	lastMod = time.Now()
ModDetect:
	for {
		if err = filepath.Walk(".", isMod); err != nil {
			if err == errIsMod {
				break ModDetect
			}
			errorExit("fail detecting modifications")
		}
		time.Sleep(500 * time.Millisecond)
	}

	// Interrupt
	if running {
		if err = runCmd.Process.Signal(os.Interrupt); err != nil {
			runCmd.Process.Kill()
		}
	}

	// Rerun
	log.Print("Rerunning server…\n\n")
	build(name, args...)
}

// isMod is a walk function to detect app modifications.
func isMod(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fi.ModTime().After(lastMod) {
		return errIsMod
	}
	return nil
}
