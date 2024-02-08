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
	fileName := os.Args[2]

	switch option {
	case "-c":
		printFileSize(fileName)
	default:
		fmt.Printf("ccwc: illegal option -- %s\n", option)
		fmt.Println("usage: ccwc [-Lclmw] [file ...]")
	}
}

func printFileSize(fileName string) {
	fileInfo, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(fmt.Errorf("wc: %s: open: No such file or directory", fileName).Error())
		return
	}

	fmt.Println(fileInfo.Size(), fileName)
}
