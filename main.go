package main

import (
	"bufio"
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

	if !fileExists(fileName) {
		fmt.Println(fmt.Errorf("wc: %s: open: No such file or directory", fileName).Error())
		return
	}

	switch option {
	case "-c":
		printFileSize(fileName)
	case "-l":
		printNumberOfLines(fileName)
	default:
		fmt.Printf("ccwc: illegal option -- %s\n", option)
		fmt.Println("usage: ccwc [-Lclmw] [file ...]")
	}
}

func printNumberOfLines(fileName string) {
	scanner, file := readFile(fileName)

	defer file.Close()

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(lineCount, fileName)
}

func printFileSize(fileName string) {
	fileInfo, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fileInfo.Size(), fileName)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	return !info.IsDir()
}

func readFile(fileName string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file
}
