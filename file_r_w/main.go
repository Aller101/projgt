package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func cheTest(st1 string, st2 string) string {
	slice1 := strings.Split(st1, "")
	slice2 := strings.Split(st2, "")
	if len(st1) == len(st2) {
		for i, _ := range slice1 {
			if slice1[i] != slice2[i] {
				return fmt.Sprintf("element: %v != elemrnt %v\n", slice1[i], slice2[i])
			}
		}
	}
	return "good"

}

func main() {

	var s_my = "massage"
	var tuzov = "message"

	fmt.Println(cheTest(s_my, tuzov))

	fmt.Println(s_my == tuzov)

	// var list1 []int
	// list2 := []int{}

	// fmt.Println(list1, " ", list2)

	// // list := []int{0, 58, -12, 99}
	// listF := Append(list2, 3)
	// fmt.Println(listF)

	// file, err := os.Open("cats.jpg")
	// if err != nil {
	// 	fmt.Printf("Err read file: %s", err)
	// 	return
	// }

	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner)
	// }

}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)

	b = digitRegexp.Find(b)

	res := make([]byte, len(b))

	copy(res, b)

	return res
}

func Append(list []int, elem int) []int {
	var res []int
	resList := len(list) + 1
	if resList <= cap(list) {
		res = list[:resList]
	} else {
		resCap := resList
		if resCap < len(list)*2 {
			resCap = len(list) * 2
		}
		res = make([]int, resList, resCap)
		copy(res, list)
	}

	res[len(list)] = elem

	return res
}
