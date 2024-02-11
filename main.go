package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	byteFlag := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	lineFlag := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	charFLag := flag.Bool("m", false, "The number of characters in each input file is written to the standard output.")
	wordFLag := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")

	flag.Parse()

	args := flag.Args()

	content, fileName := getContent(args)

	switch {
	case *byteFlag:
		fmt.Println(len(content), fileName)

	case *lineFlag:
		lines := getNumLines(content)

		fmt.Println(lines, fileName)

	case *charFLag:
		fmt.Println(len(bytes.Runes(content)), fileName)

	case *wordFLag:
		fmt.Println(len(bytes.Fields(content)), fileName)

	default:
		fmt.Println(len(content), getNumLines(content), len(bytes.Fields(content)), fileName)
	}
}

func getContent(args []string) ([]byte, string) {
	if len(args) > 0 {
		fileName := args[0]

		return readFile(fileName), fileName
	}

	return readStdin(), ""
}

func getNumLines(content []byte) int {
	count := 0

	for _, b := range content {
		if b == '\n' {
			count++
		}
	}

	return count
}

func readStdin() []byte {
	content, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	return content
}

func readFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	return content
}
