package selector

import (
  "fmt"
  "exercise-generator/exercises/determinant"
  "exercise-generator/exercises/gramschmidt"
  "exercise-generator/exercises/crossproduct"
  )

func SelectExercise(exercise string, n, m int){
  switch exercise{
  case "crossproduct":
    crossproduct.Generate(exercise)
  case "determinant":
    determinant.Generate(exercise,n)
  case "gramschmidt":
    gramschmidt.Generate(exercise,n,m)
  default:
      fmt.Printf("The exercise: %v is not implemented.\n",exercise)
      fmt.Println("Use \"exercise-generator --help\" to get a list of exercises.")

  }
}
