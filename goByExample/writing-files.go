package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  d1 := []byte("hello\ngo\n")
  err := ioutil.WriteFile("/tmp/dat1", d1, 0644)  //dump string into file
  check(err)

  f, err := os.Create("/tmp/dat2")  //for more granular writes, open file for 
                                    // writing
  check(err)

  defer f.Close()

  d2 := []byte{115, 111, 109, 101, 10}  //this is ascii values. s o m e \n
  n2, err := f.Write(d2)
  check(err)
  fmt.Printf("n2 write %d bytes\n", n2)

  n3,err := f.WriteString("writes\n")
  check(err)
  fmt.Printf("n3 wrote %d bytes\n", n3)

  f.Sync()

  w := bufio.NewWriter(f)
  n4, err := w.WriteString("buffered\n")
  check(err)
  fmt.Printf("n4 wrote %d bytes\n", n4)

  w.Flush()
}
