package main

import (
  "fmt"
  "log"
  "io/ioutil"
)

func main() {
  content, err := ioutil.ReadFile("addresses.csv");
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("File contents: %s", content)
}
