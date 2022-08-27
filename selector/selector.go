package selector

import (
  "fmt"
  "exercise-generator/exercises"
  )

func SelectExercise(exercise string, n, m int){
  fmt.Printf("Exercise: %v, n: %v, m: %v\n", exercise, n, m)

  switch exercise{
  case "determinant":
      determinant.Generate(exercise,n)
    default:
      fmt.Printf("The exercise: %v is not implemented.\n",exercise)

  }
}
