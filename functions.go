package main
import "fmt"

func add(x int, y int) int {
  return x + y
}

func add2(x, y int) int {
  return x + y
}

func swap(x,y string) (string, string) {
  return y,x
}

func main() {
  fmt.Println(add(42,13))
  fmt.Println(add2(2,53))

  a, b := swap("hello", "world")
  fmt.Println(a,b)
}
