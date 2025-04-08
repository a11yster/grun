package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("gRUN: graph runner")

	if len(os.Args) < 2 {
		fmt.Println("Usage: gRUN <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		fmt.Println("Running the graph...")
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
