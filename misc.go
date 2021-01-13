package main

import (
	"fmt"
	"log"
	"time"
)

// ExitOnError does a log.Fatal if error is not nil
func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Wait waits for the specified number of seconds and prints a progress bar
func Wait(seconds int) {
	fmt.Print("Waiting ")
	for i := 0; i < seconds; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()
}
