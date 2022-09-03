package gramschmidt
import (
  "fmt"
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
)

//n : Dimension of Vector
//m : Number of Vectors
func Generate(name string,n,m int){
  if n <2{
    n =3
  }
  if m <2 || m >n{
    m =n
  }
  //m :=" \\cdot "
  newl := "\\\\ \n \\ \\\\ \n"
  exercise := "Orthogonalize the basis of vectors using the gram schmidt method :" +newl +"$"
  solution := "Use Gram-Schmidt process:" +newl
  solution += "$u_j = v_j - \\sum\\limits_{i=1}^{j-1} \\frac{<u_i,v_j,>}{<u_i,ui>} \\cdot u_i $"+newl
  var mat mathtypes.Matrix
  for {
    mat = mathtypes.GetNXMMatrix(n,n)
    if mat.Det() != 0{
      break
    }
  }

  var matF []mathtypes.VectorF
  for i:=0; i< m; i++{
    matF = append(matF,mat[i].ToVectorF())
    exercise+="v_{"+mathtypes.Number(i+1).Tex()+"}="+matF[i].Tex()+", \\;\\;\\;\\; \n"
  }
  for i:=0 ; i<len(matF) ;i++{
    solution += "Vector 1:"+newl
    solution +="$u_{"+mathtypes.Number(i+1).Tex()+"}=v_{"+mathtypes.Number(i+1).Tex()+"}"
    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{<u_{"+mathtypes.Number(j+1).Tex()+"}v_{"+mathtypes.Number(i+1).Tex()+"}>}{2}"
    }
    solution += newl + "=" + matF[i].Tex()

    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{1}{2}"
    }
    solution += "$"+newl
  }

  fmt.Printf("Mat:= %v \n",mat.Tex())
  fmt.Printf("Still implementing\n")
  fmt.Printf("name: %v,n: %v, m: %v\n",name,n,m)
  fmt.Printf(exercise)
  textofile.SaveTofile(name,exercise, solution)
}
