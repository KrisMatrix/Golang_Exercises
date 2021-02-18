package main

import (
  "fmt"
  "time"
)

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)
  fmt.Println("Ok do this")
  time.Sleep(time.Second)
  fmt.Println("Ok done this")
  <-done
  fmt.Println("got here")
}
