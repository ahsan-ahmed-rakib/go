package main

import "fmt"

func main() {
	fmt.Println("Welcome to our Todolist App!")

	var taskItems = []string {"Watch go tutorial", "Do go code after see video", "Build API with go"}

	for index, task := range taskItems {
		// fmt.Println(index + 1,":", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}

	// fmt.Println("Tasks: ", taskItems)
}
