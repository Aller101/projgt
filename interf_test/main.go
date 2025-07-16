package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Runner interface {
	Run() string
}

func (h *Human) Run() string {
	return h.name
}

func main() {

	z := 42
	fmt.Println(&z)
	fmt.Println(*&z)

	fmt.Println()
	fmt.Println("---------------------")
	var runner Runner
	fmt.Printf("%T: %v\n", runner, runner)
	if runner == nil {
		fmt.Println("r1 == nil")
	}
	fmt.Println("---------------------")
	fmt.Println()

	fmt.Println("var unnamedRunner *Human")
	var unnamedRunner *Human
	prntHuman(unnamedRunner)
	fmt.Println()
	var newPoiner = new(Human)
	prntHuman(newPoiner)
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("runner = unnamedRunner")
	runner = unnamedRunner
	prntHuman(runner)
	fmt.Println("---------------------")
	fmt.Println()
	fmt.Printf("type - %T: value - %v , pointer - %v\n", runner, runner, &runner)
	if runner == nil {
		fmt.Println("r1 == nil")
	}
	fmt.Println("---------------------")

	runner = &Human{}
	fmt.Printf("type - %T: value - %v , pointer - %v\n", runner, runner, &runner)
	if runner == nil {
		fmt.Println("r1 == nil")
	}

	fmt.Println("---------------------")

}

func prntHuman[T comparable](value T) {
	fmt.Printf("%T: %v \n", value, value)
}
