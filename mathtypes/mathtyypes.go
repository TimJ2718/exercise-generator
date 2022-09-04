package mathtypes
import (
	"strconv"
	"math/rand"
  "math"
  "time"
)


//Calculations with integer
type Number int
type Vector []Number
type Matrix []Vector


func(mat Matrix) Tex() string{
	ret :="\\left(\\begin{array}{"
		for i:=0; i<len(mat[0]); i++ {
			ret+="c"
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
func GetNXMMatrix(n,m,r int) Matrix{
		var mat Matrix
		for i :=0 ; i<n ; i++{
				var x Vector
	      for j := 0; j <m; j++{
	        x= append(x,Number(RandomIntExcept(-r,r,[]int{0})))
	      }
				mat = append(mat,x)
	  }
	  return mat
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

func(mat Matrix) Rotate() Matrix{
	var ret Matrix
	for i:=0; i<len(mat[0]);i++{
		var x Vector
		for j:=0; j<len(mat);j++{
			x = append(x, mat[len(mat)-j-1][i])
		}
		ret = append(ret, x)
	}
	return ret
}


func GetVector(n int) Vector{
	var ret Vector
	for i :=0 ; i<n; i++{
		ret = append(ret,Number(RandomIntExcept(-9,9,[]int{0})))
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

//Calculations with fractions

type Fraction struct {
	n Number //numerator
	d Number //denominator
}

type VectorF []Fraction

func (frac Fraction) Reduce() Fraction{
	var min int
	sign := Number(1)
	if frac.n ==0{
		return Fraction{0,1}
	}
	if frac.n <0{
		frac.n = -frac.n
		sign = -sign
	}
	if frac.d <0{
		frac.d = -frac.d
		sign = -sign
	}
	if frac.n < frac.d{
			min =int(frac.n)
		} else{
			min =int(frac.d)
		}
	for i:=1 ; i<min; i++{
		if frac.n % Number(min-i) ==0 && frac.d % Number(min-i) ==0{
			return Fraction{Number(float64(sign*frac.n)/float64(min-i)),Number(float64(frac.d)/float64(min-i))}
		}
	}
	return Fraction{sign*frac.n,frac.d}
}

func (num Number) ToFraction() Fraction{
	return Fraction{num,Number(1)}
}

func (frac Fraction) ToNumber() Number{
	return frac.n
}

func (vec Vector) ToVectorF() VectorF{
	var ret VectorF
	for i:=0;i<len(vec);i++{
		ret = append(ret,vec[i].ToFraction())
	}
	return ret
	}

func (frac Fraction) Tex() string{
	if frac.d == Number(1){
		return frac.n.Tex()
	}
	return "\\frac{"+frac.n.Tex()+"}{"+frac.d.Tex()+"}"
}

func (vec VectorF) Tex() string{
	ret := "\\left(\\begin{array}{c} \n "
	for i:= 0; i <len(vec); i++{
		ret += vec[i].Tex() +"\\\\ \n"
	}
	ret +="\\end{array}\\right)"
	return ret
}


func Dotproduct(v1, v2 VectorF) Fraction{
	ret := Fraction{0,1}
	for i:=0 ;i < len(v1);i++{
		ret = AddF(ret,MultiplyF(v1[i],v2[i]))
	}
	return ret
}
func MultiplyF(f1, f2 Fraction) Fraction{
	n:= f1.n*f2.n
	d:= f1.d*f2.d
	return Fraction{n,d}.Reduce()
}
func DivideF(f1, f2 Fraction) Fraction{
	n:= f1.n*f2.d
	d:= f1.d*f2.n
	return Fraction{n,d}.Reduce()
}


func ScalarVectorF(s Fraction, v VectorF) VectorF{
	var ret VectorF
	for i:=0 ; i< len(v); i++{
		ret = append(ret,MultiplyF(s,v[i]))
	}
	return ret
}

func AddVectorF(v1, v2 VectorF) VectorF{
	var ret VectorF
	for i:=0 ;i <len(v1); i++{
		ret = append(ret,AddF(v1[i],v2[i]))
	}
	return ret
}
func SubtractVectorF(v1, v2 VectorF) VectorF{
	var ret VectorF
	for i:=0 ;i <len(v1); i++{
		ret = append(ret,SubtractF(v1[i],v2[i]))
	}
	return ret
}

func SubtractF(f1, f2 Fraction) Fraction{
	n:= f1.n*f2.d-f2.n*f1.d
	d:= f1.d * f2.d
	return Fraction{n,d}.Reduce()
}

func IsPrime(num Number) bool {
	x := int(num)
	for i:=2; i<x; i++{
		if x%i ==0{
			return false
		}
	}
	return true
}

func AddF(f1, f2 Fraction) Fraction{
	n:= f1.n*f2.d+f2.n*f1.d
	d:= f1.d * f2.d
	return Fraction{n,d}.Reduce()
}
