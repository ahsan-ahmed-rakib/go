package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 10
	fmt.Println(intArr[0]) // 10
	fmt.Println(intArr[1:3]) // 0 0

	fmt.Println(&intArr[0]) // 0x21da3f876160

	var intArr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr1) // [1, 2, 3]
	
	intArr2 := []int32{4, 5, 6}
	fmt.Println(intArr2) // [4, 5, 6]
	
	intArr2 = append(intArr2, 7)
	fmt.Println(intArr2) // [4, 5, 6, 7]
	
	var intArr3 []int32 = []int32{8, 9, 10}
	intArr2 = append(intArr2, intArr3...)
	fmt.Println(intArr2) // [4, 5, 6, 7, 8, 9, 10]

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // map[]

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"]) // 23
	fmt.Println(myMap2["Rakib"]) // 0 as default
}