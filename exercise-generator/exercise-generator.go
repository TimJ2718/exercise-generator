package main

import (
	"exercise-generator/selector"
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
	selector.SelectExercise(exercise, n, m)
}

func recognise_parameter(i int, v string) {
	switch v {
	case "--help":
		printhelp()
	case "-h":
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
	help:="This program generates math exercises in Latex.files.\n"
	help+="The exercise is generated with the command: \"exercise-generator NameOfExercise -n {nvalue} -m m{value}\" \n"
	help+="The exercise is saved in the file \"NameOfExercise-ex.tex\" and the solution in \"NameOfExercise-sol.tex\"\n"
	help+="n and m are optional parameter. If no value is passed, the default value is used.\n"
	help+="Exercises:\n"
	help+=" \"crossproduct\" Calculate a vector orthogonal to two other vectors. \n"
	help+=" \"determinat -n <nvalue>\"  n ∈ {2,3,4} \n   Calculate the determinat of a nxn Matrix.\n"
	help+=" \"gramschmidt -n <nvalue> -m <mvalue>\"  n>=m, n,m ∈ {2,3,4} \n   Orthogonalise a set of m vectors with the dimension n.\n"
	fmt.Printf(help)
}
