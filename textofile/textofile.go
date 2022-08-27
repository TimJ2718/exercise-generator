package textofile

import (
  "fmt"
  "io/ioutil"
)

func getdocument(text, titel string) string{
  ret :="\\documentclass[12pt,a4paper]{article} \n"
  ret +="\\usepackage[utf8]{inputenc} \n"
  ret+="\\begin{document}\n"
  ret+="\\begin{center} \n"+titel+"\n \\end{center} \n"
  ret+=text
  ret+="\n \\end{document}"
  return ret
}

func SaveTofile(name,exercise, solution string){
  exercise = getdocument(exercise,"Exercise")
  exdata :=[]byte(exercise)
  err := ioutil.WriteFile(name+"-ex.tex",exdata,0600)
  if err!=nil{
    fmt.Printf("%v \n",err)
  }
  solution = getdocument(solution,"Solution")
  soldata :=[]byte(solution)
  err = ioutil.WriteFile(name+"-sol.tex",soldata,0600)
  if err!=nil{
    fmt.Printf("%v \n",err)
  }
}
