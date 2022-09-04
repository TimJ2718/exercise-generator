package gramschmidt
import (
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
  solution += "$u_i = v_i - \\sum\\limits_{j=1}^{i-1} \\frac{<u_j,v_i,>}{<u_j,uj>} \\cdot u_j $"+newl
  mp := " \\cdot "
  var mat mathtypes.Matrix
  OUTER:
  for {
    mat = mathtypes.GetNXMMatrix(n,n,n)
    if mat.Det() == 0{
      continue
    }
    for i:=0;i<m;i++{
      v:= mat[i].ToVectorF()
      if mathtypes.IsPrime(mathtypes.Dotproduct(v,v).ToNumber()){
        continue OUTER
      }
      if i+1<m{
        v2:= mat[i+1].ToVectorF()
        if mathtypes.IsPrime(mathtypes.Dotproduct(v,v2).ToNumber()){
          continue OUTER
        }
      }
    }
    break
  }

  var matF []mathtypes.VectorF //Vectors to orthogonalize
  var matF2 []mathtypes.VectorF //Orthogonalized vectors
  for i:=0; i< m; i++{
    matF = append(matF,mat[i].ToVectorF())
    exercise+="v_{"+mathtypes.Number(i+1).Tex()+"}="+matF[i].Tex()+", \\;\\;\\;\\; \n"
  }

  //checke größe von nenner und zähler und regenerate entsprechend
  exercise+="$"
  for i:=0 ; i<len(matF) ;i++{
    solution += "Vector "+mathtypes.Number(i+1).Tex()+": "+newl
    solution +="$u_{"+mathtypes.Number(i+1).Tex()+"}=v_{"+mathtypes.Number(i+1).Tex()+"}"
    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{<u_{"+mathtypes.Number(j+1).Tex()+"},v_{"+mathtypes.Number(i+1).Tex()+"}>}{<u_{"+mathtypes.Number(j+1).Tex()+"},u_{"+mathtypes.Number(j+1).Tex()+"}>}"+mp+"u_{"+mathtypes.Number(j+1).Tex()+"}"
    }
    solution += newl + "=" + matF[i].Tex()
    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{"+mathtypes.Dotproduct(matF2[j],matF[i]).Tex() +"}"+"{"+mathtypes.Dotproduct(matF2[j],matF2[j]).Tex()+"}"+mp+matF2[j].Tex()
    }
    if i>0{
      solution+=newl+"="+matF[i].Tex()
      }
    var v []mathtypes.VectorF
    v = append(v,matF[i])
    for j:=0; j<i ; j++{
      s1:=mathtypes.Dotproduct(matF2[j],matF[i])
      s2:=mathtypes.Dotproduct(matF2[j],matF2[j])
      s3:=mathtypes.DivideF(s1,s2)
      v = append(v,mathtypes.ScalarVectorF(s3,matF2[j]))
      solution+="-"+v[j+1].Tex()
    }
    if i>0{
      solution+=newl
    }
    for j:=1;j <len(v);j++{
        v[0] = mathtypes.SubtractVectorF(v[0],v[j])
        solution+=v[0].Tex()
        for k:=j; k < len(v)-1;k++{
          solution+="-"+v[k+1].Tex()+newl
        }
    }
    matF2 = append(matF2,v[0])
    solution += "$"+newl
  }


  textofile.SaveTofile(name,exercise, solution)
}
