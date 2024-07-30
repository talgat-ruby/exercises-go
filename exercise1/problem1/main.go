package main

func addUp(x int) int {
  var sum int = 0
  for i := 1; i <= x; i++{
    sum += i
  }
  return sum
}
