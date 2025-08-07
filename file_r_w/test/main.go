package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	simleReader()
}

func simleReader() {
	buf := make([]byte, 10)
	// var buf []byte

	data := SimpleStruct{nums: "12,34,5.67&b=8"}

	count, err := data.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(buf[8])
	fmt.Println(string(buf), count)

}

type SimpleStruct struct {
	nums string
}

func (s SimpleStruct) Read(p []byte) (n int, err error) {
	var count int

	for i := 0; i < len(s.nums); i++ {
		if s.nums[i] <= '9' && s.nums[i] >= '0' {
			p[count] = s.nums[i]
			count++
		}
	}
	return count, io.EOF
}
