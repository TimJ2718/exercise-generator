package mathtypes
import (
	"strconv"
	//"fmt"
	"math/rand"
  "math"
  "time"
)

type Number int
type Vector []Number
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
func (mat Matrix) DetTex() string{
	if len(mat)==3 && len(mat[0])==3{
		m := " \\cdot "
		ret :="="+mat[0][0].Tex()+m+mat[1][1].Tex()+m+mat[2][2].Tex()+"+"+mat[0][1].Tex()+m+mat[1][2].Tex()+m+mat[2][0].Tex()+"+"+mat[0][2].Tex()+m+mat[1][0].Tex()+m+mat[2][1].Tex()
  	ret +="-"+mat[2][0].Tex()+m+mat[1][1].Tex()+m+mat[0][2].Tex()+"-"+mat[2][1].Tex()+m+mat[1][2].Tex() +m+mat[0][1].Tex()+"-"+mat[2][2].Tex()+m+mat[1][0].Tex()+m+mat[0][1].Tex()+"\n \\\\"
  	ret +="="+(mat[0][0]*mat[1][1]*mat[2][2]).Tex()+"+"+(mat[0][1]*mat[1][2]*mat[2][0]).Tex()+"+"+(mat[0][2]+mat[1][0]+mat[2][1]).Tex()
  	ret +="-"+(mat[2][0]*mat[1][1]*mat[0][2]).Tex()+"-"+(mat[2][1]*mat[1][2]*mat[0][0]).Tex()+"-"+(mat[2][0]*mat[1][0]*mat[0][1]).Tex()+"\n \\\\"
  	ret +="="+mat.Det().Tex()
			return ret
	}
	return "This is not implemented!"
}

func(mat Matrix) Transpose() Matrix{
	var ret Matrix
	for i:=0; i<len(mat); i++{
		var zw []Number
		for j:=0; j<len(mat[i]);j++{
			zw = append(zw,mat[j][i])
		}
		ret = append(ret,zw)
	}
	return ret
}

	//create a Matrix where the abosolute value of the product of each diagonal line is smaller 150
	func Get3Matrix() Matrix{
	  var max float64 = 150
	  x := Matrix{{0,0,0},{0,0,0},{0,0,0}}
	  for {
	    for i :=0 ; i<3 ; i++{
	      for {
	        for j := 0; j <3; j++{
	          x[i][j]= Random()
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
	  return Matrix{{x[0][0],x[1][0],x[2][0]},{x[2][1],x[0][1],x[1][1]},{x[1][2],x[2][2],x[0][2]}}
	}

func AddVectorToRow(mat Matrix, vec Vector, row int) Matrix{
	var ret Matrix
	for i:=0; i<len(mat); i++{
		var x []Number
		for j:=0; j<len(mat[0]); j++{
			if i==row{
				x=append(x,(mat[i][j]+vec[j]))
			} else{
					x=append(x,mat[i][j])
			}
		}
		ret=append(ret,x)
	}
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

func AddRow(mat Matrix, vec Vector, k int) Matrix {
	var ret Matrix
	for i:=0; i <len(mat); i++{
			if i==k{
				ret = append(ret,vec)
			}
			ret = append(ret,mat[i])
	}
	if k >=len(mat){
		ret = append(ret,vec)
	}
	return ret
}

func AddColumn(mat Matrix, vec Vector, k int) Matrix {
	var ret Matrix
	for i:=0; i <len(mat); i++{
		var zw []Number
		for j:=0; j < len(mat[i]); j++{
			if j==k{
				zw = append(zw,vec[i])
			}
			zw = append(zw,mat[i][j])
		}
		if k >=len(mat){
			zw = append(zw,vec[i])
		}
		ret = append(ret,zw)
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




func GetVector(n int) Vector{
	var ret Vector
	for i :=0 ; i<n;i++{
		ret = append(ret,Random())
		}
	return ret
}

func AddElement(vec Vector,value Number, pos int) Vector{
	var ret Vector
	for i:=0;i<len(vec);i++{
		if i==pos{
			ret = append(ret,value)
		}
		ret = append(ret,vec[i])
	}
	if pos >=len(vec){
		ret = append(ret,value)
	}
	return ret
}

func ScalarMultiplication(m Number, vec Vector) Vector{
	var ret Vector
	for i:=0;i<len(vec);i++{
		ret = append(ret,m*vec[i])
	}
	return ret
}

func Random() Number{
	rand.Seed(time.Now().UnixNano())
	return Number(rand.Intn(19)-9)
}

func RandomInt(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1)+min
}

func RandomIntExcept(min,max int, blacklist []int) int{
	rand.Seed(time.Now().UnixNano())
		OUTER:
		for {
			ret := (rand.Intn(max-min+1)+min)
			for i:=0; i<len(blacklist);i++{
				if ret == blacklist[i]{
					continue OUTER
				}
			return ret
			}
		}
	}

func (num Number) Tex() string{
	if num<0{
		return "(-"+strconv.Itoa(int(-num))+")"
	} else {
		return strconv.Itoa(int(num))
	}
}
