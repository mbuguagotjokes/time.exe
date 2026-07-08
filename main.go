package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please feed to the program another program to execute")
		os.Exit(1)
	}
	
	start := time.Now()
	cmd := exec.Command(args[0], args[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nStdoutPipe Err: \n\n\t %v\n\n", err)
		os.Exit(1)
	}
	
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError during execution: \n\n\t %v\n\n", err)
		os.Exit(1)
	}

	output, err := io.ReadAll(stdout)
	if err != nil {
	    fmt.Fprintf(os.Stderr, "\nError reading stdout: \n\n\t%v\n\n", err)
	    os.Exit(1)
	}
	
	if err := cmd.Wait(); err != nil {
	  	fmt.Fprintf(os.Stderr, "\nError during execution: \n\n\t %v\n\n", err)
		os.Exit(1)
	}

	if len(output) > 0 {
		fmt.Printf("\nStdout: \n\n\t %v\n\n", string(output))
	} else {
		fmt.Println("Run succesfuly")
	}

	
	defer func() {
		fmt.Println(time.Since(start).String())
	} ()
}