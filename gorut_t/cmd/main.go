package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// func echo(in io.Reader, out io.Writer) {
// 	io.Copy(out, in)
// }

func main() {
	// go echo(os.Stdin, os.Stdout)
	// time.Sleep(10 * time.Second)
	// fmt.Println("timeout")
	// os.Exit(0)

	// fmt.Println("outside a gorunine")
	// go func() {
	// 	fmt.Println("inside a gorunine")
	// }()
	// fmt.Println("outside again")
	// runtime.Gosched()

	// for _, file := range os.Args[1:] {
	// 	compress(file)
	// }

	// s := make([]int, 3, 4)
	// s[0] = 2
	// s[2] = 56
	// fmt.Println(s)
	// s2 := s[1:3]
	// fmt.Println(s2)
	// // s2[3] = 4 err
	// s2 = append(s2, 4)
	// fmt.Println(s2)
	// fmt.Println(s)
	// s = append(s, 101)
	// fmt.Println(s2)
	// fmt.Println(s)

	// fmt.Printf("pointer: %p ; pointer2: %p \n", s, s2)
	// fmt.Printf("s: %d ; s2: %d \n", len(s), len(s2))

	// s2 = append(s2, 981)
	// fmt.Printf("pointer: %p ; pointer2: %p \n", s, s2)
	// s2[0] = -25
	// fmt.Println(s2)
	// fmt.Println(s)

}

func compress(filename string) {
	in, err := os.Open(filename)
	if err != nil {
		log.Fatal("err open file")
	}
	defer in.Close()

	s := strings.Split(filename, ".")
	for _, ss := range s {
		fmt.Printf("%v ", ss)
	}
	fmt.Println()

	out, err := os.Create(s[0] + ".gz")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer out.Close()

	// bufio.NewWriter()

	gzout := gzip.NewWriter(out)
	i, err := io.Copy(gzout, in)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(i)
}


