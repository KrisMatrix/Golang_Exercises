package main

import "fmt"

func main() {
  i := 1

  for i<=3 {  //step through i from 1 to 3
    fmt.Println(i)
    i = i + 1
  }

  for j:= 7; j <= 9; j++ {  //step from 7 through 9
    fmt.Println(j)
  }

  for { //print loop once
    fmt.Println("loop")
    break
  }

  for n:=0; n <= 5; n++ { //print odd numbers
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}
