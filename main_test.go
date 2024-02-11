package main

import (
	"io"
	"os"
	"testing"
)

func TestShouldPrintNumberOfBytesInFile(t *testing.T) {
	os.Args = []string{"ccwc", "-c", "test.txt"}

	expected := "42 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNumberOfLinesInFile(t *testing.T) {
	os.Args = []string{"ccwc", "-l", "test.txt"}

	expected := "1 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNumberOfCharactersInFile(t *testing.T) {
	os.Args = []string{"ccwc", "-m", "test.txt"}

	expected := "42 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNumberOfWordsInFile(t *testing.T) {
	os.Args = []string{"ccwc", "-w", "test.txt"}

	expected := "9 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintDefaultOption(t *testing.T) {
	os.Args = []string{"ccwc", "test.txt"}

	expected := "42 1 9 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func assert(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected: %q", expected)
		t.Errorf("Actual:   %q", actual)
	}
}

// captureStdout calls a function f and returns its stdout side-effect as string
func captureStdout(f func()) string {
	defer func(orig *os.File) {
		os.Stdout = orig
	}(os.Stdout)

	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}
