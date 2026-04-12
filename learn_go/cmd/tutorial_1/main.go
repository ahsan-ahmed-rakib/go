package main 

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum) // Output: -32768

	var floatNum  float64 = 12345678.9
	fmt.Println(floatNum) // Output: 12345678.9

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result) // Output: 12.1

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // Output 1
	fmt.Println(intNum1 % intNum2) // Output 1
	fmt.Println(float64(intNum1) / float64(intNum2)) // Output 1.5

	var myString string = "Hello, World!"
	fmt.Println(myString) // Output: Hello, World!

	var str1 string = `Hello,
	Line 2,`
	fmt.Println(str1)

	fmt.Println(len("Rakib")) // Output: 5
	fmt.Println(utf8.RuneCountInString("Y")) // Output: 1

	var myBoolean bool = true
	fmt.Println(myBoolean) // OP: true
	
	myVar, myVar2 := "Var shortcut", 2
	fmt.Println(myVar, myVar2)

	const pi float64 = 3.1416
	fmt.Println(pi)
}