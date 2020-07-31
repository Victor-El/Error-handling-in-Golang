package main

import (
  "fmt"
  "errors"
)

func main() {
  fmt.Println(divide(20, 5))
  val, err := divide(40, 5)
  fmt.Println(err, val)
  fmt.Println(divide(20, 0))
  fmt.Println(divide2(100, 2))
  fmt.Println(divide2(100, 0))
  fmt.Println(divide2(100, 5))
}

func divide(dividend int, divisor int) (quotient int, err error) {
  quotient = 0
  err = nil
  if divisor == 0 {
    err = errors.New("Division by zero error")
    return
  }
  quotient = dividend / divisor
  return
}

func divide2(dividend int, divisor int) (int){
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("recovered from panic")
    }
  }()

  res := dividend / divisor

  return res

}