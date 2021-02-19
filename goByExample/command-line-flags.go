package main

import (
  "flag"
  "fmt"
)

func main() {
  //Flags:
  // flag1 := flag.DataType(flagName, flagDefault, flagDescription)
  //  -flagName DataType
  //    flagDescription (default flagDefault)
  //
  // wordPtr := flag.String("word", "foo", "a string")
  //
  //-word string      
  //  	a string (default "foo")

  //Setup of flags
  wordPtr := flag.String("word", "foo", "a string")
  numbPtr := flag.Int("numb", 42, "an int")
  boolPtr := flag.Bool("fork", false, "an bool")

  var svar string
  flag.StringVar(&svar, "svar", "bar", "a string var")
  flag.Parse()  //parse the flags provided

  //dereference with * to get the value for the flags

  fmt.Println("word:", *wordPtr)  
  fmt.Println("numb:", *numbPtr)
  fmt.Println("bool:", *boolPtr)
  fmt.Println("svar:", svar)
  fmt.Println("tail:", flag.Args())
}
