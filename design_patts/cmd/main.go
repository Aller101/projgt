package main

import (
	p "design_patts/singleton"
	"fmt"
)

func main() {
	// s := p.Singleton{}

	// tests := [10]struct {
	// 	name string
	// 	id   int
	// }{}
	// fmt.Println(tests)
	// fmt.Println(len(tests))

	// for i := range 10 {

	// 	tests[i].name = fmt.Sprintf("case: %d", i)
	// 	tests[i].id = i
	// }

	s := p.NewSingleton("Kolia")
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s.GetName())

	s2 := p.NewSingleton("Nastia")
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s2.GetName())

	fmt.Printf("%t: %v --- %v\n", *s2 == *s, *s2, *s)
	fmt.Printf("%t: %v --- %v\n\n", &s2 == &s, &s2, &s)

	s = p.NewSingleton("Vasia")

	fmt.Printf("%v\n", s.GetName())

}
