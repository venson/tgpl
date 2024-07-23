package dup

import (
	"os"
	"testing"
)

func TestDup(t *testing.T) {

	originalArgs := os.Args

	// Set os.Args to contain only the program name and the test file arguments
	os.Args = append([]string{originalArgs[0]}, "./test1.txt", "./test2.txt")
	dup2()
}
