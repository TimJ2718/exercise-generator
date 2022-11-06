package crossproduct
import (
  "exercise-generator/mathtypes"
  "exercise-generator/textofile"
)


func Generate(name string){
    exercise :="Calculate a vector orthogonal to the two vectors:  \\\\ \n\\ \\\\ \n"
    solution := "Calculate cross product to determine an orthogonal vector: \\\\ \n\\ \\\\ \n"
    var v1, v2 mathtypes.Vector
    for{
        v1 = mathtypes.GetVector(3)
        v2 = mathtypes.GetVector(3)
        if mathtypes.Dotproduct(v1,v2) !=0{
            break
        }
    }
    exercise+="$ v_1="
    exercise += v1.Tex()
    exercise+="\\; \\;  v_2= "
    exercise+= v2.Tex()  
    exercise+="$ \\\\ \n\\ \\\\ \n"
    v3 := mathtypes.Crossproduct(v1,v2)
    solution+="$v_1 \\times v_2 = " + v1.Tex() + " \\times " + v2.Tex() +"="+v3.Tex()+"$ \\\\ \n\\ \\\\ \n"
    var i int
    for{
        i = mathtypes.RandomInt(0,2)
        if v3[i] !=0{
            break
        }   
    }
    m := mathtypes.RandomIntExcept(-2,2, []int{0,1}) 
    exercise+="The "  +mathtypes.Number(i+1).Tex() +". element should be:  $v_"+mathtypes.Number(i+1).Tex()+"="+(mathtypes.Number(m)*v3[i]).Tex() +"$"
    solution+="Multiply with: " +mathtypes.Number(m).Tex() + " to get $v_"+mathtypes.Number(i+1).Tex()+"="+(mathtypes.Number(m)*v3[i]).Tex() +"$: \\\\ \n\\ \\\\ \n"
    solution+="$" + mathtypes.Number(m).Tex() + "\\cdot" + v3.Tex() +"=" + mathtypes.ScalarMultiplication(mathtypes.Number(m),v3).Tex() + "$"
    textofile.SaveTofile(name,exercise, solution)
}  