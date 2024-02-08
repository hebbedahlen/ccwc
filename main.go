package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ccwc [-Lclmw] [file ...]")
		return
	}

	option := os.Args[1]

	switch option {
	default:
		fmt.Printf("ccwc: illegal option -- %s\n", option)
		fmt.Println("usage: ccwc [-Lclmw] [file ...]")
	}
}
