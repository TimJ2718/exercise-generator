package gramschmidt
import (
  "fmt"
)

func Generate(name string,n,m int){
  if n <2{
    n =3
  }
  if m <2 || m >n{
    m =3
  }
  fmt.Printf("Still implementing")
  fmt.Printf("name: %v,n: %v, m: %v\n",name,n,m)
}
