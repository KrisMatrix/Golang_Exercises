package main

import (
  "bufio"   //NewReader
  "fmt"     //Println, Printf
  "io"      //ReadAtLeast
  "io/ioutil" //ReadFile
  "os"      //open and close file
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  dat, err := ioutil.ReadFile("addresses.csv")  //reads entire file
  check(err)
  fmt.Print(string(dat))

  fmt.Println("---DONE---")

  f, err := os.Open("addresses.csv")
  check(err)
  fmt.Println("f = ",f)     //f seems to be a memory address

  b1 := make([]byte, 5)
  n1, err := f.Read(b1)     //the file is read and place in slice b1 for 5 bytes
                            // n1 will be 5 which is the no. of bytes read.
  fmt.Println("n1 = ",n1)
  check(err)
  fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

  o2, err := f.Seek(6,0)
  check(err)
  b2 := make([]byte, 2)
  n2, err := f.Read(b2)
  check(err)
  fmt.Printf("%d bytes @ %d: ", n2, o2)
  fmt.Printf("%v\n", string(b2[:n2]))

  o3, err := f.Seek(6, 0)
  check(err)
  b3 := make([]byte, 2)
  n3, err := io.ReadAtLeast(f, b3, 2)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

  _, err = f.Seek(0, 0)
  check(err)

  r4 := bufio.NewReader(f)
  b4, err := r4.Peek(5)
  check(err)
  fmt.Printf("5 bytes: %s\n", string(b4))
  f.Close()
}
