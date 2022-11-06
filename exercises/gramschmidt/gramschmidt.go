package gramschmidt
import (
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
  "fmt"
)

//n : Dimension of Vector
//m : Number of Vectors
func Generate(name string,n,m int){   
  if n <2{
    n =3
  }
  if n >4{
    n=4
  }
  if m <2 || m >n{
    m =n
  }
  fmt.Printf("Higher numbers for n and m (>4) can take more time to generate. \n")
  max := mathtypes.Number(30) //Maximum size of numerator/denominator
  //m :=" \\cdot "
  newl := "\\\\ \n \\ \\\\ \n"
  exercise := ""
  solution := ""

  OUTER: //Repeat process of generating exercise until suitable exercise is generated
  for {  
    solution = "Use Gram-Schmidt process:" +newl
    exercise = "Orthogonalize the basis of vectors using the gram schmidt method :" +newl +"$"
    solution += "$u_i = v_i - \\sum\\limits_{j=1}^{i-1} \\frac{<u_j,v_i,>}{<u_j,uj>} \\cdot u_j $"+newl
    mp := " \\cdot "
  var mat mathtypes.Matrix
      //Create possible vectors for organilization
    mat = mathtypes.GetNXMMatrix(n,n,n)
    if mat.Det() == 0{ //Restart if linear dependent 
      continue
    }
    for i:=0;i<m;i++{  
      v:= mat[i].ToVectorF()
      if i+1<m{
        v2:= mat[i+1].ToVectorF()
        if mathtypes.IsPrime(mathtypes.DotproductF(v,v2).ToNumber()){ //Restart when dot products are prime numbers, because reducing is then not possible and numbers get large
          continue OUTER
        }
      }
    }
   
  
  var matF []mathtypes.VectorF //Vectors to orthogonalize
  var matF2 []mathtypes.VectorF //Orthogonalized vectors
  for i:=0; i< m; i++{
    matF = append(matF,mat[i].ToVectorF())
    exercise+="v_{"+mathtypes.Number(i+1).Tex()+"}="+matF[i].Tex()+", \\;\\;\\;\\; \n"
  }

  exercise+="$"
  
  for i:=0 ; i<len(matF) ;i++{
    solution += "Vector "+mathtypes.Number(i+1).Tex()+": "+newl
    //Write formula for vector i:  
    solution +="$u_{"+mathtypes.Number(i+1).Tex()+"}=v_{"+mathtypes.Number(i+1).Tex()+"}"
    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{<u_{"+mathtypes.Number(j+1).Tex()+"},v_{"+mathtypes.Number(i+1).Tex()+"}>}{<u_{"+mathtypes.Number(j+1).Tex()+"},u_{"+mathtypes.Number(j+1).Tex()+"}>}"+mp+"u_{"+mathtypes.Number(j+1).Tex()+"}"
    }
    //Put in numbers and already calculate dot products
    solution += newl + "=" + matF[i].Tex()
    for j:=0; j < i ; j++{
      solution += "-"+"\\frac{"+mathtypes.DotproductF(matF2[j],matF[i]).Tex() +"}"+"{"+mathtypes.DotproductF(matF2[j],matF2[j]).Tex()+"}"+mp+matF2[j].Tex()
    }
    if i>0{
      solution+=newl+"="+matF[i].Tex()
      }
    //Multiply dot products with related vectors
    var v []mathtypes.VectorF
    v = append(v,matF[i])
    for j:=0; j<i ; j++{
      s1:=mathtypes.DotproductF(matF2[j],matF[i])
      s2:=mathtypes.DotproductF(matF2[j],matF2[j])
      s3:=mathtypes.DivideF(s1,s2)
      v = append(v,mathtypes.ScalarVectorF(s3,matF2[j])) 
      for k:=0; k < len(v[j+1]); k++{ //Check for every vector that the numbers are not to large
        if v[j+1][k].Getn().Abs() > max || v[j+1][k].Getd().Abs() > max{
          continue OUTER
        }
    }
      solution+="-"+v[j+1].Tex()
    }
    if i>0{
      solution+=newl
    }
    //Subtract vectors
    for j:=1;j <len(v);j++{
        v[0] = mathtypes.SubtractVectorF(v[0],v[j])
          for k:=0; k < len(v[0]); k++{
            if v[0][k].Getn().Abs() > max || v[0][k].Getd().Abs() > max{ //Check for every vector that the numbers are not to large
              continue OUTER
          }
          }
        solution+="="+v[0].Tex()
        for k:=j; k < len(v)-1;k++{
          solution+="-"+v[k+1].Tex()+newl
        }
    

    }
    //Apend orthogonalized vectors
    matF2 = append(matF2,v[0])
    solution += "$"+newl
  }


  textofile.SaveTofile(name,exercise, solution)
  break
}
}