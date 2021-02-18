package main

import "fmt"

func zeroval(ival int) {
    ival = 0
    fmt.Println("ival = ",ival)
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)  //i = 1

    zeroval(i)
    fmt.Println("zeroval:", i)  //i = 1. ival is a local var and still 0

    zeroptr(&i)
    fmt.Println("zeroptr:", i)  //i = 0. We are passing the address of i
                                // and the address is dereferenced within
                                // the function zeroptr and set to 0

    fmt.Println("pointer:", &i) //address of i itself.
}
