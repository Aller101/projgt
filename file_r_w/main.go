package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("cats.jpg")
	if err != nil {
		fmt.Printf("Err read file: %s", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner)
	}

}
