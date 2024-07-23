package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// echo11()
	echo12()
}
func echojoin() string {
	return strings.Join(os.Args[1:], " ")
}

func echo() string {

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	return s
}

func echo11() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
}
func echo12() {
	fmt.Println("exercise 1.2")
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}

}
