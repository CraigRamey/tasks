package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Todo struct {
	Task       string `json:"task"`
	IsComplete bool   `json:"isComplete"`
}

func (t Todo) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func getTodos() []Todo {
	todos := make([]Todo, 0)
	raw, err := ioutil.ReadFile("./todos.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &todos)
	return todos
}

func main() {
	todos := getTodos()
	fmt.Println(todos)
	for _, todo := range todos {
		fmt.Println(todo)
		fmt.Println(todo.toString())
	}
}
