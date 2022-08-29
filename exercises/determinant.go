 package determinant

import (
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
  "fmt"
)



var exercise string
var solution string
// Generates matrix with easy determinant
func generate2Mat() {
	var mat mathtypes.Matrix
		for i := 0; i < 2; i++ {
			var x []mathtypes.Number
			for j := 0; j < 2; j++ {
				x = append(x, mathtypes.Random())
			}
			mat = append(mat, x)
		}
    exercise+="$"+mat.Tex()+"$"
    solution+="$\\det("+mat.Tex()+")="
    solution+=mat[0][0].Tex()+" \\cdot "+mat[1][1].Tex() +"-"+ mat[0][1].Tex() +"\\cdot" + mat[1][0].Tex()+"="
    solution+=mat.Det().Tex()+"$"
}



func generate3Mat() {
  mat := mathtypes.Get3Matrix()
  exercise+="$"+mat.Tex()+"$"
  solution+="$\\det("+mat.Tex()+") \\\\ \n"
  solution+=mat.DetTex()+"$"
}

func generate4Mat() {
  mat:= mathtypes.Get3Matrix()
  j:= mathtypes.RandomInt(0,3) //Position of new Row
  ic:=mathtypes.RandomInt(0,2) //Position of copied Column for new Column
  p:=mathtypes.Number(mathtypes.RandomIntExcept(-2,2,[]int{0})) //  Different Element for new Column
  mul:= mathtypes.Number(mathtypes.RandomIntExcept(-2,2, []int{0})) //Multiplayer new Column
  i:= mathtypes.RandomInt(0,3) //Position of new Column
  col := mathtypes.GetVector(3)
  mat = mathtypes.AddColumn(mat,col,j)
  row := make(mathtypes.Vector, len(mat[ic]))
  copy(row,mathtypes.Vector(mat[ic]))
  if ic >= i{
    ic+=1
  }
  row = mathtypes.ScalarMultiplication(mul,row)
  row[j]=row[j]+mathtypes.Number(p)
  row2 :=make(mathtypes.Vector, len(mat[0]))
  row2[j]=p
  mat2 := mathtypes.AddRow(mat,row2,i)
  mat = mathtypes.AddRow(mat,row,i)
  fmt.Printf("Test1 %v",mat.Tex())
  exercise+="$"+mat.Tex()+"$"
  solution+="Add "+mul.Tex() +" times row "+mathtypes.Number(ic+1).Tex()+" to row "+mathtypes.Number(i+1).Tex()+": \\\\ \n"
  solution+="$\\det("+mat.Tex()+")=\\det("+mat2.Tex()+")$ \\\\"

}

func Generate(name string,n int){
  exercise = "Determine the determinant of the matrix:\\\\ \n \\ \\\\ \n "
  switch n{
    case 2:
      generate2Mat()
    case 4:
      generate4Mat()
    default:
      generate3Mat()
  }

  textofile.SaveTofile(name,exercise,solution)

}
