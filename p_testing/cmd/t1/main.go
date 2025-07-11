package main

import "fmt"

func main() {
	s := Sum(2, 3)
	fmt.Println(s)
}

func Sum(a, b int) int {
	return a + b
}
