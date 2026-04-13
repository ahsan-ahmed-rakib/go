package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myString = "rèsumè"
	var myString = []rune("rèsumè")
	fmt.Println(myString)

	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed) // 114, uint8
	for i, v := range myString{
		fmt.Println(i, v)
	}
	// 	0 114  rune 0
	// 	1 232		1
	// 	3 115		2
	// 	4 117		3
	// 	5 109		4
	// 	6 232		5

	fmt.Printf("\nThe length of 'myString' is %v", len(myString)) // 8

	var strSlice = []string{"S", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder // Build string
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr) // Subscribe
}