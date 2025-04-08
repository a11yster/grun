package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("gRUN <--> GRAPH RUNNER CLI")

	if len(os.Args) < 2 {
		fmt.Println("Usage: gRUN <command>")
		os.Exit(1)
	}
}
