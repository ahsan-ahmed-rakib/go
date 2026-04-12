package main 

import (
		"fmt"
		"errors"
)

func main() {
	printMe("Rakib")

	var result, remainder, err = intDiv(24, 3)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of division is %v\n", result)
	} else {
		fmt.Printf("Result is %v with remainder is %v\n", result, remainder)
	}
}

func printMe(value string){
	fmt.Println(value)
}

func intDiv(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Can not devide by zero(0)")
		return 0, 0, err
	}

	var result = numerator / denominator
	var remainder = numerator % denominator
	
	return result, remainder, err
}