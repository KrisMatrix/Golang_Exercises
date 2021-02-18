package main

import "fmt"

func main() {
  nums := []int{2,3,4}
  sum := 0
  for _, num := range nums {  //range on arrays and slices provides both
                              // index and value for each entry. Since we
                              // don't need the index, we use '_' to ignore
    sum += num
  }

  fmt.Println("sum:", sum)    //9

  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)  // 1
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  
  for k, v := range kvs {
    fmt.Println("%s -> %s\n", k, v) //a -> apple
                                    //b -> banana
  }         

  for k := range kvs {
    fmt.Println("key:", k)          //key: a
                                    //key: b
  }

  for i, c := range "go" {
    fmt.Println(i, c)           //0 103 //note 103 is ascii 'g'
                                //1 111 //note 111 is ascii 'o'
  }
}
