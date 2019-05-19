package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	number1()
	number2()
	number3()
}

func number1() {
	var s string

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + ";"
	}
	fmt.Println(s)
}

// os.Arg[1:] trims the array to contain only element 1 on forward
func number2() {
	var s string

	for _, arg := range os.Args[1:] {
		s += arg + ";"
	}
	fmt.Println(s)
}

// shows that os.Arg[0] is the exe name
func number3() {
	fmt.Println(strings.Join(os.Args, ";"))
}
