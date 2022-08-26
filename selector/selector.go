package selector

import (
  "fmt"
  "exercise-generator/exercises"
  )

func SelectExercise(exercise string, n, m int){
  fmt.Printf("Exercise: %v, n: %v, m: %v\n", exercise, n, m)
  determinant.Generate(n)
  switch exercise{
  case "determinant":
      fmt.Printf("determinant")
    default:
      fmt.Printf("The exercise: %v is not implemented.\n",exercise)

  }
}
