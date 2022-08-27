package mathtypes
import (
	"strconv"
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

func(mat Matrix) Det() Number{
	return mat[0][0]*mat[1][1]-mat[0][1]*mat[1][0]
}

func (num Number) Tex() string{
	if num<0{
		return "(-"+strconv.Itoa(int(-num))+")"
	} else {
		return strconv.Itoa(int(num))
	}
}
