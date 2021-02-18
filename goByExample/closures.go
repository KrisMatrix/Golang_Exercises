package main

import "fmt"

func intSeq() func() int {  //the func intSeq() returns another function
                            // which in turn returns an int value
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  nextInt := intSeq()       //we call intSeq() assigning the result
                            // (which is the anonymous func that 
                            // incerments i) to nextInt. So i is
                            // initialized to 0, and every time we 
                            // call nextInt(), we are calling the
                            // anonymous function, which increments i

  fmt.Println(nextInt())    //1
  fmt.Println(nextInt())    //2
  fmt.Println(nextInt())    //3

  newInts := intSeq()       
  fmt.Println(newInts())    //1
}
