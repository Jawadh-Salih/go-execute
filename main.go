package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigCh
		fmt.Println(sig)

		done <- true
	}()

	fmt.Println("Awaiting signal")
	<-done
	fmt.Println("exiting")
}
