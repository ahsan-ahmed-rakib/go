package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch GO Crash Course"
var fullGolang = "Watch GO Full Course"
var taskItems = []string{shortGolang, fullGolang}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8000", nil)
}

func showTasks(res http.ResponseWriter, req *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(res, task)
	}
}

func helloUser(res http.ResponseWriter, req *http.Request) {
	var greet = "Hello Rakib, this is API with GO"
	fmt.Fprintln(res, greet)
}
