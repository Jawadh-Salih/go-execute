package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	var command string
	var args []string

	// Read prefix and command from arguments
	command = os.Args[1]
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTSTP, syscall.SIGINT)

	// Wait for done signal to exit
	cmd, err := runCommand(command, args)
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
	}

	go func() {
		sigs := <-sigCh
		fmt.Println("Terminating signal recieved", sigs)
		if sigs == syscall.SIGTSTP || sigs == syscall.SIGINT {
			fmt.Println("Killing the process with id", cmd.Process.Pid)
			cmd.Process.Kill()
			os.Exit(1)
		}
	}()

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
	} else {
		fmt.Println("Command finished successfully")
	}
}

func runCommand(command string, args []string) (*exec.Cmd, error) {
	// Create the full command string
	cmd := exec.Command(command, args...)

	// Set the command's output to be the same as the current process's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		fmt.Println("Is error", err.Error())
		return nil, err
	}

	return cmd, nil

}
