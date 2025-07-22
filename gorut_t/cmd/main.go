package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("hhh")
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
