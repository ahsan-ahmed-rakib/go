package main

import "fmt"

func main() {
	
	var taskItems = []string {"Watch go tutorial", "Do go code after see video", "Build API with go"}
	println("This is main function")

	// printTexts(taskItems)
	addTask(taskItems, "Learn GO function")
}

func printTexts(taskItems []string) {
	fmt.Println("Welcome to our Todolist App!")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) {
	var updatedTasks = append(taskItems, newTask)
	printTexts(updatedTasks)
}