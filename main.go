package main

import (
	"time"
	"os"
	"os/exec"
	"fmt"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please feed to the program another program to execute")
		os.Exit(1)
	}
	start := time.Now()
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError during execution: \n\t %v\n\n", err.Error())
		os.Exit(1)
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nstdout: \n\t %v\n\n", err.Error())
	} else if len(output) > 0 {
		fmt.Println(string(output))
	}

	
	defer fmt.Println(time.Since(start).String())
}