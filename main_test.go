package main

import (
	"io"
	"os"
	"testing"
)

func TestShouldPrintUsageToStdout(t *testing.T) {
	os.Args = []string{"ccwc"}

	expected := "usage: ccwc [-Lclmw] [file ...]\n"
	actual := captureStdout(main)

	if expected != actual {
		t.Errorf("Expected: %q", expected)
		t.Errorf("Actual:   %q", actual)
	}
}

func TestShouldPrintIllegalOption(t *testing.T) {
	os.Args = []string{"ccwc", "-d", "test.txt"}

	expected := "ccwc: illegal option -- -d\nusage: ccwc [-Lclmw] [file ...]\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNoFile(t *testing.T) {
	os.Args = []string{"ccwc", "-c", "non-existing.txt"}

	expected := "wc: non-existing.txt: open: No such file or directory\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNumberOfBytesOfFile(t *testing.T) {
	os.Args = []string{"ccwc", "-c", "test.txt"}

	expected := "43 test.txt\n"
	actual := captureStdout(main)

	assert(expected, actual, t)
}

func TestShouldPrintNumberOfLinesOfFile(t *testing.T) {
	os.Args = []string{"ccwc", "-l", "test.txt"}

	expected := "2 test.txt\n"
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
