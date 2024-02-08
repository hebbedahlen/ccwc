package main

import (
	"io"
	"os"
	"testing"
)

func TestShouldPrintUsageToStdout(t *testing.T) {
	os.Args = []string{"ccwc"}

	want := "usage: ccwc [-Lclmw] [file ...]\n"
	got := captureStdout(main)

	if want != got {
		t.Errorf("Expected output %q, but got %q", want, got)
	}
}

func TestShouldPrintIllegalOption(t *testing.T) {
	os.Args = []string{"ccwc", "-d"}

	expected := "ccwc: illegal option -- -d\nusage: ccwc [-Lclmw] [file ...]\n"
	actual := captureStdout(main)

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
