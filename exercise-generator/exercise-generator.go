package main

import (
	"fmt"
	"os"
	"strconv"
)

var exercise string
var n int
var m int

func main() {

	for i, v := range os.Args[1:] {
		recognise_parameter(i, v)
	}
	fmt.Printf("n: %v\n", n)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("exercise: %v\n", exercise)
}

func recognise_parameter(i int, v string) {
	switch v {
	case "-help":
		printhelp()
	case "-n":
		n, _ = strconv.Atoi(os.Args[i+2]) //+2 due to os.Args[1:]
	case "-m":
		m, _ = strconv.Atoi(os.Args[i+2])
	default:
		if _, err := strconv.Atoi(v); err != nil {
			exercise = v
		}
	}
}

func printhelp() {
	fmt.Printf("Help will be added soon \n")
}
