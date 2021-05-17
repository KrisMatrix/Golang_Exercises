package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type response1 struct {
  Page int
  Fruits []string
}

type response2 struct {
  Page int
  Fruits []string
}

func main() {
  bolB, _ = json.Marshal(true)
  fmt.Println(string(bolB))

  intB, _ = json.Marshal(1)
  fmt.Println(string(intB))

  fltB, _ = json.Marshal(2.34)
  fmt.Println(string(fltB))

  strB, _ = json.Marshal("gopher"
}
