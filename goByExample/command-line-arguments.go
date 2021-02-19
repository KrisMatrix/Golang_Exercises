package main

import (
  "fmt"
  "os"
)

func main() {
  argsWithProg := os.Args         //includes the calling script and all args
  argsWithoutProg := os.Args[1:]  //excludes the calling scripts but has all 
                                  // other args

  arg := os.Args[3]
  
  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)
  fmt.Println(arg)
}
