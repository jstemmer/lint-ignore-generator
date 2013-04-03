package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <input.xml> <output.xml>\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("reading %s and writing %s\n", os.Args[1], os.Args[2])
}
