package determinant

import (
  "math/rand"
  "math"
  "time"
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
//  "fmt"
)



var exercise string
var solution string
// Generates matrix with easy determinant
func generate2Mat() {
	rand.Seed(time.Now().UnixNano())
	var mat mathtypes.Matrix
		for i := 0; i < 2; i++ {
			var x []mathtypes.Number
			for j := 0; j < 2; j++ {
				x = append(x, mathtypes.Number(rand.Intn(19)-9))
			}
			mat = append(mat, x)
		}
    exercise+="$"+mat.Tex()+"$"
    solution+="$\\det("+mat.Tex()+")="
    solution+=mat[0][0].Tex()+" \\cdot "+mat[1][1].Tex() +"-"+ mat[0][1].Tex() +"\\cdot" + mat[1][0].Tex()+"="
    solution+=mat.Det().Tex()+"$"
}

//create a Matrix where the abosolute value of the product of each diagonal line is smaller 150
func get3Matrix() mathtypes.Matrix{
  var max float64 = 150
  rand.Seed(time.Now().UnixNano())
  x := mathtypes.Matrix{{0,0,0},{0,0,0},{0,0,0}}
  for {
    for i :=0 ; i<3 ; i++{
      for {
        for j := 0; j <3; j++{
          x[i][j]= mathtypes.Number(rand.Intn(19)-9)
        }
        if math.Abs(float64(x[i][0]*x[i][1]*x[i][2])) < max{
          break
        }
      }
    }
    if (math.Abs(float64(x[0][0]*x[1][0]*x[2][0]))<max) && (math.Abs(float64(x[0][1]*x[1][1]*x[2][1]))<max) && (math.Abs(float64(x[0][2]*x[1][2]*x[2][2]))<max){
      break
    }
  }
  return mathtypes.Matrix{{x[0][0],x[1][0],x[2][0]},{x[2][1],x[0][1],x[1][1]},{x[1][2],x[2][2],x[0][2]}}

}

func generate3Mat() {
  mat := get3Matrix()
  exercise+="$"+mat.Tex()+"$"
  solution+="$\\det("+mat.Tex()+") \\\\ \n"
  m := " \\cdot "
  solution+="="+mat[0][0].Tex()+m+mat[1][1].Tex()+m+mat[2][2].Tex()+"+"+mat[0][1].Tex()+m+mat[1][2].Tex()+m+mat[2][0].Tex()+"+"+mat[0][2].Tex()+m+mat[1][0].Tex()+m+mat[2][1].Tex()+"\n"
  solution+="-"+mat[2][0].Tex()+m+mat[1][1].Tex()+m+mat[0][2].Tex()+"-"+mat[2][1].Tex()+m+mat[1][2].Tex() +m+mat[0][1].Tex()+"-"+mat[2][2].Tex()+m+mat[1][0].Tex()+m+mat[0][1].Tex()+"\n"

  solution +="$"
}


func Generate(name string,n int){
  exercise = "Determine the determinant of the matrix:\\\\ \n \\ \\\\ \n "
  switch n{
  case 20:
      generate2Mat()
    default:
      generate3Mat()
  }

  textofile.SaveTofile(name,exercise,solution)

}
