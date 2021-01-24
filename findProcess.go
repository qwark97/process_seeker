package main

import (
	"log"
	"os/exec"
)

func find(name, python, script string) int {
	cmd := exec.Command(python, script, name)

	err := cmd.Start()
	handleErr(err)

	_ = cmd.Wait()

	var stdr []byte
	exitCode := exec.ExitError{
		ProcessState: cmd.ProcessState,
		Stderr:       stdr,
	}
	log.Printf("Searched: %s\n", name)
	log.Println("Result:", exitCode.ExitCode())
	return exitCode.ExitCode()
}

func handleErr(err error) {
	if err != nil {
		log.Panic("Error occured:", err)
	}
}
