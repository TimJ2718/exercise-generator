package mathtypes
import (
	"strconv"
	//"fmt"
	"math"
)

type Number int
type Matrix [][]Number


func(mat Matrix) Tex() string{
	ret :="\\left(\\begin{array}{"
		for i:=0; i<len(mat[0]); i++ {
			ret+="r"
		}
		ret+="} \n"
		for n:=0; n< len(mat); n++{
			x :=""
			for m:=0; m<len(mat[n]) ; m++{
				x+= " "+strconv.Itoa(int(mat[n][m]))+ " &"
			}
			x = x[:len(x)-1]+"\\\\ \n"
			ret+=x
		}
		ret+="\\end{array}\\right)"
	return ret
	}

func RemoveRow(mat Matrix, k int)  Matrix {
	var ret Matrix
	if k >= len(mat)-1{
		ret := mat[:k]
		return ret
	}
	for i :=0; i < len(mat); i++{
		if i == k{
			i=i+1
		}
		ret = append(ret,mat[i])
	}
	return ret
}

func RemoveColumn(mat Matrix, k int)  Matrix {
	var ret Matrix
		for i:=0; i < len(mat); i++{
			var x []Number
			for j:=0; j <len(mat[i]); j++{
				if j ==k {
					j=j+1
					if j>=len(mat[i]){
						continue
					}
				}
				 x = append(x,mat[i][j])
			}
			ret = append(ret,x)
		}
		return ret
}


func(mat Matrix) Det() Number{
	if len(mat[0]) >2 {
		var ret Number = 0
		x := RemoveColumn(mat,0)
		for i :=0; i< len(mat);i++{
			x2:=RemoveRow(x,i)
			ret+=mat[i][0]*x2.Det()*Number(math.Pow(-1,float64(i)))
		}
		return ret
	}	else{
		return mat[0][0]*mat[1][1]-mat[0][1]*mat[1][0]
	}
}

func (num Number) Tex() string{
	if num<0{
		return "(-"+strconv.Itoa(int(-num))+")"
	} else {
		return strconv.Itoa(int(num))
	}
}
