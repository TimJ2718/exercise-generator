 package determinant

import (
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
  "math"
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
  mat := mathtypes.GetNXMMatrix(3,3,6)
  exercise+="$"+mat.Tex()+"$"
  solution+="$\\det("+mat.Tex()+") \\\\ \n"
  solution+=mat.DetTex()+"$"
}

func generate4Mat() {
  mat:= mathtypes.GetNXMMatrix(3,3,6)
  j:= mathtypes.RandomInt(0,3) //Position of new Row
  ic:=mathtypes.RandomInt(0,2) //Position of copied Column for new Column
  p:=mathtypes.Number(mathtypes.RandomIntExcept(-2,2,[]int{0})) //  Different Element for new Column
  mul:= mathtypes.Number(mathtypes.RandomIntExcept(-2,2, []int{0})) //Multiplayer new Column
  i:= mathtypes.RandomInt(0,3) //Position of new Column
  col := mathtypes.GetVector(3)
  mat2 := mathtypes.AddColumn(mat,col,j)
  row := make(mathtypes.Vector, len(mat2[ic]))
  copy(row,mathtypes.Vector(mat2[ic]))
  if ic >= i{
    ic+=1
  }
  row = mathtypes.ScalarMultiplication(mul,row)
  row[j]=row[j]+mathtypes.Number(p)
  row2 :=make(mathtypes.Vector, len(mat2[0]))
  row2[j]=p
  mat3 := mathtypes.AddRow(mat2,row2,i)
  mat2 = mathtypes.AddRow(mat2,row,i)
  variant := mathtypes.RandomInt(0,1) != 0
  variant = false
  m :=" \\cdot "
  newl := "\\\\ \n \\ \\\\ \n"
  if variant{
  exercise+="$"+mat2.Tex()+"$"
  solution+="Add "+(-mul).Tex() +" times row "+mathtypes.Number(ic+1).Tex()+" to row "+mathtypes.Number(i+1).Tex()+": " + newl
  solution+="$\\det("+mat2.Tex()+")=\\det("+mat3.Tex()+")$ "+newl
  solution+="Expand matrix along row "+mathtypes.Number(i+1).Tex()+":" +newl
  solution+="$\\det("+mat3.Tex()+")="+"(-1)^{"+mathtypes.Number(i+1).Tex()+"+"+mathtypes.Number(j+1).Tex()+"}"+m+p.Tex()+m+"\\det("+mat.Tex()+")$" +newl
  solution+="Calculate determinant of smaller matrix:" + newl
  solution+="$\\det("+mat.Tex()+")" +newl
  solution+=mat.DetTex()+"$"+newl
  solution+="Combine solutions:" +newl
  solution+="$\\det("+mat2.Tex()+")="+ mathtypes.Number(math.Pow(-1,float64(i+j))).Tex() +m+p.Tex()+m+mat.Det().Tex()+"="+mat2.Det().Tex()+"$"
  }else{
    exercise+="$"+mat2.Rotate().Tex()+"$"
    solution+="Add "+(-mul).Tex() +" times column "+mathtypes.Number(4-ic).Tex()+" to column "+mathtypes.Number(4-i).Tex()+": " + newl
    solution+="$\\det("+mat2.Rotate().Tex()+")=\\det("+mat3.Rotate().Tex()+")$ "+newl
    solution+="Expand matrix along row "+mathtypes.Number(4-i).Tex()+":" +newl
    solution+="$\\det("+mat3.Rotate().Tex()+")="+"(-1)^{"+mathtypes.Number(j+1).Tex()+"+"+mathtypes.Number(4-i).Tex()+"}"+m+p.Tex()+m+"\\det("+mat.Rotate().Tex()+")$" +newl
    solution+="Calculate determinant of smaller matrix:" + newl
    solution+="$\\det("+mat.Rotate().Tex()+")" +newl
    solution+=mat.Rotate().DetTex()+"$"+newl
    solution+="Combine solutions:" +newl
    solution+="$\\det("+mat2.Rotate().Tex()+")="+ mathtypes.Number(math.Pow(-1,float64(j-i+5))).Tex() +m+p.Tex()+m+mat.Rotate().Det().Tex()+"="+mat2.Rotate().Det().Tex()+"$"
  }

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
